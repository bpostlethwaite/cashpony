package stage

import (
	"testing"
	"time"

	"github.com/bpostlethwaite/cashpony/message"
	"github.com/bpostlethwaite/cashpony/piper"
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

func TestStoreUpdate(t *testing.T) {

	s := NewStage()
	p := piper.NewPiped(0, 0)

	p.Pipe(s)

	smsg := message.Smsg{
		Record: record.Record{
			Id:   "zap",
			Name: "a",
		},
	}

	p.WriteTo <- smsg

	time.Sleep(10 * time.Millisecond)

	rec := s.lookUpById("zap")

	if rec.Name != "a" {
		t.Error("Expected Stored Record with Name 'a' but got", rec.Name)
	}

	smsg.Record.Name = "b"

	// make sure internally stored record isn't linked
	// by reference to the record embedded in the smsg
	rec = s.lookUpById("zap")
	if rec.Name == "b" {
		t.Error("Modification via reference to stored Records not allowed")
	}

	// Sending modified record into Stage should update
	// the internal record
	p.WriteTo <- smsg
	time.Sleep(10 * time.Millisecond)

	rec = s.lookUpById("zap")

	if rec.Name != "b" {
		t.Error("Expected Stored Record with Name 'b' but got", rec.Name)
	}

}

func TestUpdatePipe(t *testing.T) {

	s := NewStage()
	p1 := piper.NewPiped(2, 2)
	p2 := piper.NewPiped(0, 0)

	p1.Pipe(s).Pipe(p2)

	smsg1 := message.Smsg{
		Record: record.Record{
			Id:   "zap",
			Name: "a",
		},
	}

	p1.WriteTo <- smsg1
	smsg := <-p2.ReadFrom

	if smsg1.Record.Name != smsg.Record.Name {
		t.Error("First messaage passes through should match message passed in")
	}

	smsg2 := message.Smsg{
		Record: record.Record{
			Id:   "zap",
			Name: "a",
		},
	}

	smsg3 := message.Smsg{
		Record: record.Record{
			Id:   "zap",
			Name: "zorh",
		},
	}

	p1.WriteTo <- smsg2
	p1.WriteTo <- smsg3

	smsg = <-p2.ReadFrom

	if smsg.Record.Name != smsg3.Record.Name {
		t.Error("Expected updated message with name Zorh, instead got", smsg3.Record.Name)
	}
}
