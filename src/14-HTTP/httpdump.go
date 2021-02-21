package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

func DumpRequest(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, string(requestDump))
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dumprequestG", DumpRequest).Methods("GET")
	router.HandleFunc("/dumprequestP", DumpRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}
