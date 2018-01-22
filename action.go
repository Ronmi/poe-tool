package main

import (
	"github.com/gotk3/gotk3/gtk"
)

type ButtonHandler interface {
	Handle(evtData interface{})
	Key() string
	Init(l Logger, cfg *ConfigFile)
	SetObject(btn *gtk.Button)
	GetObject() *gtk.Button
}

type AbstractHandler struct {
	cfg *ConfigFile
	l   Logger
	btn *gtk.Button
}

func (h *AbstractHandler) Init(l Logger, cfg *ConfigFile) {
	h.l = l
	h.cfg = cfg
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
