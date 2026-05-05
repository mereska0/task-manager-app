package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"tmanager-app/service"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: s,
	}
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorWrite(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Task string `json:"task"`
	}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print(err)
	}
	if req.Task == "" {
		ErrorWrite(w, "empty task", http.StatusBadRequest)
		return
	}
	task, err := h.service.AddTask(r.Context(), req.Task)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			ErrorWrite(w, "timeout", http.StatusBadGateway)
			return
		}
		ErrorWrite(w, "db error", 500)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorWrite(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	tasks, err := h.service.GetTasks(r.Context())
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			ErrorWrite(w, "timeout", http.StatusBadGateway)
		}
		ErrorWrite(w, "db error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ErrorWrite(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorWrite(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req struct {
		Changes any `json:"changes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorWrite(w, "invalid json", http.StatusBadRequest)
		return
	}
	switch v := req.Changes.(type) {
	case bool:
		err := h.service.UpdateTaskStatus(r.Context(), v, id)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				ErrorWrite(w, "timeout", http.StatusGatewayTimeout)
				return
			}
			ErrorWrite(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case string:
		err := h.service.UpdateTaskText(r.Context(), v, id)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				ErrorWrite(w, "timeout", http.StatusGatewayTimeout)
				return
			}
			ErrorWrite(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		ErrorWrite(w, "changes must be string or bool", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		ErrorWrite(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorWrite(w, "invalid id", http.StatusBadRequest)
		return
	}
	err = h.service.DeleteTask(r.Context(), id)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			ErrorWrite(w, "timeout", http.StatusBadGateway)
			return
		}
		ErrorWrite(w, "db error", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) ClearTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		ErrorWrite(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	err := h.service.ClearTasks(r.Context())
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			ErrorWrite(w, "timeout", http.StatusBadGateway)
			return
		}
		ErrorWrite(w, "db error", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func ErrorWrite(w http.ResponseWriter, text string, status int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": text,
	})
}
