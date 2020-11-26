package routes

import (
	"cockpitserver/shared"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/websocket/v2"
)

func Ws(c *websocket.Conn) {
	_, simChannel := shared.CreateSimConnectConnection("socketconn")
	for {
		result := <-simChannel
		changedVars := map[string]interface{}{}
		for _, simVar := range result {
			f, err := simVar.GetFloat64()
			if err != nil {
				panic(err)
			}
			changedVars[simVar.Name] = f
		}
		marshalled, _ := json.Marshal(changedVars)
		fmt.Printf("%+v", string(marshalled))
		time.Sleep(3 * time.Second)
		c.WriteMessage(websocket.TextMessage, marshalled)
	}

}
