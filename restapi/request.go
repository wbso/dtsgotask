package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func decodeRequest(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return fmt.Errorf("invalid request: %w", err)
	}
	return nil
}
