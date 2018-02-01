package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

var (
	Handlers []ButtonHandler = []ButtonHandler{}
)

func switchLocale(lang string) {
	locale = lang
	for _, h := range Handlers {
		UpdateText(h)
	}
}

func setActionBtn(b *gtk.Builder, l Logger, cfg *ConfigFile) (signals map[string]interface{}) {
	signals = make(map[string]interface{})
	for _, x := range Handlers {
		obj, err := b.GetObject(x.Key())
		if err != nil {
			log.Fatalf("failed to load button %s: %s", x.Key, err)
		}

		btn, ok := obj.(*gtk.Button)
		if !ok {
			log.Fatalf("found an object named as %s, but not button", x.Key)
		}

		x.SetObject(btn)
		x.Init(l, cfg)
		UpdateText(x)
		signals[x.Key()] = x.Handle
	}

	return
}
