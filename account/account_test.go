package account

import "testing"

func TestSort(t *testing.T) {

	var act = &Account{Records: nil}

	act.LoadCSV("test-data.csv", TDCC)
	act.Records.Sort()

	if act.Records[0].Transaction != "AMERICAN APPAREL         FRABEN HEIT" {
		t.Error("Expected 'AMERICAN APPAREL         FRABEN HEIT' but got",
			act.Records[0].Transaction)
	}
	if act.Records[1].Transaction != "CNP*THE NEW YORKER       800-825-25" {
		t.Error("Expected 'CNP*THE NEW YORKER       800-825-25' but got",
			act.Records[1].Transaction)
	}
	if act.Records[2].Transaction != "VDEV                     FRABEN HEIT" {
		t.Error("Expected 'VDEV                     FRABEN HEIT' but got",
			act.Records[2].Transaction)
	}

}
