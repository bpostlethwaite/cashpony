package record

import (
	"fmt"
	"time"

	"github.com/salviati/symutils/fuzzy"
)

type Record struct {
	Date        time.Time
	Transaction string
	Debit       float64
	Credit      float64
	Balance     float64
	Label       string
}

// method for type Person
func (this Record) String() string {
	return "date: " + this.Date.String() + "\n" +
		"transaction: " + this.Transaction + "\n" +
		"debit: " + fmt.Sprintf("%.2f", this.Debit) + "\n" +
		"credit: " + fmt.Sprintf("%.2f", this.Credit) + "\n" +
		"balance: " + fmt.Sprintf("%.2f", this.Balance)
}

func (this Record) Match(rec Record) int {

	cost := fuzzy.LevenshteinCost{
		Del:  1,
		Ins:  1,
		Subs: 1,
	}

	dist := fuzzy.Levenshtein(this.Transaction, rec.Transaction, &cost)

	return dist
}
