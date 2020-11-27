package main

import (
	"cockpitserver/routes"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/markbates/pkger"
)

var (
	addr     = flag.String("addr", "127.0.0.1:8080", "http service address")
	BuildENV string
)

func main() {
	fmt.Println(BuildENV)

	if BuildENV == "production" {
		http.Handle("/", http.FileServer(pkger.Dir("/static")))
	} else {
		http.Handle("/", http.FileServer(http.Dir("./static")))
	}
	http.HandleFunc("/teleport", routes.Teleport)
	http.HandleFunc("/ws", routes.Ws)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
