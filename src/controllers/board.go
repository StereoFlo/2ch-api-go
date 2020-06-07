package controllers

import (
	"app/src/models"
	"app/src/utils"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	data := models.GetList()
	utils.Respond(w, data)
}
