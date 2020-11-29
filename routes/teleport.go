package routes

import (
	"cockpitserver/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

type TeleportRequest struct {
	PlaneAltitude  float64 `sim:"PLANE ALTITUDE" simUnit:"Feet" json:"alt"`
	PlaneLatitude  float64 `sim:"PLANE LATITUDE" simUnit:"Degrees" json:"lat"`
	PlaneLongitude float64 `sim:"PLANE LONGITUDE" simUnit:"Degrees" json:"lng"`
	Heading        float64 `sim:"PLANE HEADING DEGREES TRUE" simUnit:"Degrees" json:"hdg"`
}

var connection, _ = shared.CreateSimConnectConnection("setvar-connection")

func Teleport(w http.ResponseWriter, r *http.Request) {
	var reqParams TeleportRequest
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&reqParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set simulation variables
	connection.SetSimVarInterfaceInSim(reqParams)
	connection.SetSimVarInterfaceInSim(struct {
		Pitch float64 `sim:"PLANE PITCH DEGREES" simUnit:"Radians"`
		Bank  float64 `sim:"PLANE BANK DEGREES" simUnit:"Radians"`
	}{Pitch: 0, Bank: 0})

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Plane: %+v", reqParams)
}
