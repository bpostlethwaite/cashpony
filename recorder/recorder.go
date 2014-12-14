package recorder

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

var sortTypes = []string{"date", "transaction", "debit", "label"}

func checkSortType(t string) bool {
	for _, v := range sortTypes {
		if v == t {
			return true
		}
	}
	return false
}

type Recorder struct {
	Recs     Records
	sortMode string
}

type Records []Record

type Record struct {
	Id      string    `json:"id"`
	Date    time.Time `json:"date"`
	Name    string    `json:"name"`
	Debit   float64   `json:"debit"`
	Label   string    `json:"label"`
	Updated bool
}

func (this *Record) String() string {
	return fmt.Sprintf("date:      %s\n", this.Date.String()) +
		fmt.Sprintf("name:  %s\n", this.Name) +
		fmt.Sprintf("debit:        %.2f\n", this.Debit) +
		fmt.Sprintf("label:        %s", this.Label)
}

func (this *Record) Json() ([]byte, error) {
	j, err := json.Marshal(this)
	return j, err
}

func (r *Recorder) Add(rec Record) {
	r.Recs = append(r.Recs, rec)
}

func (r *Recorder) Len() int {
	return len(r.Recs)
}

func (r *Recorder) Less(i, j int) bool {

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

func (r *Recorder) Swap(i, j int) {
	s := r.Recs
	s[i], s[j] = s[j], s[i]
}

func (r *Recorder) SortBy(args ...string) {

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
