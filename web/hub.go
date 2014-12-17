package web

import (
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
		Flush: flush,
	}

	this.ReadFrom <- smsg

	for smsg = range flush {
		c.send <- smsg
	}

	// once flushed add connection
	// to pool to receive updates
	this.connections[c] = true
}

func (this *hub) run() {
	for {
		select {
		case c := <-this.register:
			// first flush all records to conn
			// then add to connectrion pool
			go this.flush(c)
			go func() {
				for smsg := range c.recv {
					this.ReadFrom <- smsg
				}
			}()

		case c := <-this.unregister:
			if _, ok := this.connections[c]; ok {
				delete(this.connections, c)
				close(c.send)
			}

		// broadcast system update messages to all clients
		case m := <-this.WriteTo:
			for c := range this.connections {
				select {
				case c.send <- m:
				default:
					delete(this.connections, c)
					close(c.send)
				}
			}
		}
	}
}
