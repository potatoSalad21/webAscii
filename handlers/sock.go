package handlers

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/potatoSalad21/webAscii/pkg/ascii"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func HandleWS(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		mt, buffer, err := ws.ReadMessage()
		if err != nil {
			return err
		}

		if mt != websocket.BinaryMessage {
			log.Println("[ERROR] Incorrect content type")
			continue
		}

		frame, err := ascii.Convert(buffer)
		if err != nil {
			log.Println("Cannot convert to ASCII.", err)
			continue
		}
		ws.WriteMessage(websocket.TextMessage, []byte(frame))
	}
}
