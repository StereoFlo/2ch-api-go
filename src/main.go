package main

import (
	"app/src/controllers"
	"app/src/utils"
	"errors"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const boardList string = "/v1/board"
const boardShow string = "/v1/board/{boardId}"
const defaultPost string = "8000"

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
		port = defaultPost
	}
	return port
}

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle(boardList, MainHandler(http.HandlerFunc(controllers.GetList))).Methods("GET")
	router.Handle(boardShow, MainHandler(http.HandlerFunc(controllers.GetById))).Methods("GET")

	return router
}

func Listen() {
	http.Handle("/", GetRouter())
	http.ListenAndServe(":"+GetPort(), nil)
}

func MainHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = errors.New("error")
			}
			utils.RespondError(w, err)
		}()
		h.ServeHTTP(w, r)
	})
}
