package restapi

import (
	"encoding/json"
	"net/http"
)

func decodeRequest(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
