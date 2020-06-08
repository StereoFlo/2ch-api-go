package controllers

import (
	"app/src/utils"
	"github.com/gorilla/mux"
	"github.com/stereoflo/goach/models/board"
	"github.com/stereoflo/goach/models/category"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	data := category.GetList()
	utils.Respond(w, data)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := board.GetBoardById(vars["boardId"])
	utils.Respond(w, data)
}
