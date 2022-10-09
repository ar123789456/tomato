package main

import (
	"log"
	"net/http"
	"tomato/server"

	handler "tomato/servis/delivery/http"
)

func main() {
	Run()
}

func Run() {
	//init db
	db := server.InitDB()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	//init handler
	handlers := handler.NewHandler(nil)

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainPage)

	log.Println("run localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	log.Println(err)
}
