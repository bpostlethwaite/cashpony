package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/bpostlethwaite/cashpony/account"
)

const DATADIR = "data"

func main() {

	files, _ := ioutil.ReadDir(DATADIR)
	var act = &account.Account{Records: nil}

	for _, f := range files {
		fname := filepath.Join(DATADIR, f.Name())
		act.LoadCSV(fname, account.TDCC)
	}

	act.Print()

}
