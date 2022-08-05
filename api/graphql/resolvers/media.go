package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	DataApi "github.com/photoview/photoview/api/dataapi"
	"strconv"
	"strings"
	"time"

	"github.com/photoview/photoview/api/dataloader"
	api "github.com/photoview/photoview/api/graphql"
	"github.com/photoview/photoview/api/graphql/auth"
	"github.com/photoview/photoview/api/graphql/models"
	"github.com/photoview/photoview/api/graphql/models/actions"
	"github.com/photoview/photoview/api/scanner/face_detection"
	"github.com/pkg/errors"
)

func (r *queryResolver) MyMedia(ctx context.Context, order *models.Ordering, paginate *models.Pagination) ([]*models.Media, error) {
	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, errors.New("unauthorized")
	}

	return actions.MyMedia(r.DB(ctx), user, order, paginate)
}

//修改中
func (r *queryResolver) Media(ctx context.Context, id int, tokenCredentials *models.ShareTokenCredentials) (*models.Media, error) {
	db := r.DB(ctx)
	if tokenCredentials != nil {

		shareToken, err := r.ShareToken(ctx, *tokenCredentials)
		if err != nil {
			return nil, err
		}

		if *shareToken.MediaID == id {
			return shareToken.Media, nil
		}
	}

	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, auth.ErrUnauthorized
	}

	var media models.Media

	err := db.
		Joins("Album").
		Where("media.id = ?", id).
		Where("EXISTS (SELECT * FROM user_albums WHERE user_albums.album_id = media.album_id AND user_albums.user_id = ?)", user.ID).
		Where("media.id IN (?)", db.Model(&models.MediaURL{}).Select("media_id").Where("media_urls.media_id = media.id")).
		First(&media).Error

	sql_album_se := fmt.Sprintf("select media.* from media join album")

	if err != nil {
		return nil, errors.Wrap(err, "could not get media by media_id and user_id from database")
	}

	return &media, nil
}

//修改完，未测试
func (r *queryResolver) MediaList(ctx context.Context, ids []int) ([]*models.Media, error) {
	//db := r.DB(ctx)
	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, auth.ErrUnauthorized
	}

	if len(ids) == 0 {
		return nil, errors.New("no ids provided")
	}

	var media []*models.Media
	//err := db.Model(&media).
	//	Joins("LEFT JOIN user_albums ON user_albums.album_id = media.album_id").
	//	Where("media.id IN ?", ids).
	//	Where("user_albums.user_id = ?", user.ID).
	//	Find(&media).Error
	//
	//if err != nil {
	//	return nil, errors.Wrap(err, "could not get media list by media_id and user_id from database")
	//}

	id, _ := json.Marshal(ids)
	IDs := strings.Trim(string(id), "[]")
	sql_media_se := fmt.Sprintf("select media.* from media left join user_albums ON user_albums.album_id = media.album_id where media_id in (%v) and user_albums.user_id =%v", IDs, user.ID)
	dataApi, _ := DataApi.NewDataApiClient()
	res, _ := dataApi.Query(sql_media_se)
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

type mediaResolver struct {
	*Resolver
}

func (r *Resolver) Media() api.MediaResolver {
	return &mediaResolver{r}
}

func (r *mediaResolver) Type(ctx context.Context, media *models.Media) (models.MediaType, error) {
	formattedType := models.MediaType(strings.Title(string(media.Type)))
	return formattedType, nil
}

//修改完
func (r *mediaResolver) Album(ctx context.Context, obj *models.Media) (*models.Album, error) {
	var album models.Album
	//err := r.DB(ctx).Find(&album, obj.AlbumID).Error // SELECT * FROM `albums` WHERE `albums`.`id` = 1
	sql_albums_se := "SELECT * FROM `albums` WHERE `albums`.`id` =" + strconv.Itoa(obj.AlbumID)
	dataApi, _ := DataApi.NewDataApiClient()
	res, err := dataApi.Query(sql_albums_se)
	album.ID = DataApi.GetInt(res, 0, 0)
	album.CreatedAt = time.Unix(*res[0][1].LongValue/1000, 0)
	album.UpdatedAt = time.Unix(*res[0][2].LongValue/1000, 0)
	album.Title = DataApi.GetString(res, 0, 3)
	album.ParentAlbumID = DataApi.GetIntP(res, 0, 4)
	album.Path = DataApi.GetString(res, 0, 5)
	album.PathHash = DataApi.GetString(res, 0, 6)
	album.CoverID = DataApi.GetIntP(res, 0, 7)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

//未修改
func (r *mediaResolver) Shares(ctx context.Context, media *models.Media) ([]*models.ShareToken, error) {
	var shareTokens []*models.ShareToken
	if err := r.DB(ctx).Where("media_id = ?", media.ID).Find(&shareTokens).Error; err != nil {
		return nil, errors.Wrapf(err, "get shares for media (%s)", media.Path)
	}

	return shareTokens, nil
}

//未修改
func (r *mediaResolver) Downloads(ctx context.Context, media *models.Media) ([]*models.MediaDownload, error) {

	var mediaUrls []*models.MediaURL
	if err := r.DB(ctx).Where("media_id = ?", media.ID).Find(&mediaUrls).Error; err != nil {
		return nil, errors.Wrapf(err, "get downloads for media (%s)", media.Path)
	}

	downloads := make([]*models.MediaDownload, 0)

	for _, url := range mediaUrls {

		var title string
		switch {
		case url.Purpose == models.MediaOriginal:
			title = "Original"
		case url.Purpose == models.PhotoThumbnail:
			title = "Small"
		case url.Purpose == models.PhotoHighRes:
			title = "Large"
		case url.Purpose == models.VideoThumbnail:
			title = "Video thumbnail"
		case url.Purpose == models.VideoWeb:
			title = "Web optimized video"
		}

		downloads = append(downloads, &models.MediaDownload{
			Title:    title,
			MediaURL: url,
		})
	}

	return downloads, nil
}

func (r *mediaResolver) HighRes(ctx context.Context, media *models.Media) (*models.MediaURL, error) {
	if media.Type != models.MediaTypePhoto {
		return nil, nil
	}

	return dataloader.For(ctx).MediaHighres.Load(media.ID)
}

func (r *mediaResolver) Thumbnail(ctx context.Context, media *models.Media) (*models.MediaURL, error) {
	return dataloader.For(ctx).MediaThumbnail.Load(media.ID)
}

func (r *mediaResolver) VideoWeb(ctx context.Context, media *models.Media) (*models.MediaURL, error) {
	if media.Type != models.MediaTypeVideo {
		return nil, nil
	}

	return dataloader.For(ctx).MediaVideoWeb.Load(media.ID)
}

//未修改
func (r *mediaResolver) Exif(ctx context.Context, media *models.Media) (*models.MediaEXIF, error) {
	if media.Exif != nil {
		return media.Exif, nil
	}
	var exif models.MediaEXIF
	if err := r.DB(ctx).Model(&media).Association("Exif").Find(&exif); err != nil {
		return nil, err
	}

	return &exif, nil
}

func (r *mediaResolver) Favorite(ctx context.Context, media *models.Media) (bool, error) {
	user := auth.UserFromContext(ctx)
	if user == nil {
		return false, auth.ErrUnauthorized
	}

	return dataloader.For(ctx).UserMediaFavorite.Load(&models.UserMediaData{
		UserID:  user.ID,
		MediaID: media.ID,
	})
}

func (r *mutationResolver) FavoriteMedia(ctx context.Context, mediaID int, favorite bool) (*models.Media, error) {

	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, auth.ErrUnauthorized
	}

	return user.FavoriteMedia(r.DB(ctx), mediaID, favorite)
}

//未修改
func (r *mediaResolver) Faces(ctx context.Context, media *models.Media) ([]*models.ImageFace, error) {
	if face_detection.GlobalFaceDetector == nil {
		return []*models.ImageFace{}, nil
	}

	if media.Faces != nil {
		return media.Faces, nil
	}

	var faces []*models.ImageFace
	if err := r.DB(ctx).Model(&media).Association("Faces").Find(&faces); err != nil {
		return nil, err
	}

	return faces, nil
}
