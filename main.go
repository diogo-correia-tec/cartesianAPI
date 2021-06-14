package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diogo-correia-tec/cartesianAPI/app"
	"github.com/diogo-correia-tec/cartesianAPI/util"

	_ "github.com/diogo-correia-tec/cartesianAPI/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func init() {
	util.LoadData() // Loads the default data present in file data/point.json
}

// @title Cartesian API
// @version 1.0
// @description API to Get points in a cartesian plane within distance of a given coordinate
// @contact.name Diogo Correia Araujo Silva
// @contact.email diogo.correia.tec@gmail.com
// @host localhost:8000
// @BasePath /api/
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/points", app.GetPoints).Methods(http.MethodGet)

	// Swagger
	router.PathPrefix("/api/swagger").Handler(httpSwagger.WrapHandler)

	fmt.Printf("Serving points on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
