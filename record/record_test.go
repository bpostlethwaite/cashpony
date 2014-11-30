package record

import "testing"

func TestMatch(t *testing.T) {

	primary := Record{Transaction: "Agrimax 2"}
	near := Record{Transaction: "agrimax"}
	far := Record{Transaction: "Aquilaunch"}

	exactDist := primary.Match(primary)
	nearDist := primary.Match(near)
	farDist := primary.Match(far)

	if exactDist != 0 {
		t.Error("Expected exact distance", exactDist, " to equal 0")
	}
	if nearDist > farDist {
		t.Error("Expected ", nearDist, "to be less than ", farDist)
	}

}
