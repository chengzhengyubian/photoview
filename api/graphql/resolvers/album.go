package resolvers

import (
	"context"
	"fmt"
	DataApi "github.com/photoview/photoview/api/dataapi"
	"time"

	api "github.com/photoview/photoview/api/graphql"
	"github.com/photoview/photoview/api/graphql/auth"
	"github.com/photoview/photoview/api/graphql/models"
	"github.com/photoview/photoview/api/graphql/models/actions"
	"github.com/pkg/errors"
)

func (r *queryResolver) MyAlbums(ctx context.Context, order *models.Ordering, paginate *models.Pagination, onlyRoot *bool, showEmpty *bool, onlyWithFavorites *bool) ([]*models.Album, error) {
	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, auth.ErrUnauthorized
	}

	return actions.MyAlbums(r.DB(ctx), user, order, paginate, onlyRoot, showEmpty, onlyWithFavorites)
}

func (r *queryResolver) Album(ctx context.Context, id int, tokenCredentials *models.ShareTokenCredentials) (*models.Album, error) {
	db := r.DB(ctx)
	if tokenCredentials != nil {

		shareToken, err := r.ShareToken(ctx, *tokenCredentials)
		if err != nil {
			return nil, err
		}

		if shareToken.Album != nil {
			if *shareToken.AlbumID == id {
				return shareToken.Album, nil
			}

			subAlbum, err := shareToken.Album.GetChildren(func(sql string) string {
				//return query.Where("sub_albums.id = ?", id)
				return sql + fmt.Sprintf(" where sub_albums.id =%v", id)
			}) //这里注意一下
			if err != nil {
				return nil, errors.Wrapf(err, "find sub album of share token (%s)", tokenCredentials.Token)
			}

			if len(subAlbum) > 0 {
				return subAlbum[0], nil
			}
		}
	}

	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, auth.ErrUnauthorized
	}

	return actions.Album(db, user, id)
}

func (r *Resolver) Album() api.AlbumResolver {
	return &albumResolver{r}
}

type albumResolver struct{ *Resolver }

//修改完，测试还有点小问题
func (r *albumResolver) Media(ctx context.Context, album *models.Album, order *models.Ordering, paginate *models.Pagination, onlyFavorites *bool) ([]*models.Media, error) {
	//db := r.DB(ctx)

	//query := db.
	//	Where("media.album_id = ?", album.ID).
	//	Where("media.id IN (?)", db.Model(&models.MediaURL{}).Select("media_urls.media_id").Where("media_urls.media_id = media.id"))
	var orderby string
	var limit, offset int
	orderby = *order.OrderBy
	limit = *paginate.Limit
	offset = *paginate.Offset
	var sql_media_se string
	if onlyFavorites != nil && *onlyFavorites == true {
		user := auth.UserFromContext(ctx)
		if user == nil {
			return nil, errors.New("cannot get favorite media without being authorized")
		}
		//
		//favoriteQuery := db.Model(&models.UserMediaData{
		//	UserID: user.ID,
		//}).Where("user_media_data.media_id = media.id").Where("user_media_data.favorite = true")
		//
		//query = query.Where("EXISTS (?)", favoriteQuery)
		sql_media_se = fmt.Sprintf("SELECT * FROM `media` WHERE media.album_id = %v AND media.id IN (SELECT media_urls.media_id FROM `media_urls` WHERE media_urls.media_id = media.id) AND EXISTS (SELECT * FROM `user_media_data` WHERE user_media_data.media_id = media.id AND user_media_data.favorite = true AND `user_media_data`.`user_id` = %v) ORDER BY `%v` LIMIT %v OFFSET %v", album.ID, user.ID, orderby, limit, offset)
	} else {
		sql_media_se = fmt.Sprintf("SELECT * FROM `media` WHERE media.album_id = %v AND media.id IN (SELECT media_urls.media_id FROM `media_urls` WHERE media_urls.media_id = media.id) ORDER BY `%v` LIMIT %v OFFSET %v", album.ID, orderby, limit, offset)
	}
	//query = models.FormatSQL(query, order, paginate)

	var media []*models.Media
	//if err := query.Find(&media).Error; err != nil { //SELECT * FROM `media` WHERE media.album_id = 1 AND media.id IN (SELECT media_urls.media_id FROM `media_urls` WHERE media_urls.media_id = media.id) ORDER BY `date_shot` LIMIT 200
	//	return nil, err
	//}
	// SELECT * FROM `media` WHERE media.album_id = 146 AND media.id IN (SELECT media_urls.media_id FROM `media_urls` WHERE media_urls.media_id = media.id) AND EXISTS (SELECT * FROM `user_media_data` WHERE user_media_data.media_id = media.id AND user_media_data.favorite = true AND `user_media_data`.`user_id` = 10) ORDER BY `date_shot` LIMIT 200 OFFSET 1
	dataApi, _ := DataApi.NewDataApiClient()
	res, err := dataApi.Query(sql_media_se)
	if len(res) == 0 {
		return nil, err
	}
	num := len(res)
	for i := 0; i < num; i++ {
		var Media models.Media
		Media.ID = DataApi.GetInt(res, i, 0)
		Media.CreatedAt = time.Unix(*res[i][1].LongValue/1000, 0)
		Media.UpdatedAt = time.Unix(*res[i][2].LongValue/1000, 0)
		Media.Title = *res[i][3].StringValue
		Media.Path = *res[i][4].StringValue
		Media.PathHash = *res[i][5].StringValue
		Media.AlbumID = int(*res[i][6].LongValue)
		Media.ExifID = DataApi.GetIntP(res, i, 7)
		Media.DateShot = time.Unix(*res[i][8].LongValue/1000, 0)
		if *res[0][9].StringValue == "photo" {
			Media.Type = models.MediaTypePhoto
		} else {
			Media.Type = models.MediaTypeVideo
		}
		Media.VideoMetadataID = DataApi.GetIntP(res, i, 10)
		Media.SideCarPath = DataApi.GetStringP(res, i, 11)
		Media.SideCarHash = DataApi.GetStringP(res, i, 12)
		Media.Blurhash = DataApi.GetStringP(res, i, 13)
		media = append(media, &Media)
	}
	return media, nil
}

func (r *albumResolver) Thumbnail(ctx context.Context, album *models.Album) (*models.Media, error) {
	return album.Thumbnail(r.DB(ctx))
}

//修改完，未测试
func (r *albumResolver) SubAlbums(ctx context.Context, parent *models.Album, order *models.Ordering, paginate *models.Pagination) ([]*models.Album, error) {

	var albums []*models.Album

	query := r.DB(ctx).Where("parent_album_id = ?", parent.ID)
	query = models.FormatSQL(query, order, paginate)

	//if err := query.Find(&albums).Error; err != nil { //SELECT * FROM `albums` WHERE parent_album_id = 1 ORDER BY `title`
	//	return nil, err
	//}
	var orderby string
	orderby = *order.OrderBy
	sql_albums_se := fmt.Sprintf("SELECT * FROM `albums` WHERE parent_album_id =%v ORDER BY '%v'", parent.ID, orderby)
	dataApi, _ := DataApi.NewDataApiClient()
	res, err := dataApi.Query(sql_albums_se)
	if err != nil {
		return nil, err
	}
	num := len(res)
	if len(res) == 0 {
		return nil, err
	}
	for i := 0; i < num; i++ {
		var album models.Album
		album.ID = int(*res[i][0].LongValue)
		album.CreatedAt = time.Unix(*res[i][1].LongValue/1000, 0)
		album.UpdatedAt = time.Unix(*res[i][2].LongValue/1000, 0)
		album.Title = *res[i][3].StringValue
		album.ParentAlbumID = DataApi.GetIntP(res, i, 4)
		album.Path = *res[i][5].StringValue
		album.PathHash = *res[i][6].StringValue
		album.CoverID = DataApi.GetIntP(res, i, 7)
		albums = append(albums, &album)
	}
	return albums, nil
}

func (r *albumResolver) ParentAlbum(ctx context.Context, obj *models.Album) (*models.Album, error) {
	panic("not implemented")
}

func (r *albumResolver) Owner(ctx context.Context, obj *models.Album) (*models.User, error) {
	panic("not implemented")
}

//修改完，未测试
func (r *albumResolver) Shares(ctx context.Context, album *models.Album) ([]*models.ShareToken, error) {

	var shareTokens []*models.ShareToken
	//if err := r.DB(ctx).Where("album_id = ?", album.ID).Find(&shareTokens).Error; err != nil { //SELECT * FROM `share_tokens` WHERE album_id = 146
	//	return nil, err
	//}
	sql_shareTokens_se := fmt.Sprintf("select * from share_tokens where album_id=%v", album.ID)
	dataApi, _ := DataApi.NewDataApiClient()
	res, err := dataApi.Query(sql_shareTokens_se)
	if err != nil {
		return nil, err
	}
	num := len(res)
	if len(res) == 0 {
		return nil, err
	}
	for i := 0; i < num; i++ {
		var shareToken models.ShareToken
		shareToken.ID = DataApi.GetInt(res, i, 0)
		shareToken.CreatedAt = time.Unix(DataApi.GetLong(res, i, 1)/1000, 0)
		shareToken.UpdatedAt = time.Unix(DataApi.GetLong(res, i, 2)/1000, 0)
		shareToken.Value = DataApi.GetString(res, i, 3)
		shareToken.OwnerID = DataApi.GetInt(res, i, 4)
		times := time.Unix(DataApi.GetLong(res, i, 5)/1000, 0)
		shareToken.Expire = &times
		shareToken.Password = DataApi.GetStringP(res, i, 6)
		shareToken.AlbumID = DataApi.GetIntP(res, i, 7)
		shareToken.MediaID = DataApi.GetIntP(res, i, 8)
		shareTokens = append(shareTokens, &shareToken)
	}
	return shareTokens, nil
}

func (r *albumResolver) Path(ctx context.Context, obj *models.Album) ([]*models.Album, error) {

	user := auth.UserFromContext(ctx)
	if user == nil {
		empty := make([]*models.Album, 0)
		return empty, nil
	}

	return actions.AlbumPath(r.DB(ctx), user, obj)
}

// Takes album_id, resets album.cover_id to 0 (null)
func (r *mutationResolver) ResetAlbumCover(ctx context.Context, albumID int) (*models.Album, error) {
	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, errors.New("unauthorized")
	}

	return actions.ResetAlbumCover(r.DB(ctx), user, albumID)
}

func (r *mutationResolver) SetAlbumCover(ctx context.Context, mediaID int) (*models.Album, error) {
	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, errors.New("unauthorized")
	}

	return actions.SetAlbumCover(r.DB(ctx), user, mediaID)
}
