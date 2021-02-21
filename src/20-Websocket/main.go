package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("WebSocket Tutorial")

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		log.Println("New Connection")

		return nil
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
