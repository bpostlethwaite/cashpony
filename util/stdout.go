package util

import (
	"fmt"

	"github.com/bpostlethwaite/cashpony/message"
	"github.com/bpostlethwaite/cashpony/piper"
)

type StdOut struct {
	*piper.Piped
}

func NewStdOut() *StdOut {

	s := &StdOut{
		Piped: &piper.Piped{
			ReadFrom: make(chan message.Smsg, 5),
			WriteTo:  make(chan message.Smsg, 5),
		},
	}

	s.start(1)

	return s
}

func (this *StdOut) start(n int) {
	var smsg message.Smsg

	go func() {
		for smsg = range this.WriteTo {
			// forward from this pipe to pipee
			fmt.Println(smsg)
		}
	}()
}
