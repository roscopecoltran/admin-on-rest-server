package main

import (
	"fmt"
	"log"

	"github.com/roscopecoltran/admin-on-rest-server/server/config"
	"github.com/roscopecoltran/admin-on-rest-server/server/database"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi"
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

	database.CreateTables(true, database.DefaultTables...)
	svc := &AdminOnRestServer{Health: false} // , DataSource: database.DB}

	api := restapi.NewAPI(svc, &apiConfig)

	err = api.RunWithSigHandler()
	if err != nil {
		log.Fatal(err)
	}

}

/*

  --help                         Show context-sensitive help (also try --help-long and --help-man).
  --debug                        Enable debug logging and pprof metrics.
  --address=":7000"              Address to listen on, e.g. :8080 or 0.0.0.0:7000.
  --insecure-http                Service only HTTP.
  --tls-cert-file=TLS-CERT-FILE  Path to TLS Cert file used when serving HTTPS.
  --tls-key-file=TLS-KEY-FILE    Path to TLS Key file used when serving HTTPS.
  --disable-well-known           Disable automatic /.well-known resources.
  --disable-auth                 Disable auth for all resources.
  --token-url=TOKEN-URL          Set TokenURL used to validate oauth2 tokens.

*/
