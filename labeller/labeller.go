package label

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/bpostlethwaite/cashpony/matcher"
	"github.com/bpostlethwaite/cashpony/message"
	"github.com/bpostlethwaite/cashpony/piper"
	"github.com/bpostlethwaite/cashpony/record"
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
	*piper.Piped
	dbfile       string
	labels       map[string]string
	matchDist    int
	defaultLabel string
}

type pairs struct {
	keys   []string
	labels []string
}

func NewLabeller(dbfile string, rf, wt int) *labeller {
	l := &labeller{
		Piped: &piper.Piped{
			ReadFrom: make(chan message.Smsg, rf),
			WriteTo:  make(chan message.Smsg, wt),
		},
		dbfile:       dbfile,
		matchDist:    matchDistance,
		defaultLabel: "Unknown",
	}

	l.LoadLabels()
	l.start(1)

	return l
}

func (this *labeller) MatchLabel(name string, maxDist int) (string, error) {

	pairs := this.Pairs()
	matches := matcher.NewMatch(
		name,
		pairs.keys,
	)
	match, err := matches.Best(maxDist)
	if err != nil {
		return "", err
	}
	idx := match.Index

	return pairs.labels[idx], nil
}

func (this *labeller) AddLabel(rec record.Record) *sync.WaitGroup {
	name := rec.Name
	label := rec.Label

	this.labels[name] = label
	fmt.Println("Labeller adding name:label", name, label)

	return this.SaveLabels()
}

func (this *labeller) start(n int) {

	var smsg message.Smsg

	if n < 1 {
		panic("can't start a pipe with n less than 1")
	}

	// Define a processing function
	process := func() {
		for smsg = range this.WriteTo {
			// grab the pointer to this Rec as we are going to update it's fields
			rec := &smsg.Record
			name := rec.Name

			// if this is a label-update then apply the label to the store and
			// mark this label as updated.
			if smsg.LabelUpdate {
				smsg.LabelUpdate = false

				this.AddLabel(*rec)

				// Right now we are flush on the main write channel
				// should investigate setting up an temporary channel
				// just for this.
				flush := make(chan message.Smsg)
				smsg.Flush = &flush

				go func() {
					// eventually wrap this up in a flush Type
					for smsg = range flush {
						// to ensure we don't get any infinite loopies
						smsg.LabelUpdate = false

						this.WriteTo <- smsg
					}

					fmt.Println("Labeller finishing flush")
				}()

				// fmt.Println("Labeller Label Update", name)
				this.ReadFrom <- smsg
				continue
			}

			// else see if we can match a label to this record
			label, err := this.MatchLabel(name, this.matchDist)
			if err == nil {
				fmt.Println("Found a match!", name, label)
				if label != rec.Label {
					rec.Updated = true
					rec.Label = label
				}
			}

			// if this label is till blank, lets assign it a default
			if rec.Label == "" {
				// fmt.Println("Applying default Label")
				rec.Label = this.defaultLabel

			}

			this.ReadFrom <- smsg
		}
	}

	// Call it concurrently
	for i := 0; i < n; i++ {
		go process()
	}

}

func (this *labeller) SaveLabels() *sync.WaitGroup {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		j, err := json.Marshal(this.labels)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(this.dbfile, j, 0644)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()

	return &wg
}

func (this *labeller) LoadLabels() {
	file, e := ioutil.ReadFile(this.dbfile)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &this.labels)
}

func (this *labeller) Pairs() pairs {
	p := pairs{
		keys:   make([]string, this.Len()),
		labels: make([]string, this.Len()),
	}

	idx := 0
	for k, v := range this.labels {
		p.keys[idx] = k
		p.labels[idx] = v
		idx++
	}

	return p
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
