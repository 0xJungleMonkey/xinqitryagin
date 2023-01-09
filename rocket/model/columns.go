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


CREATE TABLE `columns` (
  `battery_id` bigint DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  `Type` varchar(255) DEFAULT NULL,
  `NumOfFloorsServed` int DEFAULT NULL,
  `Status` varchar(255) DEFAULT NULL,
  `Information` text,
  `Notes` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_columns_on_battery_id` (`battery_id`),
  CONSTRAINT `fk_rails_021eb14ac4` FOREIGN KEY (`battery_id`) REFERENCES `batteries` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=197 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "battery_id": 86,    "id": 97,    "type": "PxrsYgZIcgcTxeXDvkUlLXTuL",    "num_of_floors_served": 92,    "status": "GjkFXAlWapXkuopYwXQpwyyJs",    "information": "fJsNjkFIvMfvhSUkXeJVjmZdU",    "notes": "GbTYAlQKWShJIMeNnmsHnMjrV",    "created_at": "2234-12-01T20:46:52.822820591-05:00",    "updated_at": "2129-10-11T00:57:53.705769736-04:00"}



*/

// Columns struct is a row record of the columns table in the rocket_development database
type Columns struct {
	//[ 0] battery_id                                     bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	BatteryID int64 `gorm:"column:battery_id;type:bigint;" json:"battery_id"`
	//[ 1] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 2] Type                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Type string `gorm:"column:Type;type:varchar;size:255;" json:"type"`
	//[ 3] NumOfFloorsServed                              int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	NumOfFloorsServed int64 `gorm:"column:NumOfFloorsServed;type:int;" json:"num_of_floors_served"`
	//[ 4] Status                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Status string `gorm:"column:Status;type:varchar;size:255;" json:"status"`
	//[ 5] Information                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Information string `gorm:"column:Information;type:text;size:65535;" json:"information"`
	//[ 6] Notes                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Notes string `gorm:"column:Notes;type:text;size:65535;" json:"notes"`
	//[ 7] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 8] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var columnsTableInfo = &TableInfo{
	Name: "columns",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "battery_id",
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
			GoFieldName:        "BatteryID",
			GoFieldType:        "int64",
			JSONFieldName:      "battery_id",
			ProtobufFieldName:  "battery_id",
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "NumOfFloorsServed",
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
			GoFieldName:        "NumOfFloorsServed",
			GoFieldType:        "int64",
			JSONFieldName:      "num_of_floors_served",
			ProtobufFieldName:  "num_of_floors_served",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Columns) TableName() string {
	return "columns"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Columns) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Columns) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Columns) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Columns) TableInfo() *TableInfo {
	return columnsTableInfo
}
