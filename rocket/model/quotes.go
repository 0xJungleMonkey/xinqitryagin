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


CREATE TABLE `quotes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `building_type` varchar(255) DEFAULT NULL,
  `service_quality` varchar(255) DEFAULT NULL,
  `number_of_apartments` varchar(255) DEFAULT NULL,
  `number_of_floors` varchar(255) DEFAULT NULL,
  `number_of_businesses` varchar(255) DEFAULT NULL,
  `number_of_basements` varchar(255) DEFAULT NULL,
  `number_of_parking` varchar(255) DEFAULT NULL,
  `number_of_cages` varchar(255) DEFAULT NULL,
  `number_of_occupants` varchar(255) DEFAULT NULL,
  `number_of_hours` varchar(255) DEFAULT NULL,
  `number_of_elevators_needed` varchar(255) DEFAULT NULL,
  `price_per_unit` varchar(255) DEFAULT NULL,
  `elevator_price` varchar(255) DEFAULT NULL,
  `installation_fee` varchar(255) DEFAULT NULL,
  `final_price` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `company_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `department` varchar(255) DEFAULT NULL,
  `project_name` varchar(255) DEFAULT NULL,
  `project_description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "id": 35,    "building_type": "JGpLdKsZKDeqVoqWoEdctvset",    "service_quality": "PMCYavAVWTDdamQMsNoJmVZYj",    "number_of_apartments": "DIxWOBbUIwORVdqOdPRCZFhMt",    "number_of_floors": "OpyQpeXQjXvajxuIYThFWSJsA",    "number_of_businesses": "mISLdGCjBetoYElnNDSReThwq",    "number_of_basements": "IKUpoQbLrRUihjrfqWQBqZlap",    "number_of_parking": "JLrlWUSHPylxJwNqWMZUeAxpq",    "number_of_cages": "AhrmILyCXjIThlOwDggGxDZCr",    "number_of_occupants": "MJAmKZftpVJSRAdOpsBHYJOGU",    "number_of_hours": "CDjqLPcZepeASLsxGMivxrQhU",    "number_of_elevators_needed": "iSjKExjevNEjIXShayTpaGIIQ",    "price_per_unit": "QpDRtIZNXLusbVvjpvHfuSkWR",    "elevator_price": "PmyjPAEPFsZifIXuSrbraTeQr",    "installation_fee": "mhfcSXYHCighHJVnwLEhURhBL",    "final_price": "jQksyBpPLnXVAFSHolQnKfWKZ",    "created_at": "2030-09-10T19:57:39.015851676-04:00",    "updated_at": "2196-10-07T13:12:26.311671055-04:00",    "name": "wOvyakfZpmpObTGVWaDIRtgUa",    "company_name": "jeRNMMBnvXsPAZLDnetFasWqA",    "email": "bTieDjseVwXxfkdatQMOxBiOq",    "phone": "TrEUlBUajilvDroykNIcCfdCN",    "department": "FpdVqpgUPrgivetletDpdKepb",    "project_name": "FOXcFdQyPtTRuyMdcDWjnUiPD",    "project_description": "nItGUtmJnwNZcpvDGVEjMmtxs"}



*/

// Quotes struct is a row record of the quotes table in the rocket_development database
type Quotes struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] building_type                                  varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	BuildingType string `gorm:"column:building_type;type:varchar;size:255;" json:"building_type"`
	//[ 2] service_quality                                varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ServiceQuality string `gorm:"column:service_quality;type:varchar;size:255;" json:"service_quality"`
	//[ 3] number_of_apartments                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfApartments string `gorm:"column:number_of_apartments;type:varchar;size:255;" json:"number_of_apartments"`
	//[ 4] number_of_floors                               varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfFloors string `gorm:"column:number_of_floors;type:varchar;size:255;" json:"number_of_floors"`
	//[ 5] number_of_businesses                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfBusinesses string `gorm:"column:number_of_businesses;type:varchar;size:255;" json:"number_of_businesses"`
	//[ 6] number_of_basements                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfBasements string `gorm:"column:number_of_basements;type:varchar;size:255;" json:"number_of_basements"`
	//[ 7] number_of_parking                              varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfParking string `gorm:"column:number_of_parking;type:varchar;size:255;" json:"number_of_parking"`
	//[ 8] number_of_cages                                varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfCages string `gorm:"column:number_of_cages;type:varchar;size:255;" json:"number_of_cages"`
	//[ 9] number_of_occupants                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfOccupants string `gorm:"column:number_of_occupants;type:varchar;size:255;" json:"number_of_occupants"`
	//[10] number_of_hours                                varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfHours string `gorm:"column:number_of_hours;type:varchar;size:255;" json:"number_of_hours"`
	//[11] number_of_elevators_needed                     varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfElevatorsNeeded string `gorm:"column:number_of_elevators_needed;type:varchar;size:255;" json:"number_of_elevators_needed"`
	//[12] price_per_unit                                 varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	PricePerUnit string `gorm:"column:price_per_unit;type:varchar;size:255;" json:"price_per_unit"`
	//[13] elevator_price                                 varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ElevatorPrice string `gorm:"column:elevator_price;type:varchar;size:255;" json:"elevator_price"`
	//[14] installation_fee                               varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	InstallationFee string `gorm:"column:installation_fee;type:varchar;size:255;" json:"installation_fee"`
	//[15] final_price                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FinalPrice string `gorm:"column:final_price;type:varchar;size:255;" json:"final_price"`
	//[16] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[17] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	//[18] name                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;" json:"name"`
	//[19] company_name                                   varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CompanyName string `gorm:"column:company_name;type:varchar;size:255;" json:"company_name"`
	//[20] email                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Email string `gorm:"column:email;type:varchar;size:255;" json:"email"`
	//[21] phone                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Phone string `gorm:"column:phone;type:varchar;size:255;" json:"phone"`
	//[22] department                                     varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Department string `gorm:"column:department;type:varchar;size:255;" json:"department"`
	//[23] project_name                                   varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProjectName string `gorm:"column:project_name;type:varchar;size:255;" json:"project_name"`
	//[24] project_description                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProjectDescription string `gorm:"column:project_description;type:varchar;size:255;" json:"project_description"`
}

var quotesTableInfo = &TableInfo{
	Name: "quotes",
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
			Name:               "building_type",
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
			GoFieldName:        "BuildingType",
			GoFieldType:        "string",
			JSONFieldName:      "building_type",
			ProtobufFieldName:  "building_type",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "service_quality",
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
			GoFieldName:        "ServiceQuality",
			GoFieldType:        "string",
			JSONFieldName:      "service_quality",
			ProtobufFieldName:  "service_quality",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "number_of_apartments",
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
			GoFieldName:        "NumberOfApartments",
			GoFieldType:        "string",
			JSONFieldName:      "number_of_apartments",
			ProtobufFieldName:  "number_of_apartments",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "number_of_floors",
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
			GoFieldName:        "NumberOfFloors",
			GoFieldType:        "string",
			JSONFieldName:      "number_of_floors",
			ProtobufFieldName:  "number_of_floors",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "number_of_businesses",
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
			GoFieldName:        "NumberOfBusinesses",
			GoFieldType:        "string",
			JSONFieldName:      "number_of_businesses",
			ProtobufFieldName:  "number_of_businesses",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "number_of_basements",
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
			GoFieldName:        "NumberOfBasements",
			GoFieldType:        "string",
			JSONFieldName:      "number_of_basements",
			ProtobufFieldName:  "number_of_basements",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "number_of_parking",
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
			GoFieldName:        "NumberOfParking",
			GoFieldType:        "string",
			JSONFieldName:      "number_of_parking",
			ProtobufFieldName:  "number_of_parking",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "number_of_cages",
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
			GoFieldName:        "NumberOfCages",
			GoFieldType:        "string",
			JSONFieldName:      "number_of_cages",
			ProtobufFieldName:  "number_of_cages",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "number_of_occupants",
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
			GoFieldName:        "NumberOfOccupants",
			GoFieldType:        "string",
			JSONFieldName:      "number_of_occupants",
			ProtobufFieldName:  "number_of_occupants",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "number_of_hours",
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
			GoFieldName:        "NumberOfHours",
			GoFieldType:        "string",
			JSONFieldName:      "number_of_hours",
			ProtobufFieldName:  "number_of_hours",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "number_of_elevators_needed",
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
			GoFieldName:        "NumberOfElevatorsNeeded",
			GoFieldType:        "string",
			JSONFieldName:      "number_of_elevators_needed",
			ProtobufFieldName:  "number_of_elevators_needed",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "price_per_unit",
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
			GoFieldName:        "PricePerUnit",
			GoFieldType:        "string",
			JSONFieldName:      "price_per_unit",
			ProtobufFieldName:  "price_per_unit",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "elevator_price",
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
			GoFieldName:        "ElevatorPrice",
			GoFieldType:        "string",
			JSONFieldName:      "elevator_price",
			ProtobufFieldName:  "elevator_price",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "installation_fee",
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
			GoFieldName:        "InstallationFee",
			GoFieldType:        "string",
			JSONFieldName:      "installation_fee",
			ProtobufFieldName:  "installation_fee",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "final_price",
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
			GoFieldName:        "FinalPrice",
			GoFieldType:        "string",
			JSONFieldName:      "final_price",
			ProtobufFieldName:  "final_price",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "company_name",
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
			GoFieldName:        "CompanyName",
			GoFieldType:        "string",
			JSONFieldName:      "company_name",
			ProtobufFieldName:  "company_name",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "email",
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
			GoFieldName:        "Email",
			GoFieldType:        "string",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "phone",
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
			GoFieldName:        "Phone",
			GoFieldType:        "string",
			JSONFieldName:      "phone",
			ProtobufFieldName:  "phone",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "department",
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
			GoFieldName:        "Department",
			GoFieldType:        "string",
			JSONFieldName:      "department",
			ProtobufFieldName:  "department",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "project_name",
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
			GoFieldName:        "ProjectName",
			GoFieldType:        "string",
			JSONFieldName:      "project_name",
			ProtobufFieldName:  "project_name",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "project_description",
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
			GoFieldName:        "ProjectDescription",
			GoFieldType:        "string",
			JSONFieldName:      "project_description",
			ProtobufFieldName:  "project_description",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *Quotes) TableName() string {
	return "quotes"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *Quotes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *Quotes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *Quotes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *Quotes) TableInfo() *TableInfo {
	return quotesTableInfo
}
