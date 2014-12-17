package record

import "testing"

func TestIsEmpty(t *testing.T) {

	var r = Record{}

	if !r.IsEmpty() {
		t.Error("Expected IsEmpty to return true instead got", r.IsEmpty())
	}

	var r2 = Record{
		Name: "bert",
	}

	if r2.IsEmpty() {
		t.Error("Expected IsEmpty to return false instead got", r2.IsEmpty())
	}

}

func TestUpdateWith(t *testing.T) {

	var r1 = Record{
		Name:  "boots",
		Debit: 1001.2,
		Label: "gonz",
	}

	var r2 = Record{
		Name:  "bags",
		Debit: 1001.2,
		Label: "gonz",
	}

	var r3 = Record{
		Name:  "bags",
		Debit: 1001.2,
		Label: "gonz",
	}

	updated := r1.UpdateWith(r2)

	if !updated {
		t.Error("Expected record update with new info to return true instead got", updated)
	}

	updated = r2.UpdateWith(r3)

	if updated {
		t.Error("Expected record update with same info to return false instead got", updated)
	}

}
