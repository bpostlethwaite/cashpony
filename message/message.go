package message

import "github.com/bpostlethwaite/cashpony/record"

type Smsg struct {
	Msg         string        `json:"msg"`
	LabelUpdate bool          `json:"labelupdate"`
	Flush       *chan Smsg    `json:"-"`
	Record      record.Record `json:"record"`
}

func (this *Smsg) HasRecInfo() bool {

	return !this.Record.IsEmpty()

}
