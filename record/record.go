package record

import (
	"encoding/json"
	"fmt"
	"time"
)

type Records []Record

type Record struct {
	Id      string    `json:"id"`
	Date    time.Time `json:"date"`
	Name    string    `json:"name"`
	Debit   float64   `json:"debit"`
	Label   string    `json:"label"`
	Updated bool
}

// func (this *Recorder)

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

func (this *Record) UpdateWith(rec *Record) bool {

	updated := false
	if this.Id != rec.Id {
		this.Id = rec.Id
		updated = true
	}
	if this.Date != rec.Date {
		this.Date = rec.Date
		updated = true
	}
	if this.Name != rec.Name {
		this.Name = rec.Name
		updated = true
	}
	if this.Debit != rec.Debit {
		this.Debit = rec.Debit
		updated = true
	}
	if this.Label != rec.Label {
		this.Label = rec.Label
		updated = true
	}
	if this.Updated != rec.Updated {
		this.Updated = rec.Updated
		updated = true
	}

	return updated
}
