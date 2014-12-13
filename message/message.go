package message

import "github.com/bpostlethwaite/cashpony/recorder"

type Smsg struct {
	LabelUpdate bool
	Recycle     bool
	Record      *recorder.Record
}
