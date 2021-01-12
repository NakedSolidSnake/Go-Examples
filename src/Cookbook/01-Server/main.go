package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func main() {
	http.HandleFunc("/", helloWorld)
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
