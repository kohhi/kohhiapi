package response

import (
	"encoding/json"
	"net/http"
)

// RespondByJSON is respond by json function.
func RespondByJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

// RespondByMessage is respond by message(json) function
func RespondByMessage(w http.ResponseWriter, statusCode int, message string) {
	RespondByJSON(w, statusCode, map[string]string{"message": message})
}
