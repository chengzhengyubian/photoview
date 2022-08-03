package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	DataApi "github.com/photoview/photoview/api/dataapi"
	"time"

	"gorm.io/gorm"
)

type Album struct {
	Model
	Title         string `gorm:"not null"`
	ParentAlbumID *int   `gorm:"index"`
	ParentAlbum   *Album `gorm:"constraint:OnDelete:SET NULL;"`
	// OwnerID       int `gorm:"not null"`
	// Owner         User
	Owners   []User `gorm:"many2many:user_albums;constraint:OnDelete:CASCADE;"`
	Path     string `gorm:"not null"`
	PathHash string `gorm:"unique"`
	CoverID  *int
}

func (a *Album) FilePath() string {
	return a.Path
}

func (a *Album) BeforeSave(tx *gorm.DB) (err error) {
	hash := md5.Sum([]byte(a.Path))
	a.PathHash = hex.EncodeToString(hash[:])
	return nil
}

// GetChildren performs a recursive query to get all the children of the album.
// An optional filter can be provided that can be used to modify the query on the children.
func (a *Album) GetChildren(db *gorm.DB, filter func(*gorm.DB) *gorm.DB) (children []*Album, err error) {
	return GetChildrenFromAlbums(db, filter, []int{a.ID})
}

//修改中，递归有问题
func GetChildrenFromAlbums(db *gorm.DB, filter func(*gorm.DB) *gorm.DB, albumIDs []int) (children []*Album, err error) {
	query := db.Model(&Album{}).Table("sub_albums")
	//标记一下，应该是递归
	if filter != nil {
		query = filter(query)
	}

	err = db.Raw(`
	WITH recursive sub_albums AS (
		SELECT * FROM albums AS root WHERE id IN (?)
		UNION ALL
		SELECT child.* FROM albums AS child JOIN sub_albums ON child.parent_album_id = sub_albums.id
	)

	?
	`, albumIDs, query).Find(&children).Error
	/* WITH recursive sub_albums AS (
	           SELECT * FROM albums AS root WHERE id IN (110)
	           UNION ALL
	           SELECT child.* FROM albums AS child JOIN sub_albums ON child.parent_album_id = sub_albums.id
	   )

	   SELECT * FROM `sub_albums`
	*/
	return children, err
}

func (a *Album) GetParents(db *gorm.DB, filter func(*gorm.DB) *gorm.DB) (parents []*Album, err error) {
	return GetParentsFromAlbums(db, filter, a.ID)
}

//未修改，递归有问题
func GetParentsFromAlbums(db *gorm.DB, filter func(*gorm.DB) *gorm.DB, albumID int) (parents []*Album, err error) {
	query := db.Model(&Album{}).Table("super_albums")

	if filter != nil {
		query = filter(query)
	}

	err = db.Raw(`
	WITH recursive super_albums AS (
		SELECT * FROM albums AS leaf WHERE id = ?
		UNION ALL
		SELECT parent.* from albums AS parent JOIN super_albums ON parent.id = super_albums.parent_album_id
	)

	?
	`, albumID, query).Find(&parents).Error
	/*WITH recursive super_albums AS (
	          SELECT * FROM albums AS leaf WHERE id = 1
	          UNION ALL
	          SELECT parent.* from albums AS parent JOIN super_albums ON parent.id = super_albums.parent_album_id
	  )

	  SELECT * FROM `super_albums` WHERE id IN (1)*/

	return parents, err
}

//修改完，未测试
func (a *Album) Thumbnail(db *gorm.DB) (*Media, error) {
	var media Media

	if a.CoverID == nil {
		if err := db.Raw(`
			WITH recursive sub_albums AS (
				SELECT * FROM albums AS root WHERE id = ?
				UNION ALL
				SELECT child.* FROM albums AS child JOIN sub_albums ON child.parent_album_id = sub_albums.id
			)
		
			SELECT * FROM media WHERE media.album_id IN (
				SELECT id FROM sub_albums
			) AND media.id IN (
				SELECT media_id FROM media_urls WHERE media_urls.media_id = media.id
			) ORDER BY id LIMIT 1
		`, a.ID).Find(&media).Error; err != nil {
			return nil, err
		}
		/* WITH recursive sub_albums AS (
		           SELECT * FROM albums AS root WHERE id = 1
		           UNION ALL
		           SELECT child.* FROM albums AS child JOIN sub_albums ON child.parent_album_id = sub_albums.id
		   )

		   SELECT * FROM media WHERE media.album_id IN (
		           SELECT id FROM sub_albums
		   ) AND media.id IN (
		           SELECT media_id FROM media_urls WHERE media_urls.media_id = media.id
		   ) ORDER BY id LIMIT 1
		*/
		//sql_media_se := fmt.Sprintf("WITH recursive sub_albums AS (SELECT * FROM albums AS root WHERE id = %v UNION ALL SELECT child.* FROM albums AS child JOIN sub_albums ON child.parent_album_id = sub_albums.id) SELECT * FROM media WHERE media.album_id IN (SELECT id FROM sub_albums) AND media.id IN (SELECT media_id FROM media_urls WHERE media_urls.media_id = media.id) ORDER BY id LIMIT 1", a.ID)
		//dataApi, _ := DataApi.NewDataApiClient()
		//res, err := dataApi.Query(sql_media_se)
		//media.ID = DataApi.GetInt(res, 0, 0)
		//media.CreatedAt = time.Unix(*res[0][1].LongValue/1000, 0)
		//media.UpdatedAt = time.Unix(*res[0][2].LongValue/1000, 0)
		//media.Title = *res[0][3].StringValue
		//media.Path = *res[0][4].StringValue
		//media.PathHash = *res[0][5].StringValue
		//media.AlbumID = int(*res[0][6].LongValue)
		//media.ExifID = DataApi.GetIntP(res, 0, 7)
		//media.DateShot = time.Unix(*res[0][8].LongValue/1000, 0)
		//if *res[0][9].StringValue == "photo" {
		//	media.Type = MediaTypePhoto
		//} else {
		//	media.Type = MediaTypeVideo
		//}
		//media.VideoMetadataID = DataApi.GetIntP(res, 0, 10)
		//media.SideCarPath = DataApi.GetStringP(res, 0, 11)
		//media.SideCarHash = DataApi.GetStringP(res, 0, 12)
		//media.Blurhash = DataApi.GetStringP(res, 0, 13)
		//if err != nil {
		//	return nil, err
		//}
	} else {
		if err := db.Where("id = ?", a.CoverID).Find(&media).Error; err != nil {
			return nil, err
		}
		sql_media_se := fmt.Sprintf("select * from media where id=%v", a.CoverID)
		dataApi, _ := DataApi.NewDataApiClient()
		res, err := dataApi.Query(sql_media_se)
		media.ID = DataApi.GetInt(res, 0, 0)
		media.CreatedAt = time.Unix(*res[0][1].LongValue/1000, 0)
		media.UpdatedAt = time.Unix(*res[0][2].LongValue/1000, 0)
		media.Title = *res[0][3].StringValue
		media.Path = *res[0][4].StringValue
		media.PathHash = *res[0][5].StringValue
		media.AlbumID = int(*res[0][6].LongValue)
		media.ExifID = DataApi.GetIntP(res, 0, 7)
		media.DateShot = time.Unix(*res[0][8].LongValue/1000, 0)
		if *res[0][9].StringValue == "photo" {
			media.Type = MediaTypePhoto
		} else {
			media.Type = MediaTypeVideo
		}
		media.VideoMetadataID = DataApi.GetIntP(res, 0, 10)
		media.SideCarPath = DataApi.GetStringP(res, 0, 11)
		media.SideCarHash = DataApi.GetStringP(res, 0, 12)
		media.Blurhash = DataApi.GetStringP(res, 0, 13)
		if err != nil {
			return nil, err
		}
	}

	return &media, nil
}
