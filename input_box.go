package main

import (
	"github.com/gotk3/gotk3/gtk"
)

func InputURL(l Logger) (str string) {
	box, err := gtk.DialogNew()
	if err != nil {
		l.Logf("cannot create dialog: %s", err)
		return
	}
	defer box.Destroy()
	entry, err := gtk.EntryNew()
	if err != nil {
		l.Logf("cannot create dialog: %s", err)
		return
	}
	label, err := gtk.LabelNew("pastebin URL")
	if err != nil {
		l.Logf("cannot create dialog: %s", err)
		return
	}

	box.AddButton("gtk-ok", gtk.RESPONSE_OK)
	box.AddButton("gtk-cancel", gtk.RESPONSE_CANCEL)

	entry.SetMaxLength(32)

	content, err := box.GetContentArea()
	if err != nil {
		l.Logf("cannot create dialog: %s", err)
		return
	}
	content.Add(label)
	content.Add(entry)
	box.ShowAll()

	if resp := box.Run(); resp == int(gtk.RESPONSE_OK) {
		str, _ = entry.GetText()
	}

	return
}
