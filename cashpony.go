package main

import (
	"log"
	"path/filepath"

	"github.com/bpostlethwaite/cashpony/labeller"
	"github.com/bpostlethwaite/cashpony/stage"
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
	stager := stage.NewStage()

	server := web.NewServer()
	hub := server.Hub

	transacter.Pipe(labeller)
	labeller.Pipe(stager)
	stager.Pipe(hub)
	hub.Pipe(labeller)

	server.ListenAndServe()

}
