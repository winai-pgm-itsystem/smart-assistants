package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func KeduaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You SUCCESS DEPLOY GOLANG page 2 <====</h1>")
}

func Dua() {
	r := mux.NewRouter()
	r.HandleFunc("/", KeduaHandler).Methods("GET")

	http.ListenAndServe(":", r)
}
