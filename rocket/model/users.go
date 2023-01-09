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


CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL DEFAULT '',
  `encrypted_password` varchar(255) NOT NULL DEFAULT '',
  `reset_password_token` varchar(255) DEFAULT NULL,
  `reset_password_sent_at` datetime DEFAULT NULL,
  `remember_created_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_users_on_email` (`email`),
  UNIQUE KEY `index_users_on_reset_password_token` (`reset_password_token`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "id": 1,    "email": "vjSdjoDgumqQlJlxFihQZGyPa",    "encrypted_password": "CZNSTCULPMEvUJEieihWrHUZI",    "reset_password_token": "fPUNOReaEhylwnpFDQkcGGqww",    "reset_password_sent_at": "2142-08-08T17:24:58.564625957-04:00",    "remember_created_at": "2064-03-20T01:57:59.366927069-04:00",    "created_at": "2153-07-25T07:50:09.284813174-04:00",    "updated_at": "2188-06-24T23:37:19.246481176-04:00"}



*/

// Users_ struct is a row record of the users table in the rocket_development database
type Users_ struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] email                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Email string `gorm:"column:email;type:varchar;size:255;" json:"email"`
	//[ 2] encrypted_password                             varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	EncryptedPassword string `gorm:"column:encrypted_password;type:varchar;size:255;" json:"encrypted_password"`
	//[ 3] reset_password_token                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ResetPasswordToken null.String `gorm:"column:reset_password_token;type:varchar;size:255;" json:"reset_password_token"`
	//[ 4] reset_password_sent_at                         datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	ResetPasswordSentAt null.Time `gorm:"column:reset_password_sent_at;type:datetime;" json:"reset_password_sent_at"`
	//[ 5] remember_created_at                            datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	RememberCreatedAt null.Time `gorm:"column:remember_created_at;type:datetime;" json:"remember_created_at"`
	//[ 6] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 7] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var usersTableInfo = &TableInfo{
	Name: "users",
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
			Name:               "email",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "encrypted_password",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "EncryptedPassword",
			GoFieldType:        "string",
			JSONFieldName:      "encrypted_password",
			ProtobufFieldName:  "encrypted_password",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "reset_password_token",
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
			GoFieldName:        "ResetPasswordToken",
			GoFieldType:        "null.String",
			JSONFieldName:      "reset_password_token",
			ProtobufFieldName:  "reset_password_token",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "reset_password_sent_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ResetPasswordSentAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "reset_password_sent_at",
			ProtobufFieldName:  "reset_password_sent_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "remember_created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "RememberCreatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "remember_created_at",
			ProtobufFieldName:  "remember_created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *Users_) TableName() string {
	return "users"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *Users_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *Users_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *Users_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *Users_) TableInfo() *TableInfo {
	return usersTableInfo
}
