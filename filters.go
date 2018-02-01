package main

import (
	"bytes"
	"errors"
	"log"
	"os/user"
	"path"
	"path/filepath"

	"github.com/Ronmi/pastebin"
)

const filterName = "poe.filter"

func GetFilterDir() (dir string, err error) {
	u, err := user.Current()
	if err != nil {
		return
	}

	if u.HomeDir == "" {
		return "", errors.New("cannot_get_user_dir")
	}

	return filepath.Join(
		u.HomeDir,
		"Documents",
		"My Games",
		"Path of Exile",
	), nil
}

type filterHnd struct {
	DevKey  string
	UserKey string
	api     *pastebin.API
}

var hnd *filterHnd

func FilterHnd(cfg *ConfigFile) *filterHnd {
	if hnd == nil {
		hnd = &filterHnd{
			DevKey:  cfg.DevKey,
			UserKey: cfg.UserKey,
			api:     &pastebin.API{Key: cfg.DevKey},
		}

		if hnd.UserKey == "" {
			if cfg.Username == "" || cfg.Password == "" {
				panic(errors.New("you must provide pastebin username/password or user key"))
			}

			k, err := hnd.api.UserKey(cfg.Username, cfg.Password)
			if err != nil {
				panic(err)
			}

			hnd.UserKey = k
		}
	}

	return hnd
}

func (h *filterHnd) rmRemoteFile(key string) (err error) {
	return h.api.Delete(h.UserKey, key)
}

func (h *filterHnd) findRemoteFilter() (key string, err error) {
	files, err := h.api.List(h.UserKey, 1000)
	if err != nil {
		return
	}

	for _, f := range files {
		log.Printf("%+v", f)
		if f.Title == filterName {
			return f.Key, nil
		}
	}

	return
}

func (h *filterHnd) listLocalFilters(dir string) (ret []string, err error) {
	return filepath.Glob(filepath.Join(dir, "*.filter"))
}

// leaving Filename, Key and Finished in zero value to indicate unexpected error
type DLState struct {
	Filename string
	Key      string
	Finished bool
	Err      error
}

func (h *filterHnd) DownloadFilters(l Logger, key string) {
	dir, err := GetFilterDir()
	if err != nil {
		l.Log(L("cannot_get_user_dir"))
		return
	}

	data, err := h.api.UserPaste(h.UserKey, key)
	if err != nil {
		l.Log(LErr(err))
		return
	}
	d := &downloader{
		R:   bytes.NewBuffer(data),
		Dir: dir,
		L:   l,
	}
	if err = d.Extract(); err != nil {
		l.Log(LErr(err))
	}
}

func (h *filterHnd) UploadFilters(l Logger) {
	dir, err := GetFilterDir()
	if err != nil {
		l.Log(L("cannot_get_user_dir"))
		return
	}
	files, err := h.listLocalFilters(dir)
	if err != nil {
		l.Log(LErr(err))
		return
	}
	if len(files) < 1 {
		return
	}

	key, err := h.findRemoteFilter()
	if err != nil {
		l.Log(LErr(err))
		return
	}
	if err := h.rmRemoteFile(key); err != nil {
		l.Log(LErr(err))
		return
	}
	u := &uploader{L: l}
	if err := u.Tar(files); err != nil {
		l.Log(LErr(err))
	}

	p := &pastebin.Paste{
		Title:   filterName,
		Content: string(u.Data()),
		UserKey: h.UserKey,
	}

	if _, err := h.api.Post(p); err != nil {
		l.Log(LErr(err))
	}
}

func init() {
	Handlers = append(
		Handlers,
		&ULFHandler{},
		&DLFHandler{},
		&ImportHandler{},
	)
}

type DLFHandler struct {
	AbstractHandler
	h *filterHnd
}

func (h *DLFHandler) Init(l Logger, cfg *ConfigFile) {
	h.AbstractHandler.Init(l, cfg)
	h.h = FilterHnd(cfg)
}

func (h *DLFHandler) Handle(data interface{}) {
	go func() {
		key, err := h.h.findRemoteFilter()
		if err != nil {
			h.l.Log(LErr(err))
			return
		}
		if key == "" {
			h.l.Log(L("remote_filter_404"))
			return
		}
		h.h.DownloadFilters(h.l, key)
	}()
}

func (h *DLFHandler) Key() string {
	return "btn_download_filters"
}

type ULFHandler struct {
	AbstractHandler
	h *filterHnd
}

func (h *ULFHandler) Init(l Logger, cfg *ConfigFile) {
	h.AbstractHandler.Init(l, cfg)
	h.h = FilterHnd(cfg)
}

func (h *ULFHandler) Handle(data interface{}) {
	go h.h.UploadFilters(h.l)
}

func (h *ULFHandler) Key() string {
	return "btn_upload_filters"
}

type ImportHandler struct {
	AbstractHandler
	h *filterHnd
}

func (h *ImportHandler) Init(l Logger, cfg *ConfigFile) {
	h.AbstractHandler.Init(l, cfg)
	h.h = FilterHnd(cfg)
}

func (h *ImportHandler) Handle(data interface{}) {
	uri := InputURL(h.l)
	go h.h.DownloadFilters(h.l, path.Base(uri))
}

func (h *ImportHandler) Key() string {
	return "btn_import_filter"
}
