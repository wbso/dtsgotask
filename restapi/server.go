package restapi

import (
	"github.com/wbso/dtsgotask/app"
)

type Server struct {
	App *app.App
}

func New(app *app.App) *Server {
	return &Server{
		App: app,
	}
}
