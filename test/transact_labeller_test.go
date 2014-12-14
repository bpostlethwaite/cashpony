package integration

import (
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/bpostlethwaite/cashpony/labeller"
	"github.com/bpostlethwaite/cashpony/transacter"
)

const DATADIR = "data"

func TestTransLabel(t *testing.T) {

	dir, err := filepath.Abs(DATADIR)
	if err != nil {
		log.Fatal("fatal filepath %s", err)
	}
	t := transact.NewTransact(dir)
	l := label.NewLabeller(filepath.Join(dir, "labels.json"))
	stdout := newStdOut()

	t.Pipe(l).Pipe(stdout)

	time.Sleep(3 * time.Second)

	t.Error("things")
}
