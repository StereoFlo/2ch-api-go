package main

import (
	"app/src/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	LoadEnv()
	Listen()
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}
	return port
}

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/v1/board", controllers.GetList).Methods("GET")
	router.HandleFunc("/v1/board/{boardId}", controllers.GetById).Methods("GET")

	return router
}

func Listen() {
	err := http.ListenAndServe(":"+GetPort(), GetRouter())

	if err != nil {
		fmt.Print(err)
		return
	}
}
