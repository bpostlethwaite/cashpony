package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/bpostlethwaite/cashpony/message"
	"github.com/bpostlethwaite/cashpony/piper"
	"github.com/gorilla/websocket"
)

type Server struct {
	Hub *hub
}

func NewServer(rf, wt int) *Server {

	server := &Server{
		Hub: &hub{
			register:    make(chan *connection),
			unregister:  make(chan *connection),
			connections: make(map[*connection]bool),
			Piped: &piper.Piped{
				ReadFrom: make(chan message.Smsg, rf),
				WriteTo:  make(chan message.Smsg, wt),
			},
		},
	}

	return server
}

func (server *Server) ListenAndServe() {

	go server.Hub.run()

	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))

	http.HandleFunc("/ws", server.wsHandler)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func (server *Server) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	c := &connection{
		send: make(chan message.Smsg, 0),
		recv: make(chan message.Smsg, 0),
		ws:   ws,
	}

	server.Hub.register <- c

	defer func() {
		server.Hub.unregister <- c
	}()

	// init the connection
	go c.reader()
	go c.writer()
}
