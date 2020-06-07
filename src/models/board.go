package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type board struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type category struct {
	Name       string  `json:"name"`
	BoardCount int32   `json:"board_count"`
	Boards     []board `json:"boards"`
}

func GetList() []category {
	resp, respErr := http.Get("https://api.stereoflo.ru/v1/board")

	if respErr != nil {
		log.Fatal(respErr)
	}

	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)
	var categories []category

	jsonErr := json.Unmarshal(byteValue, &categories)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return categories
}
