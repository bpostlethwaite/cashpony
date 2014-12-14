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

	"github.com/bpostlethwaite/cashpony/message"
	"github.com/bpostlethwaite/cashpony/recorder"
)

type transact struct {
	datadir string
	pipe    chan *message.Smsg
}

func NewTransact(datadir string) *transact {
	transact := &transact{datadir: datadir}

	this.pipe = make(chan *message.Smsg, 100)
	transact.LoadAll()

	return transact

}

func (this *transact) Pipe(pipee *message.Pipeline) {
	// transacter.Pipe(labeller)

	var smsg *message.Smsg
	thispipe := this.pipe
	thatpipe := pipee.Fitting()

	go func() {
		select {
		case smsg = <-thispipe:
			// forward
			thatpipe <- smsg
		}
	}()

}

func (this *transact) LoadAll() {
	files, _ := ioutil.ReadDir(this.datadir)
	for _, f := range files {
		if filepath.Ext(f.Name()) != "csv" {
			continue
		}
		fname := filepath.Join(this.datadir, f.Name())
		this.LoadCSV(fname, TDCC)
	}

}

func (this *transact) LoadCSV(file string, template csvTemplate) {
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

		r := recorder.Record{
			Date:  date,
			Name:  transaction,
			Debit: debit,
			Id:    file + ":" + string(i),
		}

		smsg := &message.Smsg{
			Record: &r,
		}

		go func() {

			this.pipe <- smsg

		}()
	}
}
