//go:generate go-bindata -pkg=main -ignore=\\.gitignore -o files.go main.ui

package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	cfg := loadConfig()
	gtk.Init(&os.Args)

	b := builder()
	w := window(b)
	defer w.Destroy()
	l := initLogger(b)

	initAction(l, cfg)
	b.ConnectSignals(setActionBtn(b))

	w.SetDefaultSize(600, 400)
	w.Connect("destroy", gtk.MainQuit)
	w.ShowAll()

	gtk.Main()
}

func builder() *gtk.Builder {
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatalf("failed to load glade ui file: %s", err)
	}

	data, _ := Asset("main.ui")
	if err = b.AddFromString(string(data)); err != nil {
		log.Fatalf("failed to load glade ui file: %s", err)
	}

	return b
}

func window(b *gtk.Builder) *gtk.Window {
	obj, err := b.GetObject("wroot")
	if err != nil {
		log.Fatalf("cannot load window object: %s", err)
	}

	w, ok := obj.(*gtk.Window)
	if !ok {
		log.Fatal("found an object with correct id, but not window")
	}

	return w
}
