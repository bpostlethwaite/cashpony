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
