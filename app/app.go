package app

import (
	"github.com/wbso/dtsgotask/store"
)

type App struct {
	Repo *store.Queries
}

func New(repo *store.Queries) *App {
	return &App{
		Repo: repo,
	}
}
