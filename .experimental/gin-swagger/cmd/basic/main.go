package main

import (
	"fmt"
	"log"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/config"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/database"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/restapi"
)

var (
	apiConfig restapi.Config
)

func main() {

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	fmt.Printf("Database status: %t\n", database.Status)

	err := apiConfig.Parse()
	if err != nil {
		log.Fatal(err)
	}

	svc := &AdminOnRestServer{Health: false}

	api := restapi.NewAPI(svc, &apiConfig)
	// pp.Println(api)

	/*
		names := structs.Names(api.Routes)
		for _, name := range names {
			if name != "Engine" {
				fmt.Printf("name: %s\n", name)
				// tables = append(tables, models.(name))
			}
		}
	*/
	database.CreateTables(true, database.DefaultTables...)

	err = api.RunWithSigHandler()
	if err != nil {
		log.Fatal(err)
	}

}
