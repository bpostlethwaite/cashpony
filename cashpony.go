package main

import (
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

	transacter := transact.NewTransact(dir)
	labeller := label.NewLabeller(filepath.Join(dir, "labels.json"))

	// regular always-on Smsg channels
	transacter.Pipe(labeller).Pipe(recorder).Pipe(client).Pipe(labeller)

	// flush channels. (channels of channels)
	labeller.FlushFrom(recorder)
	client.FlushFrom(recorder)

	server := web.WebClient

	go server.ListenAndServe()

}
