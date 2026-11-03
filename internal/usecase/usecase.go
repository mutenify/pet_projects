package usecase

import "errors"

const (
	StatusReady  = "ready"
	ResultDone   = "done"
	TaskNotFound = "task not found"
)

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

var ErrNotFound = errors.New("not found")
