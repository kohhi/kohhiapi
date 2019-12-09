package requests

import (
	"kohhiapi/kohhi"
	"kohhiapi/response"
	"net/http"
)

// KohhiProfileRequest is kohhi's profile request function.
func KohhiProfileRequest(w http.ResponseWriter, r *http.Request) {
	res, err := kohhi.GetProfile()
	if err != nil {
		response.RespondByMessage(w, http.StatusInternalServerError, err.Error())
	} else {
		response.RespondByJSON(w, http.StatusOK, res)
	}
}
