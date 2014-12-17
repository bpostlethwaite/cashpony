package piper

import "github.com/bpostlethwaite/cashpony/message"

type Piper interface {
	getPipe(chan struct{}) *chan message.Smsg
	Pipe(Piper) Piper
	UnPipe()
	WaitUnPipe()
	Destroy()
}

type Piped struct {
	ReadFrom chan message.Smsg // Output
	WriteTo  chan message.Smsg // Input
	pipee    Piper
	done     chan struct{}
}

func NewPiped(rbuf, wbuf int) *Piped {
	p := &Piped{}
	p.ReadFrom = make(chan message.Smsg, rbuf)
	p.WriteTo = make(chan message.Smsg, wbuf)

	p.start(1)
	return p
}

func (this *Piped) Pipe(p Piper) Piper {

	if this.done == nil {
		this.done = make(chan struct{})
	}
	WriteTo := *p.getPipe(this.done)

	go func() {
		for smsg := range this.ReadFrom {
			select {
			case WriteTo <- smsg:
			case <-this.done:
				return
			}
		}
	}()

	return p
}

func (this *Piped) getPipe(done chan struct{}) *chan message.Smsg {
	if this.done == nil || this.done == done {
		this.done = done
	} else {
		panic("Can not pipe to a Piper already being Piped to")
	}
	return &this.WriteTo
}

func (this *Piped) UnPipe() {
	if this.done == nil {
		panic("Can not call UnPipe on an unpiped Piper")
	}
	close(this.done)
}

func (this *Piped) WaitUnPipe() {
	<-this.done
}

func (this *Piped) Destroy() {
	if this.done != nil {
		this.UnPipe()
	}
	this.done = nil
	close(this.ReadFrom)
	close(this.WriteTo)
}

// start - begin processing on the incoming WriteTo channel
// and writing onto the ReadFrom channel.
// This default start function just copies from the Pipes Write
// to Read channels without processing.
// This is the implementation of the Pass-Through Stream
func (this *Piped) start(n int) {

	if n < 1 {
		panic("can't start a pipe with n less than 1")
	}

	for i := 0; i < n; i++ {
		go func() {
			for smsg := range this.WriteTo {
				select {
				case this.ReadFrom <- smsg:
				case <-this.done:
					return
				}
			}
		}()
	}
}
