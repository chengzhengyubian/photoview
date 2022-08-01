package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/photoview/photoview/api/graphql/models"
	"github.com/photoview/photoview/api/scanner"
)

//修改中，未完
func RegisterPhotoRoutes(db *gorm.DB, router *mux.Router) {

	router.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		mediaName := mux.Vars(r)["name"]

		var mediaURL models.MediaURL
		result := db.Model(&models.MediaURL{}).Joins("Media").Select("media_urls.*").Where("media_urls.media_name = ?", mediaName).Scan(&mediaURL) //SELECT media_urls.*,`Media`.`id` AS `Media__id`,`Media`.`created_at` AS `Media__created_at`,`Media`.`updated_at` AS `Media__updated_at`,`Media`.`title` AS `Media__title`,`Media`.`path` AS `Media__path`,`Media`.`path_hash` AS `Media__path_hash`,`Media`.`album_id` AS `Media__album_id`,`Media`.`exif_id` AS `Media__exif_id`,`Media`.`date_shot` AS `Media__date_shot`,`Media`.`type` AS `Media__type`,`Media`.`video_metadata_id` AS `Media__video_metadata_id`,`Media`.`side_car_path` AS `Media__side_car_path`,`Media`.`side_car_hash` AS `Media__side_car_hash`,`Media`.`blurhash` AS `Media__blurhash` FROM `media_urls` LEFT JOIN `media` `Media` ON `media_urls`.`media_id` = `Media`.`id` WHERE media_urls.media_name = 'thumbnail_凡凡暑假生活_jpg_KwLL0ZoX.jpg'
		if err := result.Error; err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404"))
			return
		}
		//sql_media_urls_se := "SELECT media_urls.*,`Media`.`id` AS `Media__id`,`Media`.`created_at` AS `Media__created_at`,`Media`.`updated_at` AS `Media__updated_at`,`Media`.`title` AS `Media__title`,`Media`.`path` AS `Media__path`,`Media`.`path_hash` AS `Media__path_hash`,`Media`.`album_id` AS `Media__album_id`,`Media`.`exif_id` AS `Media__exif_id`,`Media`.`date_shot` AS `Media__date_shot`,`Media`.`type` AS `Media__type`,`Media`.`video_metadata_id` AS `Media__video_metadata_id`,`Media`.`side_car_path` AS `Media__side_car_path`,`Media`.`side_car_hash` AS `Media__side_car_hash`,`Media`.`blurhash` AS `Media__blurhash` FROM `media_urls` LEFT JOIN `media` `Media` ON `media_urls`.`media_id` = `Media`.`id` WHERE media_urls.media_name =\"" + mediaName + "\""
		//dataAPi, _ := DataApi.NewDataApiClient()
		//res, err := dataAPi.Query(sql_media_urls_se)
		//if len(res) == 0 {
		//	w.WriteHeader(http.StatusNotFound)
		//	w.Write([]byte("404"))
		//	return
		//}
		//mediaURL.ID = int(*res[0][0].LongValue)
		//mediaURL.CreatedAt = time.Unix(*res[0][1].LongValue/1000, 0)
		//mediaURL.UpdatedAt = time.Unix(*res[0][2].LongValue/1000, 0)
		//mediaURL.MediaID = int(*res[0][3].LongValue)
		//mediaURL.MediaName = *res[0][4].StringValue
		//mediaURL.Width = int(*res[0][5].LongValue)
		//mediaURL.Height = int(*res[0][6].LongValue)
		//switch *res[0][7].StringValue {
		//case "thumbnail":
		//	mediaURL.Purpose = models.PhotoThumbnail
		//case "high-res":
		//	mediaURL.Purpose = models.PhotoHighRes
		//case "original":
		//	mediaURL.Purpose = models.MediaOriginal
		//case "video-web":
		//	mediaURL.Purpose = models.VideoWeb
		//case "video-thumbnail":
		//	mediaURL.Purpose = models.VideoThumbnail
		//}
		//mediaURL.ContentType = *res[0][8].StringValue
		//mediaURL.FileSize = *res[0][9].LongValue

		media := mediaURL.Media //这里注意一下media有没有值，另外考虑吧下怎么赋值

		if success, response, status, err := authenticateMedia(media, db, r); !success {
			if err != nil {
				log.Printf("WARN: error authenticating photo: %s\n", err)
			}
			w.WriteHeader(status)
			w.Write([]byte(response))
			return
		}

		cachedPath, err := mediaURL.CachedPath()
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		if _, err := os.Stat(cachedPath); os.IsNotExist((err)) {
			err := db.Transaction(func(tx *gorm.DB) error {
				if err = scanner.ProcessSingleMedia(tx, media); err != nil {
					log.Printf("ERROR: processing image not found in cache (%s): %s\n", cachedPath, err)
					return err
				}

				if _, err = os.Stat(cachedPath); err != nil {
					log.Printf("ERROR: after reprocessing image not found in cache (%s): %s\n", cachedPath, err)
					return err
				}

				return nil
			})

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error"))
				return
			}
		}

		// Allow caching the resource for 1 day
		w.Header().Set("Cache-Control", "private, max-age=86400, immutable")

		http.ServeFile(w, r, cachedPath)
	})
}
