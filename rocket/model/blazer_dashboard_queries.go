package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `blazer_dashboard_queries` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `dashboard_id` bigint DEFAULT NULL,
  `query_id` bigint DEFAULT NULL,
  `position` int DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_blazer_dashboard_queries_on_dashboard_id` (`dashboard_id`),
  KEY `index_blazer_dashboard_queries_on_query_id` (`query_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "id": 16,    "dashboard_id": 25,    "query_id": 25,    "position": 33,    "created_at": "2083-06-27T00:34:44.350960304-04:00",    "updated_at": "2108-09-07T00:04:05.593812459-04:00"}



*/

// BlazerDashboardQueries struct is a row record of the blazer_dashboard_queries table in the rocket_development database
type BlazerDashboardQueries struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] dashboard_id                                   bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	DashboardID int `gorm:"column:dashboard_id;type:bigint;" json:"dashboard_id"`
	//[ 2] query_id                                       bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	QueryID int `gorm:"column:query_id;type:bigint;" json:"query_id"`
	//[ 3] position                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Position int `gorm:"column:position;type:int;" json:"position"`
	//[ 4] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 5] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var blazer_dashboard_queriesTableInfo = &TableInfo{
	Name: "blazer_dashboard_queries",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "dashboard_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "DashboardID",
			GoFieldType:        "int",
			JSONFieldName:      "dashboard_id",
			ProtobufFieldName:  "dashboard_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "query_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "QueryID",
			GoFieldType:        "int",
			JSONFieldName:      "query_id",
			ProtobufFieldName:  "query_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "position",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Position",
			GoFieldType:        "int",
			JSONFieldName:      "position",
			ProtobufFieldName:  "position",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BlazerDashboardQueries) TableName() string {
	return "blazer_dashboard_queries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BlazerDashboardQueries) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BlazerDashboardQueries) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BlazerDashboardQueries) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BlazerDashboardQueries) TableInfo() *TableInfo {
	return blazer_dashboard_queriesTableInfo
}
