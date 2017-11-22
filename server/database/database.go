package database

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pressly/goose"
	"github.com/roscopecoltran/admin-on-rest-server/server/config"
)

var DB *gorm.DB
var Status bool

func init() {
	var err error

	dbConfig := config.Config.DB
	if config.Config.DB.Adapter == "mysql" {
		DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		DB = DB.Set("gorm:table_options", "CHARSET=utf8")
	} else if config.Config.DB.Adapter == "postgres" {
		DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name))
	} else if config.Config.DB.Adapter == "sqlite" {
		sqliteFilePath := fmt.Sprintf("%v%v.db", dbConfig.PrefixPath, dbConfig.Name)
		if config.Config.Debug {
			fmt.Printf("DB Sqlite, filepath: %s \n", sqliteFilePath)
		}
		DB, err = gorm.Open("sqlite3", sqliteFilePath)
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		if dbConfig.Debug {
			DB.LogMode(true)
		}
		if dbConfig.SingularTable {
			DB.SingularTable(true)
		}
		Status = true
	} else {
		panic(err)
	}

}

// Initialize GORM DB instance
func Initialize(config *config.ServiceConfig) *gorm.DB {
	var db *gorm.DB
	var err error
	dbConfig := config.DB
	if dbConfig.Adapter == "mysql" {
		db, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		db = db.Set("gorm:table_options", "CHARSET=utf8")
	} else if dbConfig.Adapter == "postgres" {
		db, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name))
	} else if dbConfig.Adapter == "sqlite" {
		sqliteFilePath := fmt.Sprintf("%v%v.db", dbConfig.PrefixPath, dbConfig.Name)
		if config.Debug {
			fmt.Printf("DB Sqlite, filepath: %s \n", sqliteFilePath)
		}
		db, err = gorm.Open("sqlite3", sqliteFilePath)
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		if dbConfig.Debug {
			db.LogMode(true)
		}
		if dbConfig.SingularTable {
			db.SingularTable(true)
		}
		Status = true
	} else {
		panic(err)
	}

	return db
}

// Migrate to latest version
func Migrate(config *config.ServiceConfig, db *gorm.DB) error {
	goose.Up(db.DB(), config.DB.MigrationsDir)
	return nil
}

// GetFromContext Get DB from Gin Context
func GetFromContext(c *gin.Context) (*gorm.DB, error) {
	dbi, _ := c.Get("DB")
	db, ok := dbi.(*gorm.DB)
	if !ok {
		return nil, errors.New("No db in context")
	}
	return db, nil
}
