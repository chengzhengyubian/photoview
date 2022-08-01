package actions

import (
	"time"

	"github.com/photoview/photoview/api/database/drivers"
	"github.com/photoview/photoview/api/graphql/models"
	"gorm.io/gorm"
)

//这里注意一下
func MyTimeline(db *gorm.DB, user *models.User, paginate *models.Pagination, onlyFavorites *bool, fromDate *time.Time) ([]*models.Media, error) {

	query := db.
		Joins("JOIN albums ON media.album_id = albums.id").
		Where("albums.id IN (?)", db.Table("user_albums").Select("user_albums.album_id").Where("user_id = ?", user.ID))

	switch drivers.GetDatabaseDriverType(db) {
	case drivers.POSTGRES:
		query = query.
			Order("DATE_TRUNC('year', date_shot) DESC").
			Order("DATE_TRUNC('month', date_shot) DESC").
			Order("DATE_TRUNC('day', date_shot) DESC").
			Order("albums.title ASC").
			Order("media.date_shot DESC")
	case drivers.SQLITE:
		query = query.
			Order("strftime('%j', media.date_shot) DESC"). // convert to day of year 001-366
			Order("albums.title ASC").
			Order("TIME(media.date_shot) DESC")
	default:
		query = query.
			Order("YEAR(media.date_shot) DESC").
			Order("MONTH(media.date_shot) DESC").
			Order("DAY(media.date_shot) DESC").
			Order("albums.title ASC").
			Order("TIME(media.date_shot) DESC")
	}

	if fromDate != nil {
		query = query.Where("media.date_shot < ?", fromDate)
	}
	//SELECT `media`.`id`,`media`.`created_at`,`media`.`updated_at`,`media`.`title`,`media`.`path`,`media`.`path_hash`,`media`.`album_id`,`media`.`exif_id`,`media`.`date_shot`,`media`.`type`,`media`.`video_metadata_id`,`media`.`side_car_path`,`media`.`side_car_hash`,`media`.`blurhash` FROM `media` JOIN albums ON media.album_id = albums.id WHERE albums.id IN (SELECT user_albums.album_id FROM `user_albums` WHERE user_id = 2) ORDER BY YEAR(media.date_shot) DESC,MONTH(media.date_shot) DESC,DAY(media.date_shot) DESC,albums.title ASC,TIME(media.date_shot) DESC LIMIT 200 OFFSET 3
	if onlyFavorites != nil && *onlyFavorites {
		query = query.Where("media.id IN (?)", db.Table("user_media_data").Select("user_media_data.media_id").Where("user_media_data.user_id = ?", user.ID).Where("user_media_data.favorite"))
	}

	query = models.FormatSQL(query, nil, paginate)
	//SELECT `media`.`id`,`media`.`created_at`,`media`.`updated_at`,`media`.`title`,`media`.`path`,`media`.`path_hash`,`media`.`album_id`,`media`.`exif_id`,`media`.`date_shot`,`media`.`type`,`media`.`video_metadata_id`,`media`.`side_car_path`,`media`.`side_car_hash`,`media`.`blurhash` FROM `media` JOIN albums ON media.album_id = albums.id WHERE albums.id IN (SELECT user_albums.album_id FROM `user_albums` WHERE user_id = 2) ORDER BY YEAR(media.date_shot) DESC,MONTH(media.date_shot) DESC,DAY(media.date_shot) DESC,albums.title ASC,TIME(media.date_shot) DESC LIMIT 200 OFFSET 3
	var media []*models.Media
	if err := query.Find(&media).Error; err != nil { //SELECT `media`.`id`,`media`.`created_at`,`media`.`updated_at`,`media`.`title`,`media`.`path`,`media`.`path_hash`,`media`.`album_id`,`media`.`exif_id`,`media`.`date_shot`,`media`.`type`,`media`.`video_metadata_id`,`media`.`side_car_path`,`media`.`side_car_hash`,`media`.`blurhash` FROM `media` JOIN albums ON media.album_id = albums.id WHERE albums.id IN (SELECT user_albums.album_id FROM `user_albums` WHERE user_id = 2) AND media.id IN (SELECT user_media_data.media_id FROM `user_media_data` WHERE user_media_data.user_id = 2 AND user_media_data.favorite) ORDER BY YEAR(media.date_shot) DESC,MONTH(media.date_shot) DESC,DAY(media.date_shot) DESC,albums.title ASC,TIME(media.date_shot) DESC LIMIT 200 OFFSET 1
		return nil, err
	}

	//只显示喜爱
	//SELECT `media`.`id`,`media`.`created_at`,`media`.`updated_at`,`media`.`title`,`media`.`path`,`media`.`path_hash`,`media`.`album_id`,`media`.`exif_id`,`media`.`date_shot`,`media`.`type`,`media`.`video_metadata_id`,`media`.`side_car_path`,`media`.`side_car_hash`,`media`.`blurhash` FROM `media` JOIN albums ON media.album_id = albums.id WHERE albums.id IN (SELECT user_albums.album_id FROM `user_albums` WHERE user_id = 2) AND media.date_shot < '2022-01-01 00:00:00' AND media.id IN (SELECT user_media_data.media_id FROM `user_media_data` WHERE user_media_data.user_id = 2 AND user_media_data.favorite) ORDER BY YEAR(media.date_shot) DESC,MONTH(media.date_shot) DESC,DAY(media.date_shot) DESC,albums.title ASC,TIME(media.date_shot) DESC LIMIT 200 OFFSET 1
	return media, nil

	//从今天开始
	//SELECT `media`.`id`,`media`.`created_at`,`media`.`updated_at`,`media`.`title`,`media`.`path`,`media`.`path_hash`,`media`.`album_id`,`media`.`exif_id`,`media`.`date_shot`,`media`.`type`,`media`.`video_metadata_id`,`media`.`side_car_path`,`media`.`side_car_hash`,`media`.`blurhash` FROM `media` JOIN albums ON media.album_id = albums.id WHERE albums.id IN (SELECT user_albums.album_id FROM `user_albums` WHERE user_id = 2) ORDER BY YEAR(media.date_shot) DESC,MONTH(media.date_shot) DESC,DAY(media.date_shot) DESC,albums.title ASC,TIME(media.date_shot) DESC LIMIT 200 OFFSET 3
}
