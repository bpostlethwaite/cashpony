package web

import (
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
)

type ServerHub struct {
	Msg MsgChan
}

type MsgChan struct {
	In  chan []byte
	Out chan []byte
}

var WebClient ServerHub

func init() {

	WebClient = ServerHub{
		Msg: MsgChan{
			In:  make(chan []byte),
			Out: make(chan []byte),
		},
	}

}

func (server *ServerHub) ListenAndServe() *MsgChan {

	go h.run(server.Msg)

	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))

	http.HandleFunc("/ws", wsHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	return &server.Msg
}
