package database

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/models"
)

var (
	DefaultTables = []interface{}{
		&models.ChoiceItem{},
		&models.DataSource{},
		&models.Entity{},
		&models.JwtUser{},
		&models.Permission{},
		&models.Role{},
		&models.User{},
		&models.GrantedAuthority{},
		&models.Apply{},
		&models.JwtAuthenticationRequest{},
	}
)

func CreateTables(isTruncate bool, tables ...interface{}) {
	for _, table := range tables {
		if isTruncate {
			if err := DB.DropTableIfExists(table).Error; err != nil {
				panic(err)
			}
		}
		fmt.Printf("AutoMigrate, table: %s\n", structs.Name(table))
		DB.AutoMigrate(table)
	}
}
