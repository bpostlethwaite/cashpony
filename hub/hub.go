package hub

import "github.com/bpostlethwaite/cashpony/record"

type PubSub struct {
	pub chan *record.Record
}

func (this *PubSub) Flush(chan *record.Record) {
}

func (this *PubSub) Subscribe(in chan *Smsg) chan *record.Record {
	out := make(chan *record.Record, 100)

	go func() {
		defer close(out)

		var r *record.Record
		var msg *Smsg

		pchan := this.pub

		select {
		case msg = <-in:
			// Handle the incoming message
			this.Handle(msg, out)
		case r = <-pchan:
			// forward
			out <- r

		}
	}()

	return out
}
