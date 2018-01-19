package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type Logger interface {
	Log(string)
	Logf(string, ...interface{})
}

type logger struct {
	Obj *gtk.TextView
	Buf *gtk.TextBuffer
}

func initLogger(b *gtk.Builder) *logger {
	obj, err := b.GetObject("loggings")
	if err != nil {
		log.Fatalf("cannot load text view: %s", err)
	}

	v, ok := obj.(*gtk.TextView)
	if !ok {
		log.Fatal("found an object with correct id, but not text view")
	}

	buf, err := v.GetBuffer()
	if err != nil {
		log.Fatalf("cannot retrieve text buffer: %s", err)
	}

	return &logger{
		Obj: v,
		Buf: buf,
	}
}

func (l *logger) Log(msg string) {
	t := time.Now().Local().Format("2006-01-02 15:04:05: ")
	i := l.Buf.GetIterAtOffset(0)
	l.Buf.Insert(i, t+msg+"\n")
	l.Obj.QueueDraw()
}

func (l *logger) Logf(tmpl string, data ...interface{}) {
	l.Log(fmt.Sprintf(tmpl, data...))
}
