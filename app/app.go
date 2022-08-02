package app

import (
	"github.com/wbso/dtsgotask/store"
)

type App struct {
	Repo *store.Queries
}

func NewApp(repo *store.Queries) *App {
	return &App{
		Repo: repo,
	}
}

type CommonResponse struct {
	Err error
}

func (r *CommonResponse) SetError(err error) *CommonResponse {
	r.Err = err
	return r
}
