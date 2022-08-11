package restapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/wbso/dtsgotask/app"
)

func (s *Server) listTasks(w http.ResponseWriter, r *http.Request) {
	res, err := s.App.ListTasks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondJSON(w, r, http.StatusOK, res)
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	var input app.CreateTaskRequest
	if err := decodeRequest(r, &input); err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	res, err := s.App.CreateTask(r.Context(), input)
	if err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, r, http.StatusCreated, res)
}

func (s *Server) getTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		respondJSON(w, r, http.StatusBadRequest, err)
		return
	}

	res, err := s.App.GetTask(r.Context(), id)
	if err != nil {
		respondJSON(w, r, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, r, http.StatusOK, res)
}

func (s *Server) updateTask(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	var input app.UpdateTaskRequest
	if err := decodeRequest(r, &input); err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	res, err := s.App.UpdateTask(r.Context(), id, input)
	if err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, r, http.StatusOK, res)
}

func (s *Server) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		respondJSON(w, r, http.StatusBadRequest, err)
		return
	}
	res, err := s.App.DeleteTask(r.Context(), id)
	if err != nil {
		respondJSON(w, r, http.StatusBadRequest, err)
		return
	}
	respondJSON(w, r, http.StatusOK, res)
}

func (s *Server) setTaskDone(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		respondJSON(w, r, http.StatusBadRequest, err)
		return
	}
	res, err := s.App.SetTaskDone(r.Context(), id)
	if err != nil {
		respondJSON(w, r, http.StatusBadRequest, err)
		return
	}
	respondJSON(w, r, http.StatusOK, res)
}

func (s *Server) setTaskUndone(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		respondJSON(w, r, http.StatusBadRequest, err)
		return
	}
	res, err := s.App.SetTaskUndone(r.Context(), id)
	if err != nil {
		respondJSON(w, r, http.StatusBadRequest, err)
		return
	}
	respondJSON(w, r, http.StatusOK, res)
}
