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
			ReadFrom: make(chan *message.Smsg, 5),
			WriteTo:  make(chan *message.Smsg, 5),
		},
		sortMode: "date",
	}
	r.start(1)
	return r
}

func (this *Stage) start(n int) {
	var smsg *message.Smsg

	if n < 1 {
		panic("can't start a pipe with n less than 1")
	}

	for i := 0; i < n; i++ {
		go func() {
			for smsg = range this.WriteTo {

				if smsg.Flush != nil {
					// this.flushTo(smsg)
				}
				// forward from this pipe to pipee
				this.ReadFrom <- smsg
			}
		}()
	}
}

func (r *Stage) Add(rec record.Record) {
	r.Recs = append(r.Recs, rec)
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
