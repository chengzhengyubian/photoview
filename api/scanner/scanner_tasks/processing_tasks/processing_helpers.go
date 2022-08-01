package processing_tasks

import (
	"fmt"
	"os"
	"path"

	"github.com/photoview/photoview/api/graphql/models"
	"github.com/photoview/photoview/api/scanner/media_encoding"
	"github.com/photoview/photoview/api/scanner/media_encoding/media_utils"
	"github.com/photoview/photoview/api/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Higher order function used to check if MediaURL for a given MediaPurpose exists
func makePhotoURLChecker(tx *gorm.DB, mediaID int) func(purpose models.MediaPurpose) (*models.MediaURL, error) {
	return func(purpose models.MediaPurpose) (*models.MediaURL, error) {
		var mediaURL []*models.MediaURL
		// SELECT * FROM `media_urls` WHERE purpose = 'high-res' AND media_id = 2
		result := tx.Where("purpose = ?", purpose).Where("media_id = ?", mediaID).Find(&mediaURL) //SELECT * FROM `media_urls` WHERE purpose = 'original' AND media_id = 2

		if result.Error != nil {
			return nil, result.Error
		}

		if result.RowsAffected > 0 {
			return mediaURL[0], nil
		}

		return nil, nil
	}
}

func generateUniqueMediaNamePrefixed(prefix string, mediaPath string, extension string) string {
	mediaName := fmt.Sprintf("%s_%s_%s", prefix, path.Base(mediaPath), utils.GenerateToken())
	mediaName = models.SanitizeMediaName(mediaName)
	mediaName = mediaName + extension
	return mediaName
}

func generateUniqueMediaName(mediaPath string) string {

	filename := path.Base(mediaPath)
	baseName := filename[0 : len(filename)-len(path.Ext(filename))]
	baseExt := path.Ext(filename)

	mediaName := fmt.Sprintf("%s_%s", baseName, utils.GenerateToken())
	mediaName = models.SanitizeMediaName(mediaName) + baseExt

	return mediaName
}

func saveOriginalPhotoToDB(tx *gorm.DB, photo *models.Media, imageData *media_encoding.EncodeMediaData, photoDimensions *media_utils.PhotoDimensions) (*models.MediaURL, error) {
	originalImageName := generateUniqueMediaName(photo.Path)

	contentType, err := imageData.ContentType()
	if err != nil {
		return nil, err
	}

	fileStats, err := os.Stat(photo.Path)
	if err != nil {
		return nil, errors.Wrap(err, "reading file stats of original photo")
	}

	mediaURL := models.MediaURL{
		Media:       photo,
		MediaName:   originalImageName,
		Width:       photoDimensions.Width,
		Height:      photoDimensions.Height,
		Purpose:     models.MediaOriginal,
		ContentType: string(*contentType),
		FileSize:    fileStats.Size(),
	}

	if err := tx.Create(&mediaURL).Error; err != nil { //INSERT INTO `media_urls` (`created_at`,`updated_at`,`media_id`,`media_name`,`width`,`height`,`purpose`,`content_type`,`file_size`) VALUES ('2022-08-01 17:28:21.117','2022-08-01 17:28:21.117',0,'自我介绍_yDPkSVVK.png',1844,1074,'original','image/png',1701845)
		return nil, errors.Wrapf(err, "inserting original photo url: %d, %s", photo.ID, photo.Title)
	}

	return &mediaURL, nil
}
