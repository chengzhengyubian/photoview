package exif

import (
	"fmt"
	DataApi "github.com/photoview/photoview/api/dataapi"
	"log"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/photoview/photoview/api/graphql/models"
)

type ExifParser interface {
	ParseExif(media_path string) (*models.MediaEXIF, error)
}

var globalExifParser ExifParser

func InitializeEXIFParser() {
	// Decide between internal or external Exif parser
	exiftoolParser, err := NewExiftoolParser()

	if err != nil {
		log.Printf("Failed to get exiftool, using internal exif parser instead: %v\n", err)
		globalExifParser = NewInternalExifParser()
	} else {
		log.Println("Found exiftool")
		globalExifParser = exiftoolParser
	}
}

// SaveEXIF scans the media file for exif metadata and saves it in the database if found
//修改中
func SaveEXIF(tx *gorm.DB, media *models.Media) (*models.MediaEXIF, error) {

	{
		// Check if EXIF data already exists
		if media.ExifID != nil {

			var exif models.MediaEXIF
			if err := tx.First(&exif, media.ExifID).Error; err != nil {
				return nil, errors.Wrap(err, "get EXIF for media from database")
			}
			sql_media_exif_select := fmt.Sprintf("select * from media_exif where id=%v", media.ExifID)
			dataApi, _ := DataApi.NewDataApiClient()
			res, _ := dataApi.Query(sql_media_exif_select)
			exif.ID = DataApi.GetInt(res, 0, 0)
			exif.CreatedAt = time.Unix(DataApi.GetLong(res, 0, 1)/1000, 0)
			exif.UpdatedAt = time.Unix(DataApi.GetLong(res, 0, 2)/1000, 0)
			exif.Camera = DataApi.GetStringP(res, 0, 3)
			exif.Maker = DataApi.GetStringP(res, 0, 4)
			exif.Lens = DataApi.GetStringP(res, 0, 5)
			date := time.Unix(DataApi.GetLong(res, 0, 6)/1000, 0)
			exif.DateShot = &date
			exposure := float64(DataApi.GetLong(res, 0, 7))
			exif.Exposure = &exposure
			aperture := float64(DataApi.GetLong(res, 0, 8))
			exif.Aperture = &aperture
			exif.Iso = DataApi.GetLongP(res, 0, 9)
			focalLength := float64(DataApi.GetLong(res, 0, 10))
			exif.FocalLength = &focalLength
			exif.Flash = DataApi.GetLongP(res, 0, 11)
			exif.Orientation = DataApi.GetLongP(res, 0, 12)
			exif.ExposureProgram = DataApi.GetLongP(res, 0, 13)
			gpsLaitude := float64(DataApi.GetLong(res, 0, 14))
			exif.GPSLatitude = &gpsLaitude
			gpsLongtude := float64(DataApi.GetLong(res, 0, 15))
			exif.GPSLongitude = &gpsLongtude
			exif.Description = DataApi.GetStringP(res, 0, 16)
			return &exif, nil
		}
	}

	if globalExifParser == nil {
		return nil, errors.New("No exif parser initialized")
	}

	exif, err := globalExifParser.ParseExif(media.Path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse exif data")
	}

	if exif == nil {
		return nil, nil
	}

	// Add EXIF to database and link to media
	if err := tx.Model(&media).Association("Exif").Replace(exif); err != nil {
		return nil, errors.Wrap(err, "save media exif to database")
	}

	if exif.DateShot != nil && !exif.DateShot.Equal(media.DateShot) {
		media.DateShot = *exif.DateShot
		if err := tx.Save(media).Error; err != nil {
			return nil, errors.Wrap(err, "update media date_shot")
		}
	}

	return exif, nil
}
