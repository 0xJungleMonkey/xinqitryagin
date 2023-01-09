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


CREATE TABLE `customers` (
  `address_id` bigint DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  `CustomerCreationDate` varchar(255) DEFAULT NULL,
  `date` varchar(255) DEFAULT NULL,
  `CompanyName` varchar(255) DEFAULT NULL,
  `CompanyHQAdress` varchar(255) DEFAULT NULL,
  `FullNameOfCompanyContact` varchar(255) DEFAULT NULL,
  `CompanyContactPhone` varchar(255) DEFAULT NULL,
  `CompanyContactEMail` varchar(255) DEFAULT NULL,
  `CompanyDesc` text,
  `FullNameServiceTechAuth` varchar(255) DEFAULT NULL,
  `TechAuthPhoneService` varchar(255) DEFAULT NULL,
  `TechManagerEmailService` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_customers_on_user_id` (`user_id`),
  KEY `index_customers_on_address_id` (`address_id`),
  CONSTRAINT `fk_rails_3f9404ba26` FOREIGN KEY (`address_id`) REFERENCES `addresses` (`id`),
  CONSTRAINT `fk_rails_9917eeaf5d` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "address_id": 32,    "user_id": 90,    "id": 77,    "customer_creation_date": "mSlZoLXyCPbOcMHNMyvXhBEWL",    "date": "dhIFBXeAVrhvvAelKIBVBaOmK",    "company_name": "NiSOMtHhJDPnEdMyroTCrfyns",    "company_hq_adress": "kviOgeulFHgfhEfeZkxDAEMhx",    "full_name_of_company_contact": "imfnNivinjrcahIEEpUTqqthT",    "company_contact_phone": "yJGUeApGFoFwJtymoCucABvKa",    "company_contact_e_mail": "vvwfPNUxWeOZtHInJVKHdXPPa",    "company_desc": "ghaIJrMMpZmTaLMEmKdwlxNKl",    "full_name_service_tech_auth": "HVuBRXIJdFmtcTJBoBcIFcZsU",    "tech_auth_phone_service": "gVbTXttlQqlsOBlyQSJTVdgYa",    "tech_manager_email_service": "IAbXcKFJCJTCOEsswgmGvXeea",    "created_at": "2058-08-31T19:55:05.345388406-04:00",    "updated_at": "2143-01-10T23:08:02.834963619-05:00"}



*/

// Customers struct is a row record of the customers table in the rocket_development database
type Customers struct {
	//[ 0] address_id                                     bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	AddressID null.Int `gorm:"column:address_id;type:bigint;" json:"address_id"`
	//[ 1] user_id                                        bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	UserID null.Int `gorm:"column:user_id;type:bigint;" json:"user_id"`
	//[ 2] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 3] CustomerCreationDate                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CustomerCreationDate string `gorm:"column:CustomerCreationDate;type:varchar;size:255;" json:"customer_creation_date"`
	//[ 4] date                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Date string `gorm:"column:date;type:varchar;size:255;" json:"date"`
	//[ 5] CompanyName                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CompanyName string `gorm:"column:CompanyName;type:varchar;size:255;" json:"company_name"`
	//[ 6] CompanyHQAdress                                varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CompanyHQAdress string `gorm:"column:CompanyHQAdress;type:varchar;size:255;" json:"company_hq_adress"`
	//[ 7] FullNameOfCompanyContact                       varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FullNameOfCompanyContact string `gorm:"column:FullNameOfCompanyContact;type:varchar;size:255;" json:"full_name_of_company_contact"`
	//[ 8] CompanyContactPhone                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CompanyContactPhone string `gorm:"column:CompanyContactPhone;type:varchar;size:255;" json:"company_contact_phone"`
	//[ 9] CompanyContactEMail                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CompanyContactEMail string `gorm:"column:CompanyContactEMail;type:varchar;size:255;" json:"company_contact_e_mail"`
	//[10] CompanyDesc                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	CompanyDesc string `gorm:"column:CompanyDesc;type:text;size:65535;" json:"company_desc"`
	//[11] FullNameServiceTechAuth                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FullNameServiceTechAuth string `gorm:"column:FullNameServiceTechAuth;type:varchar;size:255;" json:"full_name_service_tech_auth"`
	//[12] TechAuthPhoneService                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	TechAuthPhoneService string `gorm:"column:TechAuthPhoneService;type:varchar;size:255;" json:"tech_auth_phone_service"`
	//[13] TechManagerEmailService                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	TechManagerEmailService string `gorm:"column:TechManagerEmailService;type:varchar;size:255;" json:"tech_manager_email_service"`
	//[14] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[15] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var customersTableInfo = &TableInfo{
	Name: "customers",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "address_id",
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
			GoFieldName:        "AddressID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "address_id",
			ProtobufFieldName:  "address_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "user_id",
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
			GoFieldName:        "UserID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "CustomerCreationDate",
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
			GoFieldName:        "CustomerCreationDate",
			GoFieldType:        "string",
			JSONFieldName:      "customer_creation_date",
			ProtobufFieldName:  "customer_creation_date",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "date",
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
			GoFieldName:        "Date",
			GoFieldType:        "string",
			JSONFieldName:      "date",
			ProtobufFieldName:  "date",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "CompanyName",
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CompanyHQAdress",
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
			GoFieldName:        "CompanyHQAdress",
			GoFieldType:        "string",
			JSONFieldName:      "company_hq_adress",
			ProtobufFieldName:  "company_hq_adress",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "FullNameOfCompanyContact",
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
			GoFieldName:        "FullNameOfCompanyContact",
			GoFieldType:        "string",
			JSONFieldName:      "full_name_of_company_contact",
			ProtobufFieldName:  "full_name_of_company_contact",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "CompanyContactPhone",
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
			GoFieldName:        "CompanyContactPhone",
			GoFieldType:        "string",
			JSONFieldName:      "company_contact_phone",
			ProtobufFieldName:  "company_contact_phone",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "CompanyContactEMail",
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
			GoFieldName:        "CompanyContactEMail",
			GoFieldType:        "string",
			JSONFieldName:      "company_contact_e_mail",
			ProtobufFieldName:  "company_contact_e_mail",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "CompanyDesc",
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
			GoFieldName:        "CompanyDesc",
			GoFieldType:        "string",
			JSONFieldName:      "company_desc",
			ProtobufFieldName:  "company_desc",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "FullNameServiceTechAuth",
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
			GoFieldName:        "FullNameServiceTechAuth",
			GoFieldType:        "string",
			JSONFieldName:      "full_name_service_tech_auth",
			ProtobufFieldName:  "full_name_service_tech_auth",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "TechAuthPhoneService",
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
			GoFieldName:        "TechAuthPhoneService",
			GoFieldType:        "string",
			JSONFieldName:      "tech_auth_phone_service",
			ProtobufFieldName:  "tech_auth_phone_service",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "TechManagerEmailService",
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
			GoFieldName:        "TechManagerEmailService",
			GoFieldType:        "string",
			JSONFieldName:      "tech_manager_email_service",
			ProtobufFieldName:  "tech_manager_email_service",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Customers) TableName() string {
	return "customers"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Customers) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Customers) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Customers) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Customers) TableInfo() *TableInfo {
	return customersTableInfo
}
