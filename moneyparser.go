package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/bpostlethwaite/cashpony/record"
)

const SHORTFORM = "01/02/2006"
const DATADIR = "data"

func main() {

	files, _ := ioutil.ReadDir(DATADIR)

	for _, f := range files {
		fname := filepath.Join(DATADIR, f.Name())
		eatit(fname)
	}

}

func eatit(file string) {

	csvfile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	for _, each := range rawCSVdata {
		date, err := time.Parse(SHORTFORM, each[0])
		if err != nil {
			log.Fatal(err)
		}

		transaction := each[1]

		debit, err := strconv.ParseFloat(each[2], 64)
		if err != nil {
			debit = 0.0
		}
		credit, err := strconv.ParseFloat(each[3], 64)
		if err != nil {
			credit = 0.0
		}
		balance, err := strconv.ParseFloat(each[4], 64)
		if err != nil {
			balance = 0.0
		}

		record := record.Record{
			Date:        date,
			Transaction: transaction,
			Debit:       debit,
			Credit:      credit,
			Balance:     balance,
		}

		fmt.Println(record.String() + "\n" + "-----------" + "\n")
	}
}
