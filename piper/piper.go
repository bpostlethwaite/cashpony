package piper

import "github.com/bpostlethwaite/cashpony/message"

type Piper interface {
	getPipe() *chan *message.Smsg
	Pipe(Piper) Piper
	start(int)
}

type Piped struct {
	ReadFrom chan *message.Smsg // Output
	WriteTo  chan *message.Smsg // Input
}

func NewPiped(rbuf, wbuf int) *Piped {
	p := &Piped{}
	p.ReadFrom = make(chan *message.Smsg, rbuf)
	p.WriteTo = make(chan *message.Smsg, wbuf)
	p.start(1)
	return p
}

func (this *Piped) Pipe(p Piper) Piper {

	var smsg *message.Smsg
	WriteTo := *p.getPipe()

	go func() {
		for smsg = range this.ReadFrom {
			// forward from this pipe to pipee
			WriteTo <- smsg
		}
	}()

	return p
}

func (this *Piped) getPipe() *chan *message.Smsg {
	return &this.WriteTo
}

// start - begin processing on the incoming WriteTo channel
// and writing onto the ReadFrom channel.
// This default start function just copies from the Pipes Write
// to Read channels without processing.
// This is the implementation of the Pass-Through Stream
func (this *Piped) start(n int) {
	var smsg *message.Smsg

	if n < 1 {
		panic("can't start a pipe with n less than 1")
	}

	for i := 0; i < n; i++ {
		go func() {
			for smsg = range this.WriteTo {
				// forward from this pipe to pipee
				this.ReadFrom <- smsg
			}
		}()
	}
}
