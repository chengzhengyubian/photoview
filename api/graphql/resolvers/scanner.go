package resolvers

import (
	"context"
	"time"

	"github.com/photoview/photoview/api/database/drivers"
	"github.com/photoview/photoview/api/graphql/models"
	"github.com/photoview/photoview/api/scanner/periodic_scanner"
	"github.com/photoview/photoview/api/scanner/scanner_queue"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (r *mutationResolver) ScanAll(ctx context.Context) (*models.ScannerResult, error) {
	err := scanner_queue.AddAllToQueue()
	if err != nil {
		return nil, err
	}

	startMessage := "Scanner started"

	return &models.ScannerResult{
		Finished: false,
		Success:  true,
		Message:  &startMessage,
	}, nil
}

//修改完，未测试
func (r *mutationResolver) ScanUser(ctx context.Context, userID int) (*models.ScannerResult, error) {
	var user models.User
	if err := r.DB(ctx).First(&user, userID).Error; err != nil { //SELECT * FROM `users` WHERE `users`.`id` = 2 ORDER BY `users`.`id` LIMIT 1
		return nil, errors.Wrap(err, "get user from database")
	}
	//sql_users_se := "SELECT * FROM `users` WHERE `users`.`id` = " + strconv.Itoa(userID) + " ORDER BY `users`.`id` LIMIT 1"
	//dataApi, _ := DataApi.NewDataApiClient()
	//res, err := dataApi.Query(sql_users_se)
	//if len(res) == 0 {
	//	return nil, errors.Wrap(err, "get user from database")
	//}
	//user.ID = int(*res[0][0].LongValue)
	//user.Username = *res[0][3].StringValue
	//user.Password = res[0][4].StringValue
	//user.Admin = *res[0][5].BooleanValue
	scanner_queue.AddUserToQueue(&user)

	startMessage := "Scanner started"
	return &models.ScannerResult{
		Finished: false,
		Success:  true,
		Message:  &startMessage,
	}, nil
}

func (r *mutationResolver) SetPeriodicScanInterval(ctx context.Context, interval int) (int, error) {
	db := r.DB(ctx)
	if interval < 0 {
		return 0, errors.New("interval must be 0 or above")
	}

	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&models.SiteInfo{}).Update("periodic_scan_interval", interval).Error; err != nil {
		return 0, err
	}

	var siteInfo models.SiteInfo
	if err := db.First(&siteInfo).Error; err != nil {
		return 0, err
	}

	periodic_scanner.ChangePeriodicScanInterval(time.Duration(siteInfo.PeriodicScanInterval) * time.Second)

	return siteInfo.PeriodicScanInterval, nil
}

func (r *mutationResolver) SetScannerConcurrentWorkers(ctx context.Context, workers int) (int, error) {
	db := r.DB(ctx)
	if workers < 1 {
		return 0, errors.New("concurrent workers must at least be 1")
	}

	if workers > 1 && drivers.DatabaseDriverFromEnv() == drivers.SQLITE {
		return 0, errors.New("multiple workers not supported for SQLite databases")
	}

	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&models.SiteInfo{}).Update("concurrent_workers", workers).Error; err != nil {
		return 0, err
	}

	var siteInfo models.SiteInfo
	if err := db.First(&siteInfo).Error; err != nil {
		return 0, err
	}

	scanner_queue.ChangeScannerConcurrentWorkers(siteInfo.ConcurrentWorkers)

	return siteInfo.ConcurrentWorkers, nil
}
