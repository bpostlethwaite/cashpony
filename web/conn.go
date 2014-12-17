package web

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bpostlethwaite/cashpony/message"
	"github.com/gorilla/websocket"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	// ws sends on chan []byte
	send chan message.Smsg
	recv chan message.Smsg
}

func (c *connection) reader() {
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading incoming msg from client: ", err)
			break
		}
		smsg := message.Smsg{}
		err = json.Unmarshal(msg, &smsg)
		if err != nil {
			log.Fatal("couldn't parse incoming client message: ", err, string(msg))
			continue
		}
		// send message into system
		c.recv <- smsg
	}
	c.ws.Close()

}

func (c *connection) writer() {
	for smsg := range c.send {
		msg, err := json.Marshal(&smsg)
		if err != nil {
			log.Fatal(err)
		}
		err = c.ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
