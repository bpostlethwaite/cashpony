package matcher

import (
	"sort"
	"strings"

	"github.com/salviati/symutils/fuzzy"
)

// A match holds the match distance,
// and the Label value the match was
// applied to. (should also hold the needle)
type match struct {
	Index int
	Dist  int
	Label string
}

type matches []match

func (slice matches) Len() int {
	return len(slice)
}

func (slice matches) Less(i, j int) bool {
	return slice[i].Dist <= slice[j].Dist
}

func (slice matches) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice matches) Sort() {
	sort.Sort(slice)
}

func (slice matches) Labels() []string {
	v := make([]string, len(slice))

	for i, m := range slice {
		v[i] = m.Label
	}

	return v
}

func (slice matches) Indicies() []int {
	v := make([]int, len(slice))

	for i, m := range slice {
		v[i] = m.Index
	}

	return v
}

type matcher struct {
	needle   string
	haystack []string
	matches  matches
}

func NewMatch(needle string, haystack []string) *matcher {
	m := &matcher{
		needle:   needle,
		haystack: haystack,
	}

	cost := fuzzy.LevenshteinCost{
		Del:  50,
		Ins:  5,
		Subs: 3,
	}

	m.matches = make([]match, len(haystack))

	for i, hay := range m.haystack {
		needle := strings.ToLower(m.needle)
		hay = strings.ToLower(hay)
		m.matches[i] = match{
			Index: i,
			Dist:  fuzzy.Levenshtein(needle, hay, &cost),
			Label: hay,
		}
	}

	return m
}

// All return All match Labels in sorted order
// of match distance.
func (this *matcher) All() matches {

	this.matches.Sort()

	return this.matches
}

// If maxDist is -1 return matches with equal distance
// else return matches with distances less than or equal to
// maxDist
func (this *matcher) Top(maxDist int) matches {

	this.matches.Sort()

	var idx int
	for i := 0; i < len(this.matches)-1; i++ {
		idx = i
		if maxDist == -1 {
			if this.matches[i].Dist != this.matches[i+1].Dist {
				idx++
				break
			}
		} else if this.matches[i].Dist > maxDist {
			break
		}
	}

	return this.matches[:idx]
}

// If maxDist is -1 return matches with equal distance
// else return matches with distances less than or equal to
// maxDist
func (this *matcher) Best(maxDist int) *match {

	this.matches.Sort()

	if this.matches[0].Dist > maxDist {
		return nil
	}

	return &this.matches[0]
}
