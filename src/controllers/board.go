package controllers

import (
	"app/src/utils"
	"goach/models"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	data := models.GetList()
	utils.Respond(w, data)
}
