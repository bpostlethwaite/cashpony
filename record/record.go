package record

import (
	"fmt"
	"time"
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
