package message

import "github.com/bpostlethwaite/cashpony/recorder"

type Smsg struct {
	msg         string
	LabelUpdate bool
	Recycle     bool
	Record      *recorder.Record
}

type Piper interface {
	getPipe() *chan *Smsg
	Pipe(Piper) Piper
	start(int)
}

type Piped struct {
	ReadFrom chan *Smsg // Output
	WriteTo  chan *Smsg // Input
}

func NewPiped(rbuf, wbuf int) *Piped {
	p := &Piped{}
	p.ReadFrom = make(chan *Smsg, rbuf)
	p.WriteTo = make(chan *Smsg, wbuf)
	p.start(1)
	return p
}

func (this *Piped) Pipe(p Piper) Piper {

	var smsg *Smsg
	WriteTo := *p.getPipe()

	go func() {
		for smsg = range this.ReadFrom {
			// forward from this pipe to pipee
			WriteTo <- smsg
		}
	}()

	return p
}

func (this *Piped) getPipe() *chan *Smsg {
	return &this.WriteTo
}

// start - begin processing on the incoming WriteTo channel
// and writing onto the ReadFrom channel.
func (this *Piped) start(n int) {
	var smsg *Smsg

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
