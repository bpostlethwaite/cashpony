package record

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/salviati/symutils/fuzzy"
)

type Record struct {
	Id          string    `json:"id"`
	Date        time.Time `json:"date"`
	Transaction string    `json:"transaction"`
	Debit       float64   `json:"debit"`
	Credit      float64   `json:"-"`
	Label       string    `json:"label"`
	Userset     bool      `json:"-"`
}

func (this *Record) String() string {
	return fmt.Sprintf("date:      %s\n", this.Date.String()) +
		fmt.Sprintf("transaction:  %s\n", this.Transaction) +
		fmt.Sprintf("debit:        %.2f\n", this.Debit) +
		fmt.Sprintf("label:        %s", this.Label)
}

func (this *Record) Json() ([]byte, error) {
	json, err := json.Marshal(this)
	return json, err
}

func (this *Record) Match(rec Record) int {

	cost := fuzzy.LevenshteinCost{
		Del:  1,
		Ins:  1,
		Subs: 1,
	}

	dist := fuzzy.Levenshtein(this.Transaction, rec.Transaction, &cost)

	return dist
}

type Records []Record

func (slice Records) Len() int {
	return len(slice)
}

func (slice Records) Less(i, j int) bool {
	return slice[i].Date.Before(slice[j].Date)
}

func (slice Records) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice Records) Sort() {
	sort.Sort(slice)
}
