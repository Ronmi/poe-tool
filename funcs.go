package main

import (
	"log"
	"path/filepath"

	"github.com/gotk3/gotk3/gtk"
)

type Action struct {
	Func func(interface{})
	Key  string
	Obj  *gtk.Button
}

func (a *Action) UpdateText() {
	if a.Obj != nil {
		a.Obj.SetLabel(L(a.Key))
		a.Obj.QueueDraw()
	}
}

var (
	Actions []*Action
)

func switchLocale(lang string) {
	locale = lang
	for _, a := range Actions {
		a.UpdateText()
	}
}

func async(f func(interface{})) func(interface{}) {
	return func(d interface{}) {
		go f(d)
	}
}

func initAction(l Logger, cfg *ConfigFile) {
	h := FilterHnd(cfg)
	Actions = []*Action{
		{
			Func: func(eventData interface{}) {
				switchLocale("en")
				l.Log(L("btn_switch_lang_en"))
			},
			Key: "btn_switch_lang_en",
		},
		{
			Func: func(eventData interface{}) {
				switchLocale("tw")
				l.Log(L("btn_switch_lang_tw"))
			},
			Key: "btn_switch_lang_tw",
		},
		{
			Func: async(func(eventData interface{}) {
				l.Log(L("dling_trademacro"))
				err := InstallTradeMacro()
				if err != nil {
					l.Log(LErr(err))
				} else {
					l.Log(L("finish_trademacro"))
				}
			}),
			Key: "btn_inst_trademacro",
		},
		{
			Func: async(func(eventData interface{}) {
				l.Log(L("dling_pob"))
				err := InstallPoB()
				if err != nil {
					l.Log(LErr(err))
				} else {
					l.Log(L("finish_pob"))
				}
			}),
			Key: "btn_inst_pob",
		},
		{
			Func: async(func(eventData interface{}) {
				if err := InstallUncap(); err != nil {
					l.Log(LErr(err))
					return
				}
				if err := Caps2Ctrl(); err != nil {
					l.Log(LErr(err))
					return
				}
				l.Log(L("bind_caps"))
			}),
			Key: "btn_caps2ctrl",
		},
		{
			Func: async(func(eventData interface{}) {
				if err := InstallUncap(); err != nil {
					l.Log(LErr(err))
					return
				}
				if err := RestoreCaps(); err != nil {
					l.Log(LErr(err))
					return
				}
				l.Log(L("restore_caps"))
			}),
			Key: "btn_restore_caps",
		},
		{
			Func: async(func(eventData interface{}) {
				dir, err := GetFilterDir()
				if err != nil {
					l.Log(LErr(err))
					return
				}

				ch := h.DownloadFilters(dir)
				for x := range ch {
					if x.Filename == "" {
						l.Log(LErr(x.Err))
						continue
					}
					fn := filepath.Base(x.Filename)

					if !x.Finished {
						l.Logf(L("log_state_begin"), fn)
						continue
					}

					if x.Err == nil {
						l.Logf(L("log_state_ok"), x.Key, fn)
						continue
					}

					l.Logf(L("log_state_fail"), fn, x.Err)
				}
			}),
			Key: "btn_download_filters",
		},
		{
			Func: async(func(eventData interface{}) {
				dir, err := GetFilterDir()
				if err != nil {
					l.Log(LErr(err))
					return
				}

				ch := h.UploadFilters(dir)
				for x := range ch {
					if x.Filename == "" {
						l.Log(LErr(x.Err))
						continue
					}

					fn := filepath.Base(x.Filename)

					if !x.Finished {
						l.Logf(L("log_state_begin"), fn)
						continue
					}

					if x.Err == nil {
						l.Logf(L("log_state_ok"), fn, x.Key)
						continue
					}

					l.Logf(L("log_state_fail"), fn, x.Err)
				}
			}),
			Key: "btn_upload_filters",
		},
		{
			Func: func(eventData interface{}) {
				l.Log(L("not_implemented"))
			},
			Key: "btn_import_filter",
		},
	}
}

func setActionBtn(b *gtk.Builder) (signals map[string]interface{}) {
	signals = make(map[string]interface{})
	for _, x := range Actions {
		obj, err := b.GetObject(x.Key)
		if err != nil {
			log.Fatalf("failed to load button %s: %s", x.Key, err)
		}

		btn, ok := obj.(*gtk.Button)
		if !ok {
			log.Fatalf("found an object named as %s, but not button", x.Key)
		}

		x.Obj = btn
		x.UpdateText()
		signals[x.Key] = x.Func
	}

	return
}
