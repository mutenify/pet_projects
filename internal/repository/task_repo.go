package repository

import (
	"github.com/mutenify/taskflow/internal/usecase"
)

type Task struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Result string `json:"result"`
}

type Repository struct {
	tasks map[string]Task
}

func NewRepository() *Repository {
	return &Repository{tasks: make(map[string]Task)}
}

func (r *Repository) Add(t Task) {
	r.tasks[t.ID] = t
}

func (r *Repository) Get(id string) (Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return Task{}, usecase.ErrNotFound
	}
	return task, nil
}

func (r *Repository) UpdateStatus(t Task, status string) {
	t = Task{ID: t.ID, Status: status, Result: t.Result}
	r.tasks[t.ID] = t
}

func (r *Repository) UpdateResult(t Task, result string) {
	t = Task{ID: t.ID, Status: t.Status, Result: result}
	r.tasks[t.ID] = t
}
