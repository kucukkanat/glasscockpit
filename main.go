package main

import (
	"cockpitserver/routes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/markbates/pkger"
)

var (
	addr     = flag.String("addr", "0.0.0.0:8080", "http service address")
	BuildENV string
)

func main() {
	fmt.Println(BuildENV)
	if BuildENV == "production" {
		http.Handle("/", http.FileServer(pkger.Dir("/web/dist")))
	} else {
		http.Handle("/", http.FileServer(http.Dir("./web/dist")))
	}
	http.HandleFunc("/teleport", routes.Teleport)
	http.HandleFunc("/ws", routes.Ws)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func CopySimConnectDLLOnProduction() {
	source, err := pkger.Open("/SimConnect.dll")
	if err != nil {
		fmt.Println(err)
	}
	defer source.Close()

	destination, err2 := os.Create("SimConnect.dll")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer destination.Close()
	_, f_err := io.Copy(destination, source)
	if f_err != nil {
		fmt.Println(f_err)
	}
}
