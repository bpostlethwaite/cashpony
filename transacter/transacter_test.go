package transact

import (
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/bpostlethwaite/cashpony/util"
)

const DATADIR = "testdata"

func TestTransLabel(t *testing.T) {

	dir, err := filepath.Abs(DATADIR)
	if err != nil {
		log.Fatal("fatal filepath %s", err)
	}
	transacter := NewTransact(dir)
	stdout := util.NewStdOut()

	transacter.Pipe(stdout)

	go func() {
		// transacter should call unPipe
		// and close down the go process
		// linking to stdout.
		time.Sleep(time.Second)
		transacter.UnPipe()
	}()

	stdout.WaitUnPipe()
}
