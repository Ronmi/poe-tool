package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

var reGithubExeURL *regexp.Regexp

func init() {
	reGithubExeURL = regexp.MustCompile(
		`/download/[a-zA-Z0-9./-]+\.exe`,
	)

	Handlers = append(
		Handlers,
		&UncapHandler{},
		&RecapHandler{})
}

func fetchUncapURL() (uri string, err error) {
	urlUncap := `https://github.com/susam/uncap/releases/latest`
	resp, err := http.Get(urlUncap)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	ret := reGithubExeURL.Find(data)
	if ret == nil {
		err = errors.New("trademacro_no_matching_url")
		return
	}

	return "https://github.com/susam/uncap/releases" + string(ret), nil
}

func InstallUncap() (err error) {
	fn := `uncap.exe`
	if _, err = os.Stat(fn); err == nil {
		return
	}

	uri, err := fetchUncapURL()
	if err != nil {
		return
	}

	return DL(uri, fn)
}

func Caps2Ctrl() error {
	return exec.Command("uncap.exe", "0x14:0xa2").Run()
}

func RestoreCaps() error {
	return exec.Command("uncap.exe", "-k").Run()
}

type UncapHandler struct {
	AbstractHandler
}

func (h *UncapHandler) Handle(data interface{}) {
	h.l.Log(L("btn_caps2ctrl"))
	go func() {
		if err := InstallUncap(); err != nil {
			h.l.Log(LErr(err))
			return
		}
		if err := Caps2Ctrl(); err != nil {
			h.l.Log(LErr(err))
			return
		}
		h.l.Log(L("bind_caps"))
	}()
}

func (h *UncapHandler) Key() string {
	return "btn_caps2ctrl"
}

type RecapHandler struct {
	AbstractHandler
}

func (h *RecapHandler) Handle(data interface{}) {
	h.l.Log(L("btn_restore_caps"))
	go func() {
		if err := InstallUncap(); err != nil {
			h.l.Log(LErr(err))
			return
		}
		if err := RestoreCaps(); err != nil {
			h.l.Log(LErr(err))
			return
		}
		h.l.Log(L("restore_caps"))
	}()
}

func (h *RecapHandler) Key() string {
	return "btn_restore_caps"
}
