package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":5001", "lsiten address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("lsitening on %q...", *listen)
	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))))
}
