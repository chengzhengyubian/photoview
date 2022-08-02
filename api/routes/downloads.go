package routes

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/photoview/photoview/api/database/drivers"
	"github.com/photoview/photoview/api/graphql/models"
	"gorm.io/gorm"
)

func RegisterDownloadRoutes(db *gorm.DB, router *mux.Router) {
	router.HandleFunc("/album/{album_id}/{media_purpose}", func(w http.ResponseWriter, r *http.Request) {
		albumID := mux.Vars(r)["album_id"]
		mediaPurpose := mux.Vars(r)["media_purpose"]
		mediaPurposeList := strings.SplitN(mediaPurpose, ",", 10)

		var album models.Album
		if err := db.Find(&album, albumID).Error; err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404"))
			return
		}

		if success, response, status, err := authenticateAlbum(&album, db, r); !success {
			if err != nil {
				log.Printf("WARN: error authenticating album for download: %v\n", err)
			}
			w.WriteHeader(status)
			w.Write([]byte(response))
			return
		}

		var mediaWhereQuery string
		if drivers.POSTGRES.MatchDatabase(db) {
			mediaWhereQuery = "\"Media\".album_id = ?"
		} else {
			mediaWhereQuery = "Media.album_id = ?"
		}

		var mediaURLs []*models.MediaURL
		if err := db.Joins("Media").Where(mediaWhereQuery, album.ID).Where("media_urls.purpose IN (?)", mediaPurposeList).Find(&mediaURLs).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}
		/*SELECT `media_urls`.`id`,`media_urls`.`created_at`,`media_urls`.`updated_at`,`media_urls`.`media_id`,`media_urls`.`media_name`,`media_urls`.`width`,`media_urls`.`height`,`media_urls`.`purpose`,`media_urls`.`content_type`,`media_urls`.`file_size`,`Media`.`id` AS `Media__id`,`Media`.`created_at` AS `Media__created_at`,`Media`.`updated_at` AS `Media__updated_at`,`Media`.`title` AS `Media__title`,`Media`.`path` AS `Media__path`,`Media`.`path_hash` AS `Media__path_hash`,`Media`.`album_id` AS `Media__album_id`,`Media`.`exif_id` AS `Media__exif_id`,`Media`.`date_shot` AS `Media__date_shot`,`Media`.`type` AS `Media__type`,`Media`.`video_metadata_id` AS `Media__video_metadata_id`,`Media`.`side_car_path` AS `Media__side_car_path`,`Media`.`side_car_hash` AS `Media__side_car_hash`,`Media`.`blurhash` AS `Media__blurhash` FROM `media_urls` LEFT JOIN `media` `Media` ON `media_urls`.`media_id` = `Media`.`id` WHERE Media.album_id = 146 AND media_urls.purpose IN ('original')*/
		if len(mediaURLs) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no media found"))
			return
		}

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", album.Title))

		zipWriter := zip.NewWriter(w)

		for _, media := range mediaURLs {
			zipFile, err := zipWriter.Create(fmt.Sprintf("%s/%s", album.Title, media.MediaName))
			if err != nil {
				log.Printf("ERROR: Failed to create a file in zip, when downloading album (%d): %v\n", album.ID, err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error"))
				return
			}

			filePath, err := media.CachedPath()
			if err != nil {
				log.Printf("ERROR: Failed to get mediaURL cache path, when downloading album (%d): %v\n", album.ID, err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error"))
				return
			}

			fileData, err := os.Open(filePath)
			if err != nil {
				log.Printf("ERROR: Failed to open file to include in zip, when downloading album (%d): %v\n", album.ID, err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error"))
				return
			}

			_, err = io.Copy(zipFile, fileData)
			if err != nil {
				log.Printf("ERROR: Failed to copy file data, when downloading album (%d): %v\n", album.ID, err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error"))
				return
			}

			if err := fileData.Close(); err != nil {
				log.Printf("ERROR: Failed to close file, when downloading album (%d): %v\n", album.ID, err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error"))
				return
			}
		}

		// close the zip Writer to flush the contents to the ResponseWriter
		zipWriter.Close()
	})
}
