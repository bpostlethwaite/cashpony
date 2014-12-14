package piper

import (
	"sync"

	"github.com/bpostlethwaite/cashpony/message"

	"testing"
)

func TestThroughput(t *testing.T) {

	p1 := NewPiped(0, 0)
	p2 := NewPiped(0, 0)

	p1.Pipe(p2)

	smsg := &message.Smsg{
		Msg: "hot damn!",
	}

	var rec *message.Smsg
	var txt string

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		p1.WriteTo <- smsg
	}()

	go func() {
		rec = <-p2.ReadFrom
		txt = rec.Msg
		wg.Done()
	}()

	wg.Wait()

	if txt != "hot damn!" {
		t.Error("Expected msg 'hot damn!' but got", txt)
	}

}

func TestDuplex(t *testing.T) {

	p1 := NewPiped(10, 10)
	p2 := NewPiped(10, 10)

	p1.Pipe(p2).Pipe(p1)

	smsg1 := &message.Smsg{
		Msg: "msg from 1",
	}
	smsg2 := &message.Smsg{
		Msg: "msg from 2",
	}

	var rec1, rec2 *message.Smsg
	var txt1 string
	// var rec1 *message.Smsg
	var txt2 string

	go func() {
		p1.WriteTo <- smsg1
	}()

	go func() {
		p2.WriteTo <- smsg2
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		rec1 = <-p2.ReadFrom
		txt1 = rec1.Msg
		wg.Done()
	}()

	go func() {
		rec2 = <-p1.ReadFrom
		txt2 = rec2.Msg
		wg.Done()
	}()

	wg.Wait()

	if txt1 != "msg from 1" {
		t.Error("Expected 'msg from 1' but got", txt1)
	}

	if txt2 != "msg from 2" {
		t.Error("Expected 'msg from 2' but got", txt2)
	}

}
