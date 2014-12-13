package matcher

import (
	"sort"
	"strings"

	"github.com/salviati/symutils/fuzzy"
)

// A match holds the match distance,
// and the label value the match was
// applied to. (should also hold the needle)
type match struct {
	dist  int
	label string
}

type matches []match

func (slice matches) Len() int {
	return len(slice)
}

func (slice matches) Less(i, j int) bool {
	return slice[i].dist <= slice[j].dist
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
		v[i] = m.label
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
		Del:  1,
		Ins:  1,
		Subs: 1,
	}

	m.matches = make([]match, len(haystack))

	for i, hay := range m.haystack {
		needle := strings.ToLower(m.needle)
		hay = strings.ToLower(hay)
		m.matches[i] = match{
			dist:  fuzzy.Levenshtein(needle, hay, &cost),
			label: hay,
		}
	}

	return m
}

// All return All match labels in sorted order
// of match distance.
func (this *matcher) List() []string {

	this.matches.Sort()

	return this.matches.Labels()
}

// If maxDist is -1 return matches with equal distance
// else return matches with distances less than or equal to
// maxDist
func (this *matcher) Top(maxDist int) []string {

	this.matches.Sort()

	var index int
	for i := 0; i < len(this.matches)-1; i++ {
		index = i
		if maxDist == -1 {
			if this.matches[i].dist != this.matches[i+1].dist {
				index++
				break
			}
		} else if this.matches[i].dist > maxDist {
			break
		}
	}

	return this.matches[:index].Labels()
}

// If maxDist is -1 return matches with equal distance
// else return matches with distances less than or equal to
// maxDist
func (this *matcher) Best(maxDist int) string {

	this.matches.Sort()

	if this.matches[0].dist > maxDist {
		return ""
	}

	return this.matches[0].label
}
