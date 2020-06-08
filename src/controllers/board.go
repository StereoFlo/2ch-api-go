package controllers

import (
	"app/src/utils"
	"github.com/gorilla/mux"
	"github.com/stereoflo/goach/models/board"
	"github.com/stereoflo/goach/models/category"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	data, err := category.GetList()
	if err != nil {
		panic(err)
	}
	utils.Respond(w, data)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := board.GetBoardById(vars["boardId"])
	if err != nil {
		panic(err)
	}
	utils.Respond(w, data)
}
