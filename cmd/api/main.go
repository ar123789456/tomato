package main

import (
	"log"
	"net/http"
	"tomato/server"
)

func main() {
	Run()
}

func Run() {
	db := server.InitDB()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	mux := http.NewServeMux()
	err := http.ListenAndServe("localhost:8080", mux)
	log.Println(err)
}
