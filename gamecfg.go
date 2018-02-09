package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/Ronmi/pastebin"
	"github.com/Ronmi/poe-tool/gamecfg"
)

const intlCfgName = "poe.intl.ini"
const intlLocalName = "production_Config.ini"

func (h *filterHnd) findRemoteConfig() (key string, err error) {
	files, err := h.api.List(h.UserKey, 1000)
	if err != nil {
		return
	}

	for _, f := range files {
		log.Printf("%+v", f)
		if f.Title == intlCfgName {
			return f.Key, nil
		}
	}

	return
}

func (h *filterHnd) loadLocalGamecfg(dir string) (cfg map[string][]string, err error) {
	cfg = make(map[string][]string)
	f, err := os.Open(filepath.Join(dir, intlLocalName))
	if err != nil {
		return
	}
	defer f.Close()

	cfg, err = gamecfg.Read(f)
	return
}

func (h *filterHnd) UploadConfig(l Logger) {
	dir, err := GetFilterDir()
	if err != nil {
		l.Log(L("cannot_get_user_dir"))
		return
	}

	cfg, err := h.loadLocalGamecfg(dir)
	if err != nil {
		l.Log(LErr(err))
		return
	}

	data := &bytes.Buffer{}
	if err := gamecfg.Write(data, cfg); err != nil {
		l.Log(LErr(err))
		return
	}

	// remove remote config if exists
	if key, err := h.findRemoteConfig(); key != "" && err == nil {
		h.api.Delete(h.UserKey, key)
	}

	p := &pastebin.Paste{
		Title:   intlCfgName,
		Content: data.String(),
		UserKey: h.UserKey,
	}
	if _, err := h.api.Post(p); err != nil {
		l.Log(LErr(err))
	} else {
		l.Log(L("done_ul_gamecfg"))
	}
}

func (h *filterHnd) DownloadConfig(l Logger) {
	dir, err := GetFilterDir()
	if err != nil {
		l.Log(L("cannot_get_user_dir"))
		return
	}

	key, err := h.findRemoteConfig()
	if err != nil {
		l.Log(LErr(err))
		return
	}
	if key == "" {
		l.Log(L("remote_gamecfg_404"))
		return
	}

	data, err := h.api.UserPaste(h.UserKey, key)
	if err != nil {
		l.Log(LErr(err))
		return
	}

	buf := bytes.NewBuffer(data)
	remote, err := gamecfg.Read(buf)
	if err != nil {
		l.Log(LErr(err))
	}

	cfg, _ := h.loadLocalGamecfg(dir)

	// overwrite only UI/ACTION_KEYS/LANGUAGE/LOGIN/NETWORKING/NOTIFICATIONS and
	// TUTORIAL_FLAGS
	want := map[string]bool{
		"[UI]":             true,
		"[ACTION_KEYS]":    true,
		"[LANGUAGE]":       true,
		"[LOGIN]":          true,
		"[NETWORKING]":     true,
		"[NOTIFICATIONS]":  true,
		"[TUTORIAL_FLAGS]": true,
	}
	for k, v := range remote {
		if _, ok := want[k]; ok {
			cfg[k] = v
		}
		if _, ok := cfg[k]; !ok {
			cfg[k] = v
		}
	}

	f, err := os.Create(filepath.Join(dir, intlLocalName))
	if err != nil {
		l.Log(LErr(err))
		return
	}
	defer f.Close()

	if err := gamecfg.Write(f, cfg); err != nil {
		l.Log(LErr(err))
		return
	} else {
		l.Log(L("done_dl_gamecfg"))
	}
}

type ULCFGHandler struct {
	AbstractHandler
	h *filterHnd
}

func (h *ULCFGHandler) Init(c buttonInitParam) {
	h.AbstractHandler.Init(c)
	h.h = FilterHnd(c.cfg)
}

func (h *ULCFGHandler) Handle(data interface{}) {
	go h.h.UploadConfig(h.l)
}

func (h *ULCFGHandler) Key() string {
	return "btn_upload_cfg"
}

type DLCFGHandler struct {
	AbstractHandler
	h *filterHnd
}

func (h *DLCFGHandler) Init(c buttonInitParam) {
	h.AbstractHandler.Init(c)
	h.h = FilterHnd(c.cfg)
}

func (h *DLCFGHandler) Handle(data interface{}) {
	go h.h.DownloadConfig(h.l)
}

func (h *DLCFGHandler) Key() string {
	return "btn_download_cfg"
}

func init() {
	Handlers = append(
		Handlers,
		&ULCFGHandler{},
		&DLCFGHandler{},
	)
}
