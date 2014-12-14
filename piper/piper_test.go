package piper

import (
	"github.com/bpostlethwaite/cashpony/message"

	"testing"
	"time"
)

func TestThroughput(t *testing.T) {

	p1 := NewPiped(0, 0)
	p2 := NewPiped(0, 0)
	p3 := NewPiped(0, 0)

	p1.Pipe(p2).Pipe(p3)

	smsg := message.Smsg{
		Msg: "hot damn!",
	}

	var rec message.Smsg
	var txt string

	go func() {
		time.Sleep(50 * time.Millisecond)
		p1.WriteTo <- smsg
	}()

	go func() {
		rec = <-p3.ReadFrom
		txt = rec.Msg
		p3.UnPipe()
	}()

	p1.WaitUnPipe()

	if txt != "hot damn!" {
		t.Error("Expected msg 'hot damn!' but got", txt)
	}

}

func TestDuplex(t *testing.T) {

	p1 := NewPiped(10, 10)
	p2 := NewPiped(10, 10)

	p2.Pipe(p1).Pipe(p2)

	smsg1 := message.Smsg{
		Msg: "msg from 1",
	}
	smsg2 := message.Smsg{
		Msg: "msg from 2",
	}

	var rec1, rec2 message.Smsg
	var txt1 string
	// var rec1 *message.Smsg
	var txt2 string

	go func() {
		p1.WriteTo <- smsg1
		rec1 = <-p2.ReadFrom
		txt1 = rec1.Msg
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		p2.WriteTo <- smsg2
		rec2 = <-p1.ReadFrom
		txt2 = rec2.Msg
		p1.UnPipe()
	}()

	p1.WaitUnPipe()

	if txt1 != "msg from 1" {
		t.Error("Expected 'msg from 1' but got", txt1)
	}

	if txt2 != "msg from 2" {
		t.Error("Expected 'msg from 2' but got", txt2)
	}

}
