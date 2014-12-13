package label

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bpostlethwaite/cashpony/matcher"
	"github.com/bpostlethwaite/cashpony/message"
	"github.com/bpostlethwaite/cashpony/recorder"
)

var Labels = map[string]bool{
	"Clothing":  true,
	"Household": true,
	"Dining":    true,
	"Transit":   true,
	"Food":      true,
	"Misc":      true,
	"Unknown":   true,
}

const matchDistance = 1

type labeller struct {
	dbfile    string
	labels    map[string]string
	matchDist int
}

func NewLabeller(dbfile string) *labeller {
	labeller := &labeller{dbfile: dbfile}

	labeller.LoadLabels()
	labeller.matchDist = matchDistance

	return labeller
}

func (this *labeller) MatchLabel(name string, maxDist int) string {

	matches := matcher.NewMatch(
		name,
		this.Values(),
	)

	return matches.Best(maxDist)
}

func (this *labeller) AddLabel(rec recorder.Record) {
	name := rec.Transaction
	label := rec.Label

	this.labels[name] = label

	go this.SaveLabels()
}

func (this *labeller) Pipe(in <-chan *message.Smsg) <-chan *message.Smsg {

	out := make(chan *message.Smsg, 10)

	go func() {

		for smsg := range in {
			out <- this.process(smsg)
		}
	}()

	return out
}

func (this *labeller) process(smsg *message.Smsg) *message.Smsg {

	rec := smsg.Record
	name := rec.Transaction

	// if this is a label-update then apply the label to the store and
	// mark this label as updated. Also add a
	if smsg.LabelUpdate {
		smsg.LabelUpdate = false

		this.labels[name] = rec.Label

		rec.Updated = true
		smsg.Recycle = true

		return smsg
	}

	// else see if we can match a label to this record
	label := this.MatchLabel(name, this.matchDist)

	if label != "" && label != rec.Label {
		rec.Updated = true
		rec.Label = label
	}

	return smsg
}

func (this *labeller) SaveLabels() {

	j, err := json.Marshal(this.labels)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(this.dbfile, j, 0644)
	if err != nil {
		panic(err)
	}
}

func (this *labeller) LoadLabels() {
	file, e := ioutil.ReadFile(this.dbfile)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &this.labels)
}

func (this *labeller) Values() []string {
	v := make([]string, this.Len())

	idx := 0
	for _, value := range this.labels {
		v[idx] = value
		idx++
	}

	return v
}

func (this *labeller) Len() int {
	return len(this.labels)
}

func (this *labeller) String() string {
	str := ""
	for k, v := range this.labels {
		str += k + " : " + v + "\n"
	}
	return str

}
