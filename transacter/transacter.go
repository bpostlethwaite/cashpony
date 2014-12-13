package transact

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/bpostlethwaite/cashpony/recorder"
)

type Transact struct {
	datadir      string
	Transactions recorder.Records
}

func NewTransact(datadir string) *Transact {
	transact := &Transact{datadir: datadir}

	transact.LoadAll()
	return transact

}

func (this *Transact) LoadAll() {
	files, _ := ioutil.ReadDir(this.datadir)
	for _, f := range files {
		if filepath.Ext(f.Name()) != "csv" {
			continue
		}
		fname := filepath.Join(this.datadir, f.Name())
		this.LoadCSV(fname, TDCC)
	}

}

func (this *Transact) LoadCSV(file string, template csvTemplate) {
	csvfile, err := os.Open(file)
	if err != nil {
		fmt.Println("transact csv open error", err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		log.Fatal("beltching on csv", file)
	}

	// sanity check, display to standard output
	for i, each := range rawCSVdata {
		date, err := time.Parse(template.DateFormat, each[template.DateField])
		if err != nil {
			log.Fatal(err)
		}

		transaction := each[template.TransactionField]

		debit, err := strconv.ParseFloat(each[template.DebitField], 64)
		if err != nil {
			debit = 0.0
		}

		t := recorder.Record{
			Date:        date,
			Transaction: transaction,
			Debit:       debit,
			Id:          file + ":" + string(i),
		}

		this.Transactions = append(this.Transactions, t)
	}
}
