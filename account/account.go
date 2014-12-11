package account

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bpostlethwaite/cashpony/matcher"
	"github.com/bpostlethwaite/cashpony/record"
	"github.com/jmcvetta/randutil"
)

type Account struct {
	Records record.Records
}

const SEARCHDIST = 3

func (this *Account) MatchingRecords(rec record.Record) record.Records {

	matches := make(matcher.Matches, len(this.Records))

	for i, r := range this.Records {
		matches[i] = matcher.Match{
			Record:   &r,
			Distance: rec.Match(r),
		}
	}

	topMatches := matches.Top(SEARCHDIST)
	matchingRecs := make(record.Records, len(topMatches))

	for i, m := range topMatches {
		matchingRecs[i] = *m.Record
	}

	return matchingRecs
}

func (this *Account) LoadCSV(file string, template csvTemplate) {

	csvfile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	for _, each := range rawCSVdata {
		date, err := time.Parse(template.DateFormat, each[template.DateField])
		if err != nil {
			log.Fatal(err)
		}

		transaction := each[template.TransactionField]

		debit, err := strconv.ParseFloat(each[template.DebitField], 64)
		if err != nil {
			debit = 0.0
		}
		credit, err := strconv.ParseFloat(each[template.CreditField], 64)
		if err != nil {
			credit = 0.0
		}

		id, err := randutil.AlphaString(8)
		if err != nil {
			log.Fatal("Why is the random string creator failing?")
		}

		record := record.Record{
			Date:        date,
			Transaction: transaction,
			Debit:       debit,
			Credit:      credit,
			Label:       "Unknown",
			Userset:     false,
			Id:          id,
		}

		this.Records = append(this.Records, record)
	}
}

func (this *Account) Print() {
	for _, record := range this.Records {
		fmt.Println(record.String() + "\n" + "-----------" + "\n")
	}
}
