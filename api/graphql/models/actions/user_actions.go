package actions

import (
	"errors"
	"os"
	"path"
	"strconv"

	"github.com/photoview/photoview/api/graphql/models"
	"github.com/photoview/photoview/api/utils"
	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB, userID int) (*models.User, error) {

	// make sure the last admin user is not deleted
	var adminUsers []*models.User
	db.Model(&models.User{}).Where("admin = true").Limit(2).Find(&adminUsers) //SELECT * FROM `users` WHERE admin = true LIMIT 2
	if len(adminUsers) == 1 && adminUsers[0].ID == userID {
		return nil, errors.New("deleting sole admin user is not allowed")
	}

	var user models.User
	deletedAlbumIDs := make([]int, 0)

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, userID).Error; err != nil { //SELECT * FROM `users` WHERE `users`.`id` = 12 ORDER BY `users`.`id` LIMIT 1
			return err
		}

		userAlbums := user.Albums
		if err := tx.Model(&user).Association("Albums").Find(&userAlbums); err != nil { //SELECT `albums`.`id`,`albums`.`created_at`,`albums`.`updated_at`,`albums`.`title`,`albums`.`parent_album_id`,`albums`.`path`,`albums`.`path_hash`,`albums`.`cover_id` FROM `albums` JOIN `user_albums` ON `user_albums`.`album_id` = `albums`.`id` AND `user_albums`.`user_id` = 12
			return err
		}

		if err := tx.Model(&user).Association("Albums").Clear(); err != nil { // DELETE FROM `user_albums` WHERE `user_albums`.`user_id` = 12
			return err
		}

		for _, album := range userAlbums {
			var associatedUsers = tx.Model(album).Association("Owners").Count() //SELECT count(*) FROM `users` JOIN `user_albums` ON `user_albums`.`user_id` = `users`.`id` AND `user_albums`.`album_id` = 108

			if associatedUsers == 0 {
				deletedAlbumIDs = append(deletedAlbumIDs, album.ID)
				if err := tx.Delete(album).Error; err != nil {
					return err
				}
			}
		}

		if err := tx.Delete(&user).Error; err != nil { // DELETE FROM `users` WHERE `users`.`id` = 12

			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// If there is only one associated user, clean up the cache folder and delete the album row
	for _, deletedAlbumID := range deletedAlbumIDs {
		cachePath := path.Join(utils.MediaCachePath(), strconv.Itoa(int(deletedAlbumID)))
		if err := os.RemoveAll(cachePath); err != nil {
			return &user, err
		}
	}
	return &user, nil
}
