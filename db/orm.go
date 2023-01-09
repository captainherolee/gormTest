package db

import (
	"gormTest/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	//db, err := gorm.Open(dbname, con+"?parseTime=true")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      con + "?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:        256,                                            // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision: true,                                           // disable datetime precision support, which not supported before MySQL 5.6

		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{})
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetUser(email string) (user models.User, err error) {
	return user, db.Where(&models.User{Email: email}).Find(&user).Error
}
