package routes

import (
	"cockpitserver/shared"
	"encoding/json"
	"net/http"
)

type TriggerEventRequestType struct {
	Name string `json:"eventname"`
}

var connectionForEvents, _ = shared.CreateSimConnectConnection("events-connection")

// Event is the http handler function
func Event(w http.ResponseWriter, r *http.Request) {
	var req TriggerEventRequestType
	json.NewDecoder(r.Body).Decode(&req)
	TriggerEvent(req.Name)
	w.Write([]byte("{\"status\":\"OK\"}"))
}

func TriggerEvent(eventName string) {
	event := connectionForEvents.NewSimEvent(shared.SimEventMap[eventName])
	event.Run()
}
