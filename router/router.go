package router

import (
	"kohhiapi/requests"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Router is router function.
func Router(port int) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", rootRequest)
	router.HandleFunc("/profile", requests.KohhiProfileRequest).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
