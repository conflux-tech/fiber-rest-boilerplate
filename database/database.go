package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/conflux-tech/fiber-rest-boilerplate/configs"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Instance function creates a new DB instance
func Instance() *gorm.DB {
	return db
}

// Connect function connects to the database
func Connect(conf *configs.DBConfig) {
	var err error
	switch strings.ToLower(conf.Driver) {
	case "pg", "postgres":
		dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v", conf.Host, conf.Port, conf.User, conf.Database, conf.Password, conf.SSLMode)
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	case "mysql", "mariadb":
		dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Database)
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{})
		if err == nil {
			db.Set("gorm:table_options", "Engine=InnoDB")
		}
	case "sqlite", "sqlite3":
		db, err = gorm.Open(sqlite.Open(conf.Database), &gorm.Config{})
	}
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database")
	}
}

// Close function closes the database connection
func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}
	sqlDB.Close()
}
