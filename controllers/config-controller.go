package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Threx-code/go-api/config"
)

func CreateDatabase(w http.ResponseWriter, r *http.Request) {
	config.CreateDatabase()
}

func CreateTable(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("../../migrations/")
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, file := range files {
		query, err := ioutil.ReadFile("../../migrations/" + file.Name())
		if err != nil {
			log.Printf("%s File reading error", err.Error())
			return
		}

		fileName := strings.Split(string(file.Name()), ".")
		config.RunMigration(string(fileName[0]), string(query))
	}
}
