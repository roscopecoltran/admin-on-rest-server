package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/roscopecoltran/admin-on-rest-server/server/config"
)

var (
	DB             *gorm.DB
	Status         bool
	SqliteFilePath string
)

type DBHandler struct {
	Client *gorm.DB
	Status bool
}

func init() {
	var err error

	dbConfig := config.Config.DB
	if config.Config.DB.Adapter == "mysql" {
		DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		DB = DB.Set("gorm:table_options", "CHARSET=utf8")
	} else if config.Config.DB.Adapter == "postgres" {
		DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name))
	} else if config.Config.DB.Adapter == "sqlite" {
		SqliteFilePath = fmt.Sprintf("./%v.db", dbConfig.Name)
		// SqliteFilePath = fmt.Sprintf("%v/%v", os.TempDir(), dbConfig.Name)
		fmt.Printf("DB Sqlite, filepath: %s \n", SqliteFilePath)
		DB, err = gorm.Open("sqlite3", SqliteFilePath)
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		if os.Getenv("DEBUG") != "" {
			DB.LogMode(true)
		}
		DB.SingularTable(true)
		Status = true
	} else {
		panic(err)
	}
}

/*
func New(cfg *config.Config) {
	var err error

	dbConfig := cfg.DB
	if cfg.DB.Adapter == "mysql" {
		DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		DB = DB.Set("gorm:table_options", "CHARSET=utf8")
	} else if cfg.DB.Adapter == "postgres" {
		DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name))
	} else if cfg.DB.Adapter == "sqlite" {
		SqliteFilePath = fmt.Sprintf("%v/%v", os.TempDir(), dbConfig.Name)
		fmt.Printf("DB Sqlite, filepath: %s \n", SqliteFilePath)
		DB, err = gorm.Open("sqlite3", SqliteFilePath)
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		if os.Getenv("DEBUG") != "" {
			DB.LogMode(true)
		}
		DB.SingularTable(true)
		Status = true
	} else {
		panic(err)
	}
}
*/
