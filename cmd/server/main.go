package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Task struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Result string `json:"result"`
}

const (
	taskNotFound = "task not found"
	statusReady  = "ready"
	resultDone   = "done"
)

// храним Task по значению (не *Task), чтобы метод Get возвращал копию структуры и внешний код не мог случайно изменить задачу в мапе в обход методов репозитория
// TODO: добавить sync.RWMutex для защиты мап от конкурентного доступа из горутин net/http
type Repository struct {
	tasks map[string]Task
}

func NewRepository() *Repository {
	return &Repository{tasks: make(map[string]Task)}
}

func (r *Repository) Add(t Task) {
	r.tasks[t.ID] = t
}

func (r *Repository) Get(id string) (Task, bool) {
	task, ok := r.tasks[id]
	return task, ok
}

type CreateResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type ResultResponse struct {
	Result string `json:"result"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Printf("encode response: %v", err)
	}
}

func main() {
	repo := NewRepository()
	mux := http.NewServeMux()

	mux.HandleFunc("POST /task", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		repo.Add(Task{ID: id, Status: statusReady, Result: resultDone})
		writeJSON(w, http.StatusCreated, CreateResponse{ID: id})
	})

	mux.HandleFunc("GET /status/{task_id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("task_id")
		task, ok := repo.Get(id)
		if !ok {
			writeJSON(w, http.StatusNotFound, ErrorResponse{Error: taskNotFound})
			return
		}
		writeJSON(w, http.StatusOK, StatusResponse{Status: task.Status})
	})

	mux.HandleFunc("GET /result/{task_id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("task_id")
		task, ok := repo.Get(id)
		if !ok {
			writeJSON(w, http.StatusNotFound, ErrorResponse{Error: taskNotFound})
			return
		}
		writeJSON(w, http.StatusOK, ResultResponse{Result: task.Result})
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("listen: %v", err)
	}
}
