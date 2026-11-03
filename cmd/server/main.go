package main

import (
	"log"
	"net/http"

	"github.com/mutenify/taskflow/internal/handler"
	"github.com/mutenify/taskflow/internal/repository"
	"github.com/mutenify/taskflow/internal/service"
)

func main() {
	repo := repository.NewRepository()
	serv := service.NewService(repo)
	h := handler.NewHandler(serv)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /task", h.PostTask)
	mux.HandleFunc("GET /status/{task_id}", h.GetStatus)
	mux.HandleFunc("GET /result/{task_id}", h.GetResult)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
