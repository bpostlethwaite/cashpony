package label

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/bpostlethwaite/cashpony/recorder"
)

const dbfile = "test_labels.json"

var testdata = map[string]string{
	"up":  "down",
	"top": "bottom",
}

func TestSetup(t *testing.T) {
	j, err := json.Marshal(testdata)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(dbfile, j, 0644)
}

func TestLabelMatching(t *testing.T) {
	l := NewLabeller(dbfile)

	label := l.MatchLabel("Top", 1)

	if label != "bottom" {
		t.Error("Expected label 'bottom' got", label, "instead")
	}

}

func TestLabelAdded(t *testing.T) {
	l := NewLabeller(dbfile)

	r := recorder.Record{
		Transaction: "strange",
		Label:       "charm",
	}

	l.AddLabel(r)

	vals := l.Values()

	ans := strings.Join(vals, " ")
	if !strings.Contains(ans, "charm") {
		t.Error("Expected test data to contain 'charm'")
	}

}

func TestPipeline(t *testing.T) {

	l := NewLabeller(dbfile)

}
