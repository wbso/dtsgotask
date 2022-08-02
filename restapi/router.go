package restapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/wbso/dtsgotask/app"
)

// Router
func (s *Server) Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res := s.App.Ping(r.Context(), app.PingRequest{
			Name: r.URL.Query().Get("name"),
		})
		respondJson(w, r, http.StatusOK, res)
	})

	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		res := s.App.ListTasks(r.Context(), app.ListTasksRequest{
			Show: r.URL.Query().Get("show"),
		})
		respondJson(w, r, http.StatusOK, res)
	})

	r.Post("/tasks", func(w http.ResponseWriter, r *http.Request) {
		var input app.CreateTaskRequest
		if err := decodeRequest(r, &input); err != nil {
			respondJson(w, r, http.StatusBadRequest, err)
			return
		}
		res, err := s.App.CreateTask(r.Context(), input)
		if err != nil {
			respondJson(w, r, http.StatusBadRequest, err)
			return
		}
		respondJson(w, r, http.StatusCreated, res)
	})

	r.Get("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			respondJson(w, r, http.StatusBadRequest, err)
			return
		}
		res := s.App.GetTask(r.Context(), id)
		respondJson(w, r, http.StatusOK, res)
	})

	r.Put("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			respondJson(w, r, http.StatusBadRequest, err)
			return
		}
		var input app.UpdateTaskRequest
		if err := decodeRequest(r, &input); err != nil {
			respondJson(w, r, http.StatusBadRequest, err)
			return
		}
		res := s.App.UpdateTask(r.Context(), id, input)
		respondJson(w, r, http.StatusOK, res)
	})

	r.Delete("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			respondJson(w, r, http.StatusBadRequest, err)
			return
		}
		res := s.App.DeleteTask(r.Context(), id)
		respondJson(w, r, http.StatusOK, res)
	})

	r.Put("/tasks/{id}/complete", func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			respondJson(w, r, http.StatusBadRequest, err)
			return
		}
		res := s.App.CompleteTask(r.Context(), id)
		respondJson(w, r, http.StatusOK, res)
	})

	return r
}
