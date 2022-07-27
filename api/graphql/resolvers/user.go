package resolvers

import (
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path"
	"rds-data-20220330/client"
	"strconv"

	api "github.com/photoview/photoview/api/graphql"
	"github.com/photoview/photoview/api/graphql/auth"
	"github.com/photoview/photoview/api/graphql/models"
	"github.com/photoview/photoview/api/graphql/models/actions"
	"github.com/photoview/photoview/api/scanner"
	"github.com/photoview/photoview/api/scanner/face_detection"
	"github.com/photoview/photoview/api/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userResolver struct {
	*Resolver
}

func (r *Resolver) User() api.UserResolver {
	return &userResolver{r}
}

//func (r *queryResolver) User(ctx context.Context, order *models.Ordering, paginate *models.Pagination) ([]*models.User, error) {
//
//	var users []*models.User
//
//	if err := models.FormatSQL(r.DB(ctx).Model(models.User{}), order, paginate).Find(&users).Error; err != nil {
//		return nil, err
//	}
//
//	return users, nil
//}

func (r *queryResolver) User(ctx context.Context, order *models.Ordering, paginate *models.Pagination) ([]*models.User, error) {
	//配置数据库连接
	database := "photoview"
	resourceArn := "acs:rds:cn-hangzhou:2304982093480287:dbInstance/rm-dingliang024"
	secretArn := "acs:rds:cn-hangzhou:2304982093480287:rds-db-credentials/aurora-taKgr1"
	var o client.ExecuteStatementRequest

	o.Database = &database
	o.ResourceArn = &resourceArn
	o.SecretArn = &secretArn

	//sql执行后的返回体
	//var res *client.ExecuteStatementResponse
	//sql执行的请求体
	var o1 *client.ExecuteStatementRequest
	//执行sql的客户端
	var client client.Client
	//客户端的配置
	var config openapi.Config
	//初始化客户端配置
	config.SetAccessKeyId("ACSTQDkNtSMrZtwL")
	config.SetAccessKeySecret("zXJ7QF79Oz")
	config.SetEndpoint("rds-data-daily.aliyuncs.com")
	client.Init(&config)
	//请求体指向真实的
	o1 = &o
	sql_users_se := "select * from users"
	o1.Sql = &sql_users_se
	res, err := client.ExecuteStatement(o1)
	if err != nil {
		fmt.Println(err)
	}
	num1 := len(res.Body.Data.Records)
	//num2 := len(res.Body.Data.Records[0])
	var users []*models.User

	for i := 0; i < num1; i++ {
		var user models.User
		user.ID = int(*res.Body.Data.Records[i][0].LongValue)
		user.Password = res.Body.Data.Records[i][4].StringValue
		user.Username = *res.Body.Data.Records[i][3].StringValue
		user.Admin = *res.Body.Data.Records[i][5].BooleanValue
		//users[i] = &user
		users = append(users, &user)
	}
	return users, nil
}

func (r *userResolver) Albums(ctx context.Context, user *models.User) ([]*models.Album, error) {
	user.FillAlbums(r.DB(ctx))

	pointerAlbums := make([]*models.Album, len(user.Albums))
	for i, album := range user.Albums {
		pointerAlbums[i] = &album
	}

	return pointerAlbums, nil
}

func (r *userResolver) RootAlbums(ctx context.Context, user *models.User) (albums []*models.Album, err error) {
	db := r.DB(ctx)

	err = db.Model(&user).
		Where("albums.parent_album_id NOT IN (?)",
			db.Table("user_albums").
				Select("albums.id").
				Joins("JOIN albums ON albums.id = user_albums.album_id AND user_albums.user_id = ?", user.ID),
		).Or("albums.parent_album_id IS NULL").
		Association("Albums").Find(&albums)

	return
}

func (r *queryResolver) MyUser(ctx context.Context) (*models.User, error) {

	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, auth.ErrUnauthorized
	}

	return user, nil
}

func (r *mutationResolver) AuthorizeUser(ctx context.Context, username string, password string) (*models.AuthorizeResult, error) {
	db := r.DB(ctx)
	user, err := models.AuthorizeUser(db, username, password)
	if err != nil {
		return &models.AuthorizeResult{
			Success: false,
			Status:  err.Error(),
		}, nil
	}

	var token *models.AccessToken

	transactionError := db.Transaction(func(tx *gorm.DB) error {
		token, err = user.GenerateAccessToken(tx)
		if err != nil {
			return err
		}

		return nil
	})

	if transactionError != nil {
		return nil, transactionError
	}

	return &models.AuthorizeResult{
		Success: true,
		Status:  "ok",
		Token:   &token.Value,
	}, nil
}

func (r *mutationResolver) InitialSetupWizard(ctx context.Context, username string, password string, rootPath string) (*models.AuthorizeResult, error) {
	db := r.DB(ctx)
	siteInfo, err := models.GetSiteInfo(db)
	if err != nil {
		return nil, err
	}

	if !siteInfo.InitialSetup {
		return nil, errors.New("not initial setup")
	}

	rootPath = path.Clean(rootPath)

	var token *models.AccessToken

	transactionError := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE site_info SET initial_setup = false").Error; err != nil {
			return err
		}

		user, err := models.RegisterUser(username, &password, true)
		if err != nil {
			return err
		}

		_, err = scanner.NewRootAlbum(tx, rootPath, user)
		if err != nil {
			return err
		}

		token, err = user.GenerateAccessToken(tx)
		if err != nil {
			return err
		}

		return nil
	})

	if transactionError != nil {
		return &models.AuthorizeResult{
			Success: false,
			Status:  err.Error(),
		}, nil
	}

	return &models.AuthorizeResult{
		Success: true,
		Status:  "ok",
		Token:   &token.Value,
	}, nil
}

func (r *queryResolver) MyUserPreferences(ctx context.Context) (*models.UserPreferences, error) {
	//配置数据库连接
	database := "photoview"
	resourceArn := "acs:rds:cn-hangzhou:2304982093480287:dbInstance/rm-dingliang024"
	secretArn := "acs:rds:cn-hangzhou:2304982093480287:rds-db-credentials/aurora-taKgr1"
	var o client.ExecuteStatementRequest

	o.Database = &database
	o.ResourceArn = &resourceArn
	o.SecretArn = &secretArn

	//sql执行后的返回体
	var res *client.ExecuteStatementResponse
	//sql执行的请求体
	var o1 *client.ExecuteStatementRequest
	//执行sql的客户端
	var client client.Client
	//客户端的配置
	var config openapi.Config
	//初始化客户端配置
	config.SetAccessKeyId("ACSTQDkNtSMrZtwL")
	config.SetAccessKeySecret("zXJ7QF79Oz")
	config.SetEndpoint("rds-data-daily.aliyuncs.com")
	client.Init(&config)
	//请求体指向真实的
	o1 = &o

	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, auth.ErrUnauthorized
	}
	userPref := models.UserPreferences{
		UserID: user.ID,
	}
	id := strconv.Itoa(user.ID)
	//if err := r.DB(ctx).Where("user_id = ?", user.ID).FirstOrCreate(&userPref).Error; err != nil {
	//	return nil, err
	//}
	sql1 := "select * from user_preferences where user_id=" + id
	o1.Sql = &sql1
	res, err := client.ExecuteStatement(o1)
	if err != nil {
		fmt.Println(err)
	}
	if len(res.Body.Data.Records) == 0 {
		sql2 := "insert into user_preferences (user_id,language,created_at) VALUES (" + id + ",\"English\",NOW())"
		o1.Sql = &sql2
		client.ExecuteStatement(o1)
		sql := "select * from user_preferences where user_id=" + id
		o1.Sql = &sql
		res, err := client.ExecuteStatement(o1)
		if err != nil {
			fmt.Println(err)
		}
		var language *string
		var langTrans *models.LanguageTranslation = nil
		language = res.Body.Data.Records[0][4].StringValue
		lng := models.LanguageTranslation(*language)
		langTrans = &lng
		userPref.Language = langTrans
		userPref.ID = int(*res.Body.Data.Records[0][0].LongValue)
	} else {
		var language *string
		language = res.Body.Data.Records[0][4].StringValue
		var langTrans *models.LanguageTranslation = nil
		if language != nil {
			lng := models.LanguageTranslation(*language)
			langTrans = &lng
		}
		userPref.Language = langTrans
		userPref.ID = int(*res.Body.Data.Records[0][0].LongValue)
	}
	return &userPref, nil
}

func (r *mutationResolver) ChangeUserPreferences(ctx context.Context, language *string) (*models.UserPreferences, error) {
	//配置数据库连接
	database := "photoview"
	resourceArn := "acs:rds:cn-hangzhou:2304982093480287:dbInstance/rm-dingliang024"
	secretArn := "acs:rds:cn-hangzhou:2304982093480287:rds-db-credentials/aurora-taKgr1"

	var o client.ExecuteStatementRequest

	o.Database = &database
	o.ResourceArn = &resourceArn
	o.SecretArn = &secretArn

	var str string
	str = *language

	//sql执行后的返回体
	var res *client.ExecuteStatementResponse
	//sql执行的请求体
	var o1 *client.ExecuteStatementRequest
	//执行sql的客户端
	var client client.Client
	//客户端的配置
	var config openapi.Config
	//初始化客户端配置
	config.SetAccessKeyId("ACSTQDkNtSMrZtwL")
	config.SetAccessKeySecret("zXJ7QF79Oz")
	config.SetEndpoint("rds-data-daily.aliyuncs.com")
	client.Init(&config)
	//请求体指向真实的
	o1 = &o

	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, auth.ErrUnauthorized
	}

	var langTrans *models.LanguageTranslation = nil
	if language != nil {
		lng := models.LanguageTranslation(*language)
		langTrans = &lng
	}

	var userPref models.UserPreferences
	//下面是对这段逻辑的改写
	var id string
	id = strconv.Itoa(user.ID)
	sql2 := "select * from user_preferences where user_id =" + id
	//如果执行这个操作之后没有值，那么执行下面的操作，否则啥也不干
	o1.Sql = &sql2
	res, err := client.ExecuteStatement(o1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Body.Data.Records)
	if len(res.Body.Data.Records) == 0 {
		//执行，插入的操作
		sql3 := "insert into user_preferences (user_id,language,created_at) VALUES (" + id + ",\"English\",NOW())"
		o1.Sql = &sql3
		//var res *client.ExecuteStatementResponse
		client.ExecuteStatement(o1)
	}

	//执行，插入的操作
	//sql3 := "insert into user_preferences (user_id) VALUES (" + id + ")\""
	//更新的操作
	sql := "update user_preferences set (language,updated_at) values(\"" + str + "\",NOW()) where user_id=" + id
	o1.Sql = &sql
	client.ExecuteStatement(o1)
	o1.Sql = &sql2
	res, err = client.ExecuteStatement(o1)
	if err != nil {
		fmt.Println(err)
	}

	o1.Sql = &sql2
	res, err = client.ExecuteStatement(o1)

	userPref.ID = int(*res.Body.Data.Records[0][0].LongValue)
	userPref.Language = langTrans
	userPref.UserID = user.ID
	return &userPref, nil
}

//Admin queries
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, username *string, password *string, admin *bool) (*models.User, error) {
	db := r.DB(ctx)

	if username == nil && password == nil && admin == nil {
		return nil, errors.New("no updates requested")
	}

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	if username != nil {
		user.Username = *username
	}

	if password != nil {
		hashedPassBytes, err := bcrypt.GenerateFromPassword([]byte(*password), 12)
		if err != nil {
			return nil, err
		}
		hashedPass := string(hashedPassBytes)

		user.Password = &hashedPass
	}

	if admin != nil {
		user.Admin = *admin
	}

	if err := db.Save(&user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to update user")
	}

	return &user, nil
}

//更改了更新用户的操作
//func (r *mutationResolver) UpdateUser(ctx context.Context, id int, username *string, password *string, admin *bool) (*models.User, error) {
//	//配置数据库连接
//	database := "photoview"
//	resourceArn := "acs:rds:cn-hangzhou:2304982093480287:dbInstance/rm-dingliang024"
//	secretArn := "acs:rds:cn-hangzhou:2304982093480287:rds-db-credentials/aurora-taKgr1"
//	var o client.ExecuteStatementRequest
//
//	o.Database = &database
//	o.ResourceArn = &resourceArn
//	o.SecretArn = &secretArn
//
//	//sql执行后的返回体
//	var res *client.ExecuteStatementResponse
//	//sql执行的请求体
//	var req *client.ExecuteStatementRequest
//	//执行sql的客户端
//	var client client.Client
//	//客户端的配置
//	var config openapi.Config
//	//初始化客户端配置
//	config.SetAccessKeyId("ACSTQDkNtSMrZtwL")
//	config.SetAccessKeySecret("zXJ7QF79Oz")
//	config.SetEndpoint("rds-data-daily.aliyuncs.com")
//	client.Init(&config)
//	//请求体指向真实的
//	req = &o
//
//	//db := r.DB(ctx)
//
//	if username == nil && password == nil && admin == nil {
//		return nil, errors.New("no updates requested")
//	}
//
//	var user models.User
//	//if err := db.First(&user, id).Error; err != nil {
//	//	return nil, err
//	//}
//	sql_users_se := "select * from users where id =" + strconv.Itoa(id) + "limit 1"
//	req.Sql = &sql_users_se
//	res, err := client.ExecuteStatement(req)
//	if err != nil {
//		return nil, err
//	}
//	//user.ID = int(*res.Body.Data.Records[0][0].LongValue)
//	user.Username = *res.Body.Data.Records[0][3].StringValue
//	fmt.Println(user.Username)
//	user.Password = res.Body.Data.Records[0][4].StringValue
//	user.Admin = *res.Body.Data.Records[0][5].BooleanValue
//	if username != nil {
//		user.Username = *username
//	}
//	if password != nil {
//		hashedPassBytes, err := bcrypt.GenerateFromPassword([]byte(*password), 12)
//		if err != nil {
//			return nil, err
//		}
//		hashedPass := string(hashedPassBytes)
//
//		user.Password = &hashedPass
//	}
//	var ad int
//	if admin != nil {
//		user.Admin = *admin
//	}
//	if user.Admin == true {
//		ad = 1
//	} else {
//		ad = 0
//	}
//	//if err := db.Save(&user).Error; err != nil {
//	//	return nil, errors.Wrap(err, "failed to update user")
//	//}
//	sql_users_up := "update users set username=" + user.Username + "\",password=\"" + *user.Password + "\",admin=" + strconv.Itoa(ad) + " where id =" + strconv.Itoa(id)
//	req.Sql = &sql_users_up
//	//if username != nil {
//	//	sql_users_up += "\"" + *username + "\""
//	//} else {
//	//	sql_users_up += "\"" + *name + "\""
//	//}
//	//sql_users_up += ",password"
//	//if password != nil {
//	//	sql_users_up += "\"" + *password + "\""
//	//} else {
//	//	sql_users_up += "\"" + *pass + "\""
//	//}
//	//sql_users_up += ",admin"
//	//if admin != nil {
//	//	sql_users_up += strconv.Itoa(ad)
//	//} else {
//	//	sql_users_up += strconv.Itoa(0)
//	//}
//	//sql_users_up += "where id ="
//	//sql_users_up += strconv.Itoa(id)
//	res, err = client.ExecuteStatement(req)
//	fmt.Println(res)
//	if err != nil {
//		return nil, err
//	}
//	//user.ID = id
//	//user.Password = password
//	//user.Admin = *admin
//	return &user, nil
//}

func (r *mutationResolver) CreateUser(ctx context.Context, username string, password *string, admin bool) (*models.User, error) {

	var user *models.User

	transactionError := r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		user, err = models.RegisterUser(username, password, admin)
		if err != nil {
			return err
		}

		return nil
	})

	if transactionError != nil {
		return nil, transactionError
	}

	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*models.User, error) {
	return actions.DeleteUser(r.DB(ctx), id)
}

func (r *mutationResolver) UserAddRootPath(ctx context.Context, id int, rootPath string) (*models.Album, error) {
	db := r.DB(ctx)

	rootPath = path.Clean(rootPath)

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	newAlbum, err := scanner.NewRootAlbum(db, rootPath, &user)
	if err != nil {
		return nil, err
	}

	return newAlbum, nil
}

func (r *mutationResolver) UserRemoveRootAlbum(ctx context.Context, userID int, albumID int) (*models.Album, error) {
	db := r.DB(ctx)

	var album models.Album
	if err := db.First(&album, albumID).Error; err != nil {
		return nil, err
	}

	var deletedAlbumIDs []int = nil

	transactionError := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Raw("DELETE FROM user_albums WHERE user_id = ? AND album_id = ?", userID, albumID).Error; err != nil {
			return err
		}

		children, err := album.GetChildren(tx, nil)
		if err != nil {
			return err
		}

		childAlbumIDs := make([]int, len(children))
		for i, child := range children {
			childAlbumIDs[i] = child.ID
		}

		result := tx.Exec("DELETE FROM user_albums WHERE user_id = ? and album_id IN (?)", userID, childAlbumIDs)
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("No relation deleted")
		}

		// Cleanup if no user owns the album anymore
		var userAlbumCount int
		if err := tx.Raw("SELECT COUNT(user_id) FROM user_albums WHERE album_id = ?", albumID).Scan(&userAlbumCount).Error; err != nil {
			return err
		}

		if userAlbumCount == 0 {
			deletedAlbumIDs = append(childAlbumIDs, albumID)
			childAlbumIDs = nil

			// Delete albums from database
			if err := tx.Delete(&models.Album{}, "id IN (?)", deletedAlbumIDs).Error; err != nil {
				deletedAlbumIDs = nil
				return err
			}
		}

		return nil
	})

	if transactionError != nil {
		return nil, transactionError
	}

	if deletedAlbumIDs != nil {
		// Delete albums from cache
		for _, id := range deletedAlbumIDs {
			cacheAlbumPath := path.Join(utils.MediaCachePath(), strconv.Itoa(id))

			if err := os.RemoveAll(cacheAlbumPath); err != nil {
				return nil, err
			}
		}

		// Reload faces as media might have been deleted
		if face_detection.GlobalFaceDetector != nil {
			if err := face_detection.GlobalFaceDetector.ReloadFacesFromDatabase(db); err != nil {
				return nil, err
			}
		}
	}

	return &album, nil
}
