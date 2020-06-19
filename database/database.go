package database

import (
	"fmt"
	"strings"

	"github.com/conflux-tech/fiber-rest-boilerplate/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Instance function creates a new DB instance
func Instance() *gorm.DB {
	return db
}

// Connect function connects to the database
func Connect(conf *config.DBConfig) {
	var err error
	switch strings.ToLower(conf.Driver) {
	case "pg":
		conn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v", conf.Host, conf.Port, conf.User, conf.Database, conf.Password)
		db, err = gorm.Open("postgres", conn)
	case "mysql", "mariadb":
		conn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Database)
		db, err = gorm.Open("mysql", conn)
		if err == nil {
			db.Set("gorm:table_options", "Engine=InnoDB")
		}
	case "sqlite", "sqlite3":
		db, err = gorm.Open("sqlite3", conf.Database)
	}
	if err != nil {
		panic("Failed to connect to database")
	}
}

// Close function closes the database connection
func Close() error {
	return db.Close()
}
