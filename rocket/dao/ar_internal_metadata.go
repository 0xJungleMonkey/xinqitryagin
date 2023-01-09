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

// GetAllArInternalMetadata is a function to get a slice of record(s) from ar_internal_metadata table in the rocket_development database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllArInternalMetadata(ctx context.Context, page, pagesize int64, order string) (results []*model.ArInternalMetadata, totalRows int, err error) {

	resultOrm := DB.Model(&model.ArInternalMetadata{})
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

// GetArInternalMetadata is a function to get a single record from the ar_internal_metadata table in the rocket_development database
// error - ErrNotFound, db Find error
func GetArInternalMetadata(ctx context.Context, argKey string) (record *model.ArInternalMetadata, err error) {
	record = &model.ArInternalMetadata{}
	if err = DB.First(record, argKey).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddArInternalMetadata is a function to add a single record to ar_internal_metadata table in the rocket_development database
// error - ErrInsertFailed, db save call failed
func AddArInternalMetadata(ctx context.Context, record *model.ArInternalMetadata) (result *model.ArInternalMetadata, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateArInternalMetadata is a function to update a single record from ar_internal_metadata table in the rocket_development database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateArInternalMetadata(ctx context.Context, argKey string, updated *model.ArInternalMetadata) (result *model.ArInternalMetadata, RowsAffected int64, err error) {

	result = &model.ArInternalMetadata{}
	db := DB.First(result, argKey)
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

// DeleteArInternalMetadata is a function to delete a single record from ar_internal_metadata table in the rocket_development database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteArInternalMetadata(ctx context.Context, argKey string) (rowsAffected int64, err error) {

	record := &model.ArInternalMetadata{}
	db := DB.First(record, argKey)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
