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


CREATE TABLE `elevators` (
  `column_id` bigint DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  `SerialNumber` int DEFAULT NULL,
  `Model` varchar(255) DEFAULT NULL,
  `Type` varchar(255) DEFAULT NULL,
  `Status` varchar(255) DEFAULT NULL,
  `CommisionDate` date DEFAULT NULL,
  `LastInspectionDate` date DEFAULT NULL,
  `InspectionCert` varchar(255) DEFAULT NULL,
  `Information` text,
  `Notes` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_elevators_on_column_id` (`column_id`),
  CONSTRAINT `fk_rails_69442d7bc2` FOREIGN KEY (`column_id`) REFERENCES `columns` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=592 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "column_id": 60,    "id": 83,    "serial_number": 0,    "model": "tsxyNuGukqybVroZfYRvoIZrv",    "type": "nMecqycwwRvxaqEyfhBPYVJIs",    "status": "wFfpFIYfvsDvtOwNSDSTpQjgg",    "commision_date": "2122-04-04T08:37:18.972803188-04:00",    "last_inspection_date": "2226-10-03T11:12:08.69314046-04:00",    "inspection_cert": "iMoxfiCQvDSPBlWXNJechfsOq",    "information": "rrRKhblwnAQpjkuZjfZecYlJc",    "notes": "iwMLRUvRKyVrCdWRpCsvAmBLZ",    "created_at": "2071-09-21T15:19:23.652481199-04:00",    "updated_at": "2126-06-30T04:24:39.760038557-04:00"}



*/

// Elevators struct is a row record of the elevators table in the rocket_development database
type Elevators struct {
	//[ 0] column_id                                      bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	ColumnID null.Int `gorm:"column:column_id;type:bigint;" json:"column_id"`
	//[ 1] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 2] SerialNumber                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SerialNumber null.Int `gorm:"column:SerialNumber;type:int;" json:"serial_number"`
	//[ 3] Model                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Model string `gorm:"column:Model;type:varchar;size:255;" json:"model"`
	//[ 4] Type                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Type string `gorm:"column:Type;type:varchar;size:255;" json:"type"`
	//[ 5] Status                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Status string `gorm:"column:Status;type:varchar;size:255;" json:"status"`
	//[ 6] CommisionDate                                  date                 null: true   primary: false  isArray: false  auto: false  col: date            len: -1      default: []
	CommisionDate null.Time `gorm:"column:CommisionDate;type:date;" json:"commision_date"`
	//[ 7] LastInspectionDate                             date                 null: true   primary: false  isArray: false  auto: false  col: date            len: -1      default: []
	LastInspectionDate null.Time `gorm:"column:LastInspectionDate;type:date;" json:"last_inspection_date"`
	//[ 8] InspectionCert                                 varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	InspectionCert string `gorm:"column:InspectionCert;type:varchar;size:255;" json:"inspection_cert"`
	//[ 9] Information                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Information string `gorm:"column:Information;type:text;size:65535;" json:"information"`
	//[10] Notes                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Notes string `gorm:"column:Notes;type:text;size:65535;" json:"notes"`
	//[11] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[12] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var elevatorsTableInfo = &TableInfo{
	Name: "elevators",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "column_id",
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
			GoFieldName:        "ColumnID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "column_id",
			ProtobufFieldName:  "column_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "SerialNumber",
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
			GoFieldName:        "SerialNumber",
			GoFieldType:        "null.Int",
			JSONFieldName:      "serial_number",
			ProtobufFieldName:  "serial_number",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "Model",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Model",
			GoFieldType:        "string",
			JSONFieldName:      "model",
			ProtobufFieldName:  "model",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "Type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Type",
			GoFieldType:        "string",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "Status",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Status",
			GoFieldType:        "string",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CommisionDate",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "CommisionDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "commision_date",
			ProtobufFieldName:  "commision_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "LastInspectionDate",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "LastInspectionDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "last_inspection_date",
			ProtobufFieldName:  "last_inspection_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "InspectionCert",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "InspectionCert",
			GoFieldType:        "string",
			JSONFieldName:      "inspection_cert",
			ProtobufFieldName:  "inspection_cert",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "Information",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Information",
			GoFieldType:        "string",
			JSONFieldName:      "information",
			ProtobufFieldName:  "information",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "Notes",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Notes",
			GoFieldType:        "string",
			JSONFieldName:      "notes",
			ProtobufFieldName:  "notes",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Elevators) TableName() string {
	return "elevators"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Elevators) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Elevators) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Elevators) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Elevators) TableInfo() *TableInfo {
	return elevatorsTableInfo
}
