package stage

import (
	"fmt"
	"sort"

	"github.com/bpostlethwaite/cashpony/message"
	"github.com/bpostlethwaite/cashpony/piper"
	"github.com/bpostlethwaite/cashpony/record"
)

var sortTypes = []string{"date", "name", "debit", "label"}

func checkSortType(t string) bool {
	for _, v := range sortTypes {
		if v == t {
			return true
		}
	}
	return false
}

type Stage struct {
	*piper.Piped
	Recs     record.Records
	sortMode string
}

func NewStage() *Stage {
	r := &Stage{
		Piped: &piper.Piped{
			ReadFrom: make(chan message.Smsg, 5),
			WriteTo:  make(chan message.Smsg, 5),
		},
		sortMode: "date",
	}
	r.start(1)
	return r
}

func (this *Stage) start(n int) {
	var smsg message.Smsg

	if n < 1 {
		panic("can't start a pipe with n less than 1")
	}

	for i := 0; i < n; i++ {
		go func() {
			for smsg = range this.WriteTo {

				var updated bool

				rec := &smsg.Record            // ptr to Incoming smsg Rec
				rin := this.lookUpById(rec.Id) // ptr to Internally stored Rec

				// swap out rec for stored db item rin. First
				// apply updates from rec to rin, then continue
				// to modify rin disposing of rec.
				if rin == nil {
					// add record to database, rin is now pointer to stored rec
					this.Add(*rec)
					rin = this.lookUpById(rec.Id)
					updated = true
				} else {
					// update stored record r
					updated = rin.UpdateWith(rec)
				}

				// replace the record

				// No longer use rec, we now use pointer to
				// stored record r
				if updated {
					// send to Client on update channel

					this.ReadFrom <- smsg
				}

				if smsg.Flush != nil {
					go func() {
						for _, r := range this.Recs {
							smsg := message.Smsg{
								Record: r,
							}
							smsg.Flush <- smsg
						}
						close(smsg.Flush)
						smsg.Flush = nil
					}()
				}
			}
		}()
	}
}

func (this *Stage) Add(rec record.Record) {
	this.Recs = append(this.Recs, rec)
}

func (this *Stage) lookUpById(id string) *record.Record {
	for i := 0; i < len(this.Recs); i++ {
		if this.Recs[i].Id == id {
			return &this.Recs[i]
		}
	}
	return nil
}

func (r *Stage) Len() int {
	return len(r.Recs)
}

func (r *Stage) Less(i, j int) bool {

	s := r.Recs

	switch r.sortMode {
	case "name":
		return s[i].Name < s[j].Name

	case "label":
		return s[i].Label < s[j].Label

	case "debit":
		return s[i].Debit < s[j].Debit

	default: // Date
		return s[i].Date.Before(s[j].Date)

	}
}

func (r *Stage) Swap(i, j int) {
	s := r.Recs
	s[i], s[j] = s[j], s[i]
}

func (r *Stage) SortBy(args ...string) {

	reverse := false

	var t string

	if len(args) == 0 {
		t = "date"
	} else {
		t = args[0]
	}

	if len(args) == 2 && args[1] == "reverse" {
		reverse = true
	}

	if !checkSortType(t) {
		fmt.Println("Error, bad sorting type --- ignoring")
	}

	r.sortMode = t

	sort.Sort(r)

	if reverse {
		s := r.Recs
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
}
