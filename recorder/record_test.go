package recorder

import (
	"testing"
	"time"
)

func TestSortRecordByDate(t *testing.T) {

	transactionOrder := []string{"A", "B", "C", "A"}

	dateA := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	dateB := time.Date(2009, time.November, 11, 23, 0, 0, 0, time.UTC)
	dateC := time.Date(2009, time.November, 12, 23, 0, 0, 0, time.UTC)

	r := Recorder{}

	r.Add(Record{Date: dateC, Transaction: "C"})
	r.Add(Record{Date: dateA, Transaction: "A"})
	r.Add(Record{Date: dateB, Transaction: "B"})

	r.SortBy("date")

	r.Add(Record{Date: dateC, Transaction: "A"})

	r.SortBy("date")

	recs := r.Recs

	for i := 0; i < len(r.Recs); i++ {
		if recs[i].Transaction != transactionOrder[i] {
			t.Error("Expected ", transactionOrder[i], "got", recs[i].Transaction)
		}

	}
}

func TestSortRecordByName(t *testing.T) {

	r := Recorder{}
	r.Add(Record{Transaction: "b"})
	r.Add(Record{Transaction: "a"})
	r.Add(Record{Transaction: "c"})
	r.Add(Record{Transaction: "a"})

	r.SortBy("transaction")

	if r.Recs[0].Transaction != "a" ||
		r.Recs[1].Transaction != "a" ||
		r.Recs[2].Transaction != "b" ||
		r.Recs[3].Transaction != "c" {

		t.Error("Transactions failed to sort by name")
	}
}
