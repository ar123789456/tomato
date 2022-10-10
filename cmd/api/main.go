package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tomato/server"
	"tomato/servis"
	"tomato/servis/repository/mock"
	"tomato/servis/repository/sqlite"

	handler "tomato/servis/delivery/http"
	"tomato/servis/usecase"
)

func main() {
	Run()
}

func Run() {
	args := os.Args
	test := false
	if len(args) != 1 {
		test = args[1] == "dev"
	}
	var repo servis.Repository
	//init db
	if !test {
		db := server.InitDB()
		defer func() {
			err := db.Close()
			if err != nil {
				log.Println(err)
			}
		}()
		//init service
		repo = sqlite.NewRepository(db)
	} else {
		fmt.Println("dev")
		repo = mock.NewRepository()
	}
	useCase := usecase.NewUseCase(repo)
	handlers := handler.NewHandler(useCase)

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainPage)
	mux.HandleFunc("/user/create", handlers.CreateUser)
	mux.HandleFunc("/user/edit", handlers.EditUser)
	mux.HandleFunc("/user/", handlers.GetUser)
	mux.HandleFunc("/tomato/create", handlers.CreateTomato)
	mux.HandleFunc("/tomato/complete", handlers.StartTomato)
	mux.HandleFunc("/tomato", handlers.GetTomato)
	mux.HandleFunc("/tomato/delete", handlers.DeleteTomato)
	mux.HandleFunc("/analytics", handlers.GetTomatoNltx)

	log.Println("run localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	log.Println(err)
}
