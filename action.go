package main

import (
	"github.com/gotk3/gotk3/gtk"
)

type buttonInitParam struct {
	l   Logger
	cfg *ConfigFile
	w   *gtk.Window
}

type ButtonHandler interface {
	Handle(evtData interface{})
	Key() string
	Init(buttonInitParam)
	SetObject(btn *gtk.Button)
	GetObject() *gtk.Button
}

type AbstractHandler struct {
	cfg  *ConfigFile
	l    Logger
	btn  *gtk.Button
	root *gtk.Window
}

func (h *AbstractHandler) Init(c buttonInitParam) {
	h.l = c.l
	h.cfg = c.cfg
	h.root = c.w
}

func (h *AbstractHandler) SetObject(btn *gtk.Button) {
	h.btn = btn
}

func (h *AbstractHandler) GetObject() *gtk.Button {
	return h.btn
}

func UpdateText(b ButtonHandler) {
	if x := b.GetObject(); x != nil {
		x.SetLabel(L(b.Key()))
		x.QueueDraw()
	}
}
