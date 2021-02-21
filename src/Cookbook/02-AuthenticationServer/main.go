package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ADMIN_USER     = "admin"
	ADMIN_PASSWORD = "admin"
)

type Connection struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port`
}

func ReadFile(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	return string(data)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello World")
}

func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(ADMIN_USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(ADMIN_PASSWORD)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the application\n"))
			return
		}
		handler(w, r)
	}
}

func main() {
	http.HandleFunc("/", BasicAuth(helloWorld, "Please enter your username and password"))
	data := ReadFile("config.json")
	var connection Connection
	json.Unmarshal([]byte(data), &connection)
	fmt.Println(connection)

	err := http.ListenAndServe(connection.Hostname+":"+connection.Port, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
