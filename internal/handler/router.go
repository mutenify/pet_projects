package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mutenify/taskflow/internal/service"
	"github.com/mutenify/taskflow/internal/usecase"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Printf("encode response: %v", err)
	}
}

func (h *Handler) PostTask(w http.ResponseWriter, r *http.Request) {
	id := h.svc.CreateTask()
	WriteJSON(w, http.StatusCreated, usecase.CreateResponse{ID: id})
}

func (h *Handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	task, err := h.svc.GetTask(r.PathValue("task_id"))
	if err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			WriteJSON(w, http.StatusNotFound, usecase.ErrorResponse{Error: usecase.TaskNotFound})
		} else {
			WriteJSON(w, http.StatusInternalServerError, usecase.ErrorResponse{Error: "internal error"})
			log.Printf("get status: %v", err)
		}
		return
	}
	WriteJSON(w, http.StatusOK, usecase.StatusResponse{Status: task.Status})
}

func (h *Handler) GetResult(w http.ResponseWriter, r *http.Request) {
	task, err := h.svc.GetTask(r.PathValue("task_id"))
	if err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			WriteJSON(w, http.StatusNotFound, usecase.ErrorResponse{Error: usecase.TaskNotFound})
		} else {
			WriteJSON(w, http.StatusInternalServerError, usecase.ErrorResponse{Error: "internal error"})
			log.Printf("get result: %v", err)
		}
		return
	}
	WriteJSON(w, http.StatusOK, usecase.ResultResponse{Result: task.Result})
}
