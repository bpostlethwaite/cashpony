package stage

import (
	"testing"
	"time"

	"github.com/bpostlethwaite/cashpony/record"
)

func TestSortRecordByDate(t *testing.T) {

	nameOrder := []string{"A", "B", "C", "A"}

	dateA := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	dateB := time.Date(2009, time.November, 11, 23, 0, 0, 0, time.UTC)
	dateC := time.Date(2009, time.November, 12, 23, 0, 0, 0, time.UTC)

	r := Stage{}

	r.Add(record.Record{Date: dateC, Name: "C"})
	r.Add(record.Record{Date: dateA, Name: "A"})
	r.Add(record.Record{Date: dateB, Name: "B"})

	r.SortBy("date")

	r.Add(record.Record{Date: dateC, Name: "A"})

	r.SortBy("date")

	recs := r.Recs

	for i := 0; i < len(r.Recs); i++ {
		if recs[i].Name != nameOrder[i] {
			t.Error("Expected ", nameOrder[i], "got", recs[i].Name)
		}

	}
}

func TestSortRecordByName(t *testing.T) {

	r := Stage{}
	r.Add(record.Record{Name: "b"})
	r.Add(record.Record{Name: "a"})
	r.Add(record.Record{Name: "c"})
	r.Add(record.Record{Name: "a"})

	r.SortBy("name")

	if r.Recs[0].Name != "a" ||
		r.Recs[1].Name != "a" ||
		r.Recs[2].Name != "b" ||
		r.Recs[3].Name != "c" {

		t.Error("Names failed to sort by name")
	}
}
