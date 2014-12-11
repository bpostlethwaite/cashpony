package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/bpostlethwaite/cashpony/account"
	"github.com/bpostlethwaite/cashpony/web"
)

const DATADIR = "data"

var Labels = map[string]bool{
	"Clothing":  true,
	"Household": true,
	"Dining":    true,
	"Transit":   true,
	"Food":      true,
	"Misc":      true,
	"Unknown":   true,
}

func main() {

	files, _ := ioutil.ReadDir(DATADIR)
	var act = &account.Account{Records: nil}

	for _, f := range files {
		fname := filepath.Join(DATADIR, f.Name())
		act.LoadCSV(fname, account.TDCC)
	}

	//act.Print()
	server := web.WebClient

	serverMsg := server.Msg

	go server.ListenAndServe()

	for {
		select {
		case <-serverMsg.Out:
			for _, rec := range act.Records {
				json, err := rec.Json()
				if err != nil {
					fmt.Println("could not marshal record")
				} else {
					serverMsg.In <- json
				}
			}
		}
	}
}
