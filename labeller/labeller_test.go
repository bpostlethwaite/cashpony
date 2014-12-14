package label

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"sync"
	"testing"

	"github.com/bpostlethwaite/cashpony/message"
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
		Name:  "strange",
		Label: "charm",
	}

	wg := l.AddLabel(r)
	wg.Wait()
	// wait for this label to commited to disk

	vals := l.Pairs().labels

	ans := strings.Join(vals, " ")
	if !strings.Contains(ans, "charm") {
		t.Error("Expected test data to contain 'charm'")
	}

}

func TestPipelineRecordUpdate(t *testing.T) {

	l := NewLabeller(dbfile)

	smsg := &message.Smsg{
		Record: &recorder.Record{
			Name:  "strange",
			Label: "boson",
		},
	}

	var label string
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		l.WriteTo <- smsg
	}()

	go func() {
		smsg = <-l.ReadFrom
		label = smsg.Record.Label
		wg.Done()
	}()

	wg.Wait()

	if label != "charm" {
		t.Error("Expected label 'charm' but got", label)
	}

	if !smsg.Record.Updated {
		t.Error("Expected record to be updated, but wasn't")
	}

}

func TestPipelineRecordRecycle(t *testing.T) {

	l := NewLabeller(dbfile)

	labels := l.Pairs().labels
	ans := strings.Join(labels, " ")
	if strings.Contains(ans, "boson") {
		t.Error("Expected boson to not be present in label store")
	}

	smsg := &message.Smsg{
		LabelUpdate: true,
		Record: &recorder.Record{
			Name:  "strange",
			Label: "boson",
		},
	}

	var label string
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		l.WriteTo <- smsg
	}()

	go func() {
		smsg = <-l.ReadFrom
		label = smsg.Record.Label
		wg.Done()
	}()

	wg.Wait()

	if label != "boson" {
		t.Error("Expected label 'boson' but got", label)
	}

	if !smsg.Record.Updated {
		t.Error("Expected record to be updated, but wasn't")
	}

	if !smsg.Recycle {
		t.Error("Expected Recycle Flag in Smsg to be set, but it wasn't")
	}

	labels = l.Pairs().labels
	ans = strings.Join(labels, " ")
	if !strings.Contains(ans, "boson") {
		t.Error("Expected boson to be in label store, not found")
	}

}
