package controllers

import (
	"app/src/utils"
	"github.com/gorilla/mux"
	"github.com/stereoflo/goach/models"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	data := models.GetList()
	utils.Respond(w, data)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := models.GetBoardById(vars["boardId"])
	utils.Respond(w, data)
}
