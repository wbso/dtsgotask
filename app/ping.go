package app

import "context"

type PingRequest struct {
	Name string `json:"name"`
}
type PingResponse struct {
	Message string `json:"message"`
	Echo    string `json:"echo"`
}

func (app *App) Ping(ctx context.Context, input PingRequest) PingResponse {
	return PingResponse{Message: "pong", Echo: input.Name}
}
