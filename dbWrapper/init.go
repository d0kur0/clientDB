package dbWrapper

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/d0kur0/clientDB/utils/configMgr"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (err error) {
	config := configMgr.Get()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.User,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Dbname,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err == nil {
		log.Println("Connected to db successful")
	}

	return errors.Wrap(err, "Connected to db successful")
}

func GetDB() *gorm.DB {
	return db
}

func Migrate() error {
	err := db.AutoMigrate(&User{})

	if err == nil {
		log.Println("Migrations successful ended")
	}

	return errors.Wrap(err, "Error with execute migrations")
}
