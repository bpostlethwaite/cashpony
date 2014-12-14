package message

import "github.com/bpostlethwaite/cashpony/record"

type Smsg struct {
	Msg         string
	LabelUpdate bool
	Flush       chan Smsg
	Record      record.Record
}
