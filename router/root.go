package router

import (
	"kohhiapi/response"
	"net/http"
)

func rootRequest(w http.ResponseWriter, r *http.Request) {
	response.RespondByMessage(w, http.StatusForbidden, "request forbidden")
}
