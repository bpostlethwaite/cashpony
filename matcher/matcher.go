package matcher

import (
	"sort"

	"github.com/bpostlethwaite/cashpony/record"
)

type Match struct {
	Record   *record.Record
	Distance int
}

type Matches []Match

func (slice Matches) Len() int {
	return len(slice)
}

func (slice Matches) Less(i, j int) bool {
	return slice[i].Distance <= slice[j].Distance
}

func (slice Matches) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice Matches) Sort() {
	sort.Sort(slice)
}

// If maxDist is -1 return the top matches with equal distance
// else return all matches with distances less than or equal to
// maxDist
func (slice Matches) Top(maxDist int) []Match {
	slice.Sort()

	var index int
	for i := 0; i < len(slice)-1; i++ {
		index = i
		if maxDist == -1 {
			if slice[i].Distance != slice[i+1].Distance {
				index++
				break
			}
		} else if slice[i].Distance > maxDist {
			break
		}
	}

	return slice[:index]
}
