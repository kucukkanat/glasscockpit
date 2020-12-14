package routes

import (
	"cockpitserver/shared"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var _, variableChannel = shared.CreateSimConnectConnection("socket-connection")

func Ws(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	// Cross origin allow
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, _err := upgrader.Upgrade(w, r, nil)
	if _err != nil {
		fmt.Println(_err)
	}
	for {
		marshalled, _ := json.Marshal(<-variableChannel)
		ws.WriteMessage(websocket.TextMessage, marshalled)
	}
}
