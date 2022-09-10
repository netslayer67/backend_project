package main

import (
	"fmt"
	"net/http"
	"testdumpflix/database"
	"testdumpflix/pkg/mysql"
	"testdumpflix/routes"

	"github.com/gorilla/mux"
)

func main() {

	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1/").Subrouter())

	fmt.Println("Starting API server localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
