package record

import (
	"testing"
	"time"
)

func TestMatch(t *testing.T) {

	primary := Record{Transaction: "Agrimax 2"}
	near := Record{Transaction: "agrimax"}
	far := Record{Transaction: "Aquilaunch"}

	exactDist := primary.Match(primary)
	nearDist := primary.Match(near)
	farDist := primary.Match(far)

	if exactDist != 0 {
		t.Error("Expected exact distance", exactDist, " to equal 0")
	}
	if nearDist > farDist {
		t.Error("Expected ", nearDist, "to be less than ", farDist)
	}
}

func TestSortRecords(t *testing.T) {

	transactionOrder := []string{"A", "A", "B", "C"}

	dateA := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	dateB := time.Date(2009, time.November, 11, 23, 0, 0, 0, time.UTC)
	dateC := time.Date(2009, time.November, 12, 23, 0, 0, 0, time.UTC)

	Recs := Records{
		Record{Date: dateB, Transaction: "B"},
		Record{Date: dateC, Transaction: "C"},
		Record{Date: dateA, Transaction: "A"},
	}

	Recs.Sort()

	Recs = append(Recs, Record{Date: dateA, Transaction: "A"})

	Recs.Sort()

	for i := 0; i < len(Recs); i++ {
		if Recs[i].Transaction != transactionOrder[i] {
			t.Error("Expected ", transactionOrder[i], "got", Recs[i].Transaction)
		}

	}
}
