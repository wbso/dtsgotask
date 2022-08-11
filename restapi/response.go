package restapi

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func respondError(w http.ResponseWriter, r *http.Request, status int, err error) {
	respondJSON(w, r, status, response{
		Data:    nil,
		Error:   err.Error(),
		Message: http.StatusText(status),
	})
}

func respondJSON(w http.ResponseWriter, _ *http.Request, status int, payload interface{}) {
	res := response{
		Error:   "",
		Message: http.StatusText(status),
		Data:    payload,
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
