package web

import (
	"fmt"

	"github.com/bpostlethwaite/cashpony/message"
	"github.com/bpostlethwaite/cashpony/piper"
)

type hub struct {
	*piper.Piped
	// Registered connections.
	connections map[*connection]bool

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

func (this *hub) flush(c *connection) {
	// send a msg into the sytem with a
	// flush pipe and connect it to client

	flush := make(chan message.Smsg)

	smsg := message.Smsg{
		Flush: &flush,
	}

	this.ReadFrom <- smsg

	for smsg = range flush {
		c.send <- smsg
	}

	// once flushed add connection
	// to pool to receive updates
	this.connections[c] = true

	fmt.Println("Hub finishing flush")
}

func (this *hub) run() {
	for {
		select {
		case c := <-this.register:
			// first flush all records to conn
			// then add to connectrion pool
			fmt.Println("receiving a connection")
			go this.flush(c)
			go func() {
				for smsg := range c.recv {
					fmt.Println("Incoming from client", smsg)
					this.ReadFrom <- smsg
				}
			}()

		case c := <-this.unregister:
			if _, ok := this.connections[c]; ok {
				fmt.Print("/////////////// Unregistering connection and closing")
				delete(this.connections, c)
				close(c.send)
				close(c.recv)
			}

		// broadcast system update messages to all clients
		case m := <-this.WriteTo:
			for c := range this.connections {
				c.send <- m
			}
		}
	}
}
