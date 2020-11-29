package main

import (
	"cockpitserver/routes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/markbates/pkger"
)

var (
	addr     = flag.String("addr", "0.0.0.0:8080", "http service address")
	BuildENV string
)

func main() {
	if BuildENV == "demo" {
		TriggerDEMOCheck()
	}
	if BuildENV == "production" || BuildENV == "demo" {
		http.Handle("/", http.FileServer(pkger.Dir("/web/dist")))
	} else {
		http.Handle("/", http.FileServer(http.Dir("./web/dist")))
	}
	http.HandleFunc("/event", routes.Event)
	http.HandleFunc("/teleport", routes.Teleport)
	http.HandleFunc("/airports", routes.Airports)
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

func TriggerDEMOCheck() {
	fmt.Println("DEMO VERSION")
	kill := make(chan bool)
	type DemoAPIResponse struct {
		IsEnabled bool `json: "isEnabled"`
		Duration  int  `json: "duration"`
	}

	url := "https://raw.githubusercontent.com/kucukkanat/cockpitappconfig/main/config.json"
	resp, err := http.Get(url)
	if err != nil {
		panic("Request error")
	}
	var remote_config DemoAPIResponse

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	jsonErr := json.Unmarshal(body, &remote_config)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// Tell kill channel to kill process
	go func() {
		time.Sleep(time.Duration(remote_config.Duration) * time.Second)
		fmt.Printf("\nDemo version only works for %v seconds!", remote_config.Duration)
		kill <- true
	}()
	if !remote_config.IsEnabled {
		panic("This demo is not enabled!")
	}
	<-kill
	panic("Demo version terminated")

}
