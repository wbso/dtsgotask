package restapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Router
func (s *Server) Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// api routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/tasks", s.listTasks)
		r.Post("/tasks", s.createTask)
		r.Get("/tasks/{id}", s.getTaskByID)
		r.Put("/tasks/{id}", s.updateTask)
		r.Delete("/tasks/{id}", s.deleteTask)
		r.Put("/tasks/{id}/done", s.setTaskDone)
		r.Put("/tasks/{id}/undone", s.setTaskUndone)
	})

	// serve frontend files
	// r.Get("/*", ReverseProxy("http://localhost:5173/"))
	r.Get("/*", HandleEmbeddedIndexFile())

	return r
}
