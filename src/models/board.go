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
	resp, err := http.Get("https://api.stereoflo.ru/v1/board")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)
	var categories []category

	json.Unmarshal(byteValue, &categories)

	return categories
}
