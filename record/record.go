package record

import (
	"fmt"
	"sort"
	"time"

	"github.com/salviati/symutils/fuzzy"
)

type Record struct {
	Date        time.Time
	Transaction string
	Debit       float64
	Credit      float64
	Label       string
}

// select.options[select.selectedIndex].value

func (this *Record) String() string {
	return "date: " + this.Date.String() + "\n" +
		"transaction: " + this.Transaction + "\n" +
		"debit: " + fmt.Sprintf("%.2f", this.Debit) + "\n" +
		"credit: " + fmt.Sprintf("%.2f", this.Credit) + "\n"
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
