package dao

import (
	"context"
	"time"

	"rocket/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllBlazerDashboards is a function to get a slice of record(s) from blazer_dashboards table in the rocket_development database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBlazerDashboards(ctx context.Context, page, pagesize int64, order string) (results []*model.BlazerDashboards, totalRows int, err error) {

	resultOrm := DB.Model(&model.BlazerDashboards{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetBlazerDashboards is a function to get a single record from the blazer_dashboards table in the rocket_development database
// error - ErrNotFound, db Find error
func GetBlazerDashboards(ctx context.Context, argID int64) (record *model.BlazerDashboards, err error) {
	record = &model.BlazerDashboards{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddBlazerDashboards is a function to add a single record to blazer_dashboards table in the rocket_development database
// error - ErrInsertFailed, db save call failed
func AddBlazerDashboards(ctx context.Context, record *model.BlazerDashboards) (result *model.BlazerDashboards, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateBlazerDashboards is a function to update a single record from blazer_dashboards table in the rocket_development database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBlazerDashboards(ctx context.Context, argID int64, updated *model.BlazerDashboards) (result *model.BlazerDashboards, RowsAffected int64, err error) {

	result = &model.BlazerDashboards{}
	db := DB.First(result, argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteBlazerDashboards is a function to delete a single record from blazer_dashboards table in the rocket_development database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBlazerDashboards(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &model.BlazerDashboards{}
	db := DB.First(record, argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
