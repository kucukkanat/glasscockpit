package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/markbates/pkger"
)

type LatLngTuple struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type AirportsRequest struct {
	SouthWest LatLngTuple `json:"southWest"`
	NorthEast LatLngTuple `json:"northEast"`
}

// Airports is the Request handler
func Airports(w http.ResponseWriter, r *http.Request) {

	var decodedRequest AirportsRequest
	json.NewDecoder(r.Body).Decode(&decodedRequest)
	filteredAirports := FilterAirportsByBoundary(decodedRequest)

	// Marshall and respond with filtered airports
	marshalled, _ := json.Marshal(filteredAirports)
	w.Write(marshalled)
}

type Airport struct {
	ID        string
	Type      string
	Name      string
	Latitude  float64
	Longitude float64
	Elevation float64
	Country   string
	Code      string
}

func ReadAllAirportData() []Airport {
	content, _ := pkger.Open("/data/airports.csv")
	airportDataRaw, _ := ioutil.ReadAll(content)
	airportDataAsString := strings.Split(string(airportDataRaw), "\n")[1:]

	response := []Airport{}
	for _, row := range airportDataAsString {
		fields := strings.Split(row, ",")

		if len(fields) >= 13 {
			lat, _ := strconv.ParseFloat(fields[4], 64)
			lng, _ := strconv.ParseFloat(fields[5], 64)
			elv, _ := strconv.ParseFloat(fields[6], 64)
			response = append(response, Airport{
				ID:        fields[0],
				Type:      fields[2],
				Name:      strings.ReplaceAll(fields[3], "\"", ""),
				Latitude:  lat,
				Longitude: lng,
				Elevation: elv,
				Country:   strings.ReplaceAll(fields[8], "\"", ""),
				Code:      strings.ReplaceAll(fields[12], "\"", ""),
			})
		}
		// break
	}
	return response
}

var allAirports = ReadAllAirportData()

func inRange(n float64, min float64, max float64) bool {
	return n < max && n > min
}

func FilterAirportsByBoundary(bounds AirportsRequest) []Airport {
	filteredAirports := []Airport{}
	for _, airport := range allAirports {
		if inRange(airport.Longitude, bounds.SouthWest.Lng, bounds.NorthEast.Lng) && inRange(airport.Latitude, bounds.SouthWest.Lat, bounds.NorthEast.Lat) && airport.Type != "heliport" {
			filteredAirports = append(filteredAirports, airport)
		}
	}
	return filteredAirports
}
