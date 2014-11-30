package matcher

import (
	"testing"

	"github.com/bpostlethwaite/cashpony/record"
)

func TestSort(t *testing.T) {

	distances := []int{3, 4, 2, 5, 1}
	sortedDistances := []int{1, 2, 3, 4, 5}

	matches := make(Matches, 5)

	for i := 0; i < 5; i++ {
		var r record.Record
		matches[i] = Match{Record: &r, Distance: distances[i]}
	}

	matches.Sort()

	for i := 0; i < 5; i++ {
		if matches[i].Distance != sortedDistances[i] {
			t.Error("Expected ",
				sortedDistances[i],
				" got ",
				matches[i].Distance)
		}
	}
}

func TestTopSingle(t *testing.T) {
	distances := []int{3, 4, 7, 1, 3}

	matches := make(Matches, 5)

	for i := 0; i < 5; i++ {
		var r record.Record
		matches[i] = Match{Record: &r, Distance: distances[i]}
	}

	topMatches := matches.Top(-1)

	if len(topMatches) != 1 {
		t.Error("Expected 1 top matches got ", len(topMatches))
	}

	if topMatches[0].Distance != 1 {
		t.Error("Top match does not have lowest distance")
	}
}

func TestTopMult(t *testing.T) {
	distances := []int{3, 4, 1, 1, 1}

	matches := make(Matches, 5)

	for i := 0; i < 5; i++ {
		var r record.Record
		matches[i] = Match{Record: &r, Distance: distances[i]}
	}

	topMatches := matches.Top(-1)

	if len(topMatches) != 3 {
		t.Error("Expected 3 top matches got ", len(topMatches))
	}

	if topMatches[0].Distance != 1 &&
		topMatches[1].Distance != 1 &&
		topMatches[2].Distance != 1 {
		t.Error("Top matches do not all have equal distance")
	}
}

func TestTopLessThan(t *testing.T) {
	distances := []int{3, 4, 7, 1, 3}

	matches := make(Matches, 5)

	for i := 0; i < 5; i++ {
		var r record.Record
		matches[i] = Match{Record: &r, Distance: distances[i]}
	}

	topMatches := matches.Top(3)

	if len(topMatches) != 3 {
		t.Error("Expected 3 top matches got ", len(topMatches))
	}

	if topMatches[0].Distance != 1 {
		t.Error("Top match does not have lowest distance")
	}
	if topMatches[1].Distance != 3 {
		t.Error("expected distance to be 3 got", topMatches[1].Distance)
	}
	if topMatches[2].Distance != 3 {
		t.Error("expected distance to be 3 got", topMatches[1].Distance)
	}

}
