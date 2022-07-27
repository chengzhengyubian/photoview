package models

import (
	"crypto/rand"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"rds-data-20220330/client"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	Model
	Username string  `gorm:"unique;size:128"`
	Password *string `gorm:"size:256"`
	// RootPath string  `gorm:"size:512`
	Albums []Album `gorm:"many2many:user_albums;constraint:OnDelete:CASCADE;"`
	Admin  bool    `gorm:"default:false"`
}

type UserMediaData struct {
	ModelTimestamps
	UserID   int  `gorm:"primaryKey;autoIncrement:false"`
	MediaID  int  `gorm:"primaryKey;autoIncrement:false"`
	Favorite bool `gorm:"not null;default:false"`
}

type UserAlbums struct {
	UserID  int `gorm:"primaryKey;autoIncrement:false;constraint:OnDelete:CASCADE;"`
	AlbumID int `gorm:"primaryKey;autoIncrement:false;constraint:OnDelete:CASCADE;"`
}

type AccessToken struct {
	Model
	UserID int       `gorm:"not null;index"`
	User   User      `gorm:"constraint:OnDelete:CASCADE;"`
	Value  string    `gorm:"not null;size:24;index"`
	Expire time.Time `gorm:"not null;index"`
}

type UserPreferences struct {
	Model
	UserID   int  `gorm:"not null;index"`
	User     User `gorm:"constraint:OnDelete:CASCADE;"`
	Language *LanguageTranslation
}

func (u *UserPreferences) BeforeSave(tx *gorm.DB) error {

	if u.Language != nil && *u.Language == "" {
		u.Language = nil
	}

	if u.Language != nil {
		lang_str := string(*u.Language)
		found_match := false
		for _, lang := range AllLanguageTranslation {
			if string(lang) == lang_str {
				found_match = true
				break
			}
		}

		if !found_match {
			return errors.New("invalid language value")
		}
	}

	return nil
}

var ErrorInvalidUserCredentials = errors.New("invalid credentials")

func AuthorizeUser(db *gorm.DB, username string, password string) (*User, error) {
	var user User

	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrorInvalidUserCredentials
		}
		return nil, errors.Wrap(result.Error, "failed to get user by username when authorizing")
	}

	if user.Password == nil {
		return nil, errors.New("user does not have a password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, ErrorInvalidUserCredentials
		} else {
			return nil, errors.Wrap(err, "compare user password hash")
		}
	}

	return &user, nil
}

func RegisterUser(username string, password *string, admin bool) (*User, error) {
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
	var req *client.ExecuteStatementRequest
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
	req = &o
	user := User{
		Username: username,
		Admin:    admin,
	}

	if password != nil {
		hashedPassBytes, err := bcrypt.GenerateFromPassword([]byte(*password), 12)
		if err != nil {
			return nil, errors.Wrap(err, "failed to hash password")
		}
		hashedPass := string(hashedPassBytes)

		user.Password = &hashedPass
	}
	//result := db.Create(&user)
	var ad int
	if admin == true {
		ad = 1
	} else {
		ad = 0
	}
	user.Admin = admin
	user.Username = username
	user.Password = password
	var sql_users_in string
	if password != nil {
		sql_users_in = "insert into users (username,password,admin) values (\"" + username + "\",\"" + *password + "\"," + strconv.Itoa(ad) + ")"
	} else {
		sql_users_in = "insert into users (username,admin) values (\"" + username + "\"," + strconv.Itoa(ad) + ")"
	}
	req.Sql = &sql_users_in
	res, err := client.ExecuteStatement(req)
	if err != nil {
		fmt.Println(res)
		return nil, errors.Wrap(err, "insert new user with password into database")
	}
	sql_users_se := "select * from users where username=\"" + username + "\""
	req.Sql = &sql_users_se
	res, err = client.ExecuteStatement(req)
	if err != nil {
		fmt.Println(err)
	}
	user.ID = int(*res.Body.Data.Records[0][0].LongValue)
	//if result.Error != nil {
	//	return nil, errors.Wrap(result.Error, "insert new user with password into database")
	//}
	return &user, nil
}

func (user *User) GenerateAccessToken(db *gorm.DB) (*AccessToken, error) {

	bytes := make([]byte, 24)
	if _, err := rand.Read(bytes); err != nil {
		return nil, errors.New(fmt.Sprintf("Could not generate token: %s\n", err.Error()))
	}
	const CHARACTERS = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i, b := range bytes {
		bytes[i] = CHARACTERS[b%byte(len(CHARACTERS))]
	}

	token_value := string(bytes)
	expire := time.Now().Add(14 * 24 * time.Hour)
	//timeStr := expire.Format("2006-01-02 15:04:05")
	token := AccessToken{
		UserID: user.ID,
		Value:  token_value,
		Expire: expire,
	}
	result := db.Create(&token)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "saving access token to database")
	}
	return &token, nil
}

// FillAlbums fill user.Albums with albums from database
func (user *User) FillAlbums(db *gorm.DB) error {
	// Albums already present
	if len(user.Albums) > 0 {
		return nil
	}

	if err := db.Model(&user).Association("Albums").Find(&user.Albums); err != nil {
		return errors.Wrap(err, "fill user albums")
	}

	return nil
}

func (user *User) OwnsAlbum(db *gorm.DB, album *Album) (bool, error) {

	if err := user.FillAlbums(db); err != nil {
		return false, err
	}

	albumIDs := make([]int, 0)
	for _, a := range user.Albums {
		albumIDs = append(albumIDs, a.ID)
	}

	filter := func(query *gorm.DB) *gorm.DB {
		return query.Where("id IN (?)", albumIDs)
	}

	ownedParents, err := album.GetParents(db, filter)
	if err != nil {
		return false, err
	}

	return len(ownedParents) > 0, nil
}

// FavoriteMedia sets/clears a media as favorite for the user
func (user *User) FavoriteMedia(db *gorm.DB, mediaID int, favorite bool) (*Media, error) {
	userMediaData := UserMediaData{
		UserID:   user.ID,
		MediaID:  mediaID,
		Favorite: favorite,
	}

	if err := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&userMediaData).Error; err != nil {
		return nil, errors.Wrapf(err, "update user favorite media in database")
	}

	var media Media
	if err := db.First(&media, mediaID).Error; err != nil {
		return nil, errors.Wrap(err, "get media from database after favorite update")
	}

	return &media, nil
}
