package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You SUCCESS DEPLOY GOLANG page 1</h1>")
}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods("GET")

	http.ListenAndServe(":", r)
}
