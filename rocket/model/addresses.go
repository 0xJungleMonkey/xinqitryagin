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


CREATE TABLE `addresses` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `address_type` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `entity` varchar(255) DEFAULT NULL,
  `number_and_street` varchar(255) DEFAULT NULL,
  `suite_or_apartment` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `postal_code` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `notes` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `latitude` float DEFAULT NULL,
  `longitude` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "id": 96,    "address_type": "XtbjsoJqhOPgJtDotUYyCyKns",    "status": "YqMSjyhePLAKZDKhqyoOkGHgX",    "entity": "fOtadXcFSWHjadhWxjKnJHgAs",    "number_and_street": "IAPaBIQlvFSqrACUUYWIgWtRQ",    "suite_or_apartment": "KSHtEgsQHIYOTgnUhZIZnphFR",    "city": "qiqHeQiMTNSDtQJKmflPxixHo",    "postal_code": "PCmDZsigqEUxZGKtLieCSNwgf",    "country": "WThCIZKTQapjTYHaOtEbPqHSh",    "notes": "TBkvFawNrvLMcttmefjlTrmMk",    "created_at": "2169-10-07T03:47:40.455929199-04:00",    "updated_at": "2276-12-02T10:28:51.918914572-05:00",    "latitude": 0.30996025,    "longitude": 0.40509677}



*/

// Addresses struct is a row record of the addresses table in the rocket_development database
type Addresses struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] address_type                                   varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AddressType string `gorm:"column:address_type;type:varchar;size:255;" json:"address_type"`
	//[ 2] status                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Status string `gorm:"column:status;type:varchar;size:255;" json:"status"`
	//[ 3] entity                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Entity string `gorm:"column:entity;type:varchar;size:255;" json:"entity"`
	//[ 4] number_and_street                              varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberAndStreet string `gorm:"column:number_and_street;type:varchar;size:255;" json:"number_and_street"`
	//[ 5] suite_or_apartment                             varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SuiteOrApartment string `gorm:"column:suite_or_apartment;type:varchar;size:255;" json:"suite_or_apartment"`
	//[ 6] city                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	City string `gorm:"column:city;type:varchar;size:255;" json:"city"`
	//[ 7] postal_code                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	PostalCode string `gorm:"column:postal_code;type:varchar;size:255;" json:"postal_code"`
	//[ 8] country                                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Country string `gorm:"column:country;type:varchar;size:255;" json:"country"`
	//[ 9] notes                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Notes string `gorm:"column:notes;type:varchar;size:255;" json:"notes"`
	//[10] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[11] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	//[12] latitude                                       float                null: true   primary: false  isArray: false  auto: false  col: float           len: -1      default: []
	Latitude float64 `gorm:"column:latitude;type:float;" json:"latitude"`
	//[13] longitude                                      float                null: true   primary: false  isArray: false  auto: false  col: float           len: -1      default: []
	Longitude float64 `gorm:"column:longitude;type:float;" json:"longitude"`
}

var addressesTableInfo = &TableInfo{
	Name: "addresses",
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
			Name:               "address_type",
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
			GoFieldName:        "AddressType",
			GoFieldType:        "string",
			JSONFieldName:      "address_type",
			ProtobufFieldName:  "address_type",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "status",
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "entity",
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
			GoFieldName:        "Entity",
			GoFieldType:        "string",
			JSONFieldName:      "entity",
			ProtobufFieldName:  "entity",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "number_and_street",
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
			GoFieldName:        "NumberAndStreet",
			GoFieldType:        "string",
			JSONFieldName:      "number_and_street",
			ProtobufFieldName:  "number_and_street",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "suite_or_apartment",
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
			GoFieldName:        "SuiteOrApartment",
			GoFieldType:        "string",
			JSONFieldName:      "suite_or_apartment",
			ProtobufFieldName:  "suite_or_apartment",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "city",
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
			GoFieldName:        "City",
			GoFieldType:        "string",
			JSONFieldName:      "city",
			ProtobufFieldName:  "city",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "postal_code",
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
			GoFieldName:        "PostalCode",
			GoFieldType:        "string",
			JSONFieldName:      "postal_code",
			ProtobufFieldName:  "postal_code",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "country",
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
			GoFieldName:        "Country",
			GoFieldType:        "string",
			JSONFieldName:      "country",
			ProtobufFieldName:  "country",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "notes",
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
			GoFieldName:        "Notes",
			GoFieldType:        "string",
			JSONFieldName:      "notes",
			ProtobufFieldName:  "notes",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "latitude",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "Latitude",
			GoFieldType:        "float64",
			JSONFieldName:      "latitude",
			ProtobufFieldName:  "latitude",
			ProtobufType:       "float",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "longitude",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "Longitude",
			GoFieldType:        "float64",
			JSONFieldName:      "longitude",
			ProtobufFieldName:  "longitude",
			ProtobufType:       "float",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Addresses) TableName() string {
	return "addresses"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Addresses) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Addresses) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Addresses) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Addresses) TableInfo() *TableInfo {
	return addressesTableInfo
}
