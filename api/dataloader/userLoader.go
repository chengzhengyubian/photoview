package dataloader

import (
	"time"

	"github.com/photoview/photoview/api/graphql/models"
	"gorm.io/gorm"
)

func NewUserLoaderByToken(db *gorm.DB) *UserLoader {
	return &UserLoader{
		maxBatch: 100,
		wait:     5 * time.Millisecond,
		fetch: func(tokens []string) ([]*models.User, []error) {

			var accessTokens []*models.AccessToken
			err := db.Where("expire > ?", time.Now()).Where("value IN (?)", tokens).Find(&accessTokens).Error
			//SELECT * FROM `access_tokens` WHERE expire > '2022-07-28 00:24:13.606' AND value IN ('WohDNhR2teZ344ldk4jkJyQL')
			if err != nil {
				return nil, []error{err}
			}
			//SELECT distinct user_id FROM `access_tokens` WHERE expire > '2022-07-28 00:24:13.682' AND value IN ('WohDNhR2teZ344ldk4jkJyQL')
			rows, err := db.Table("access_tokens").Select("distinct user_id").Where("expire > ?", time.Now()).Where("value IN (?)", tokens).Rows()
			if err != nil {
				return nil, []error{err}
			}
			userIDs := make([]int, 0)
			for rows.Next() {
				var id int
				if err := db.ScanRows(rows, &id); err != nil {
					return nil, []error{err}
				}
				userIDs = append(userIDs, id)
			}
			rows.Close()

			var userMap map[int]*models.User
			if len(userIDs) > 0 {

				var users []*models.User
				if err := db.Where("id IN (?)", userIDs).Find(&users).Error; err != nil { // SELECT * FROM `users` WHERE id IN (2)
					return nil, []error{err}
				}

				userMap = make(map[int]*models.User, len(users))
				for _, user := range users {
					userMap[user.ID] = user
				}
			} else {
				userMap = make(map[int]*models.User, 0)
			}

			tokenMap := make(map[string]*models.AccessToken, len(tokens))
			for _, token := range accessTokens {
				tokenMap[token.Value] = token
			}

			result := make([]*models.User, len(tokens))
			for i, token := range tokens {
				accessToken, tokenFound := tokenMap[token]
				if tokenFound {
					user, userFound := userMap[accessToken.UserID]
					if userFound {
						result[i] = user
					}
				}
			}

			return result, nil
		},
	}
}
