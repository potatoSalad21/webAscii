package handlers

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/potatoSalad21/webAscii/pkg/ascii"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWS(c echo.Context) error {
	// Receive, Convert to ASCII, Send back

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		mt, buffer, err := ws.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			continue
		}

		if mt != websocket.BinaryMessage {
			log.Println("[ERROR] Incorrect content type.")
			continue
		}

		ascii.Convert(buffer)
	}
}
