package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/bpostlethwaite/cashpony/labeller"
	"github.com/bpostlethwaite/cashpony/transacter"
	"github.com/bpostlethwaite/cashpony/web"
)

const DATADIR = "data"

func main() {

	dir, err := filepath.Abs(DATADIR)
	if err != nil {
		log.Fatal("fatal filepath %s", err)
	}

	tor := transact.NewTransact(dir)
	lab := label.NewLabeller(filepath.Join(dir, "labels.json"))

	server := web.WebClient

	serverMsg := server.Msg

	go server.ListenAndServe()

	for {
		select {
		case <-serverMsg.Out:
			for _, rec := range tor.Transactions {
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
