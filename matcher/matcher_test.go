package matcher

import (
	"strings"
	"testing"
)

const needle = "aaaaa"

func TestSort(t *testing.T) {

	var haystack = []string{
		"bbbaa",
		"baaaa",
		"bbbba",
		"bbbbb",
		"aaaaa",
		"bbaaa",
	}

	matches := NewMatch(needle, haystack)

	for i, m := range matches.All() {
		// should produces sorted string arrag
		// aaaaa baaaa bbaaa ...
		nb := strings.Count(m.Label, "b")
		if nb != i {
			t.Error("Match failed to sort string")
		}
	}

}

func TestTopSingle(t *testing.T) {

	var haystack = []string{
		"bbbaa",
		"aaaaa",
		"bbbba",
		"aaaaa",
		"aaaaa",
		"bbaaa",
	}

	matches := NewMatch(needle, haystack)

	// we should get 3 top matches
	tops := matches.Top(-1)

	if len(tops) != 3 {
		t.Error("Match failed to pick 3 top matches")
	}

	for i := range tops {
		if tops[i].Label != needle {
			t.Error("Match failed to pick correct top matches")
		}
	}
}
