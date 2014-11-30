package account

import (
	"github.com/bpostlethwaite/cashpony/matcher"
	"github.com/bpostlethwaite/cashpony/record"
)

type Account struct {
	Records []record.Record
}

const SEARCHDIST = 3

func (this Account) MatchingRecords(rec record.Record) []Match {

	matches := make(matcher.Matches, len(this.Records))

	for i, r := range this.Records {
		matches[i] = Match{
			Record:   r,
			Distance: rec.Match(r),
		}
	}

	topMatches = matches.Top(SEARCHDIST)
	matchingRecs := make(record.Record, len(topMatches))
	for i, r := range topMatches {
		matchingRecs[i] = topMatches[i].Record
	}

	return matchingRecs
}
