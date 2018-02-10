// +build windows

package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func fetchTradeMacroURL() (uri string, err error) {
	urlTradeMacro := "https://github.com/POE-TradeMacro/POE-TradeMacro/releases/latest"
	resp, err := http.Get(urlTradeMacro)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	ret := reGithubURL.Find(data)
	if ret == nil {
		err = errors.New("trademacro_no_matching_url")
		return
	}

	return "https://github.com/POE-TradeMacro/POE-TradeMacro/releases" + string(ret), nil
}

func init() {
	Handlers = append(Handlers, &TMHandler{})
}

type TMHandler struct {
	AbstractHandler
	ahkPath string
}

func (h *TMHandler) installAHK() (ok bool) {
	h.l.Log(L("detect_ahk"))

	k, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		`SOFTWARE\AutoHotkey`,
		registry.QUERY_VALUE)
	if err == nil {
		if s, _, err := k.GetStringValue("InstallDir"); err == nil {
			h.ahkPath = s
			h.l.Logf(L("ahk_path"), s)
			return true
		}
	}

	h.l.Log(L("try_dl_ahk"))
	uri := "https://autohotkey.com/download/ahk-install.exe"
	fn := "ahk-install.exe"
	if err := DL(uri, fn); err != nil {
		h.l.Log(L("err_dl_ahk"))
		return
	}

	h.l.Log(L("inst_ahk"))
	exec.Command(fn).Run()

	k, err = registry.OpenKey(
		registry.LOCAL_MACHINE,
		`SOFTWARE\AutoHotkey`,
		registry.QUERY_VALUE)
	if err == nil {
		if s, _, err := k.GetStringValue("InstallDir"); err == nil {
			h.ahkPath = s
			h.l.Logf(L("ahk_path"), s)
			return true
		}
	}

	h.l.Log(L("err_dl_ahk"))
	return false
}

func (h *TMHandler) InstallTradeMacro() {
	if !h.installAHK() {
		return
	}

	dir := filepath.Join(".", "POE-TradeMacro")
	tm := filepath.Join(dir, "Run_TradeMacro.ahk")

	if _, err := os.Stat(tm); err != nil {
		fn := "poe_trademacro.zip"
		// not found, download it
		h.l.Log(L("dling_trademacro"))

		uri, err := fetchTradeMacroURL()
		if err != nil {
			h.l.Log(L("err_dl_tm"))
			return
		}

		if err = DL(uri, fn); err != nil {
			h.l.Log(L("err_dl_tm"))
			return
		}
		defer os.Remove(fn)

		if _, err = Unzip(fn, dir); err != nil {
			h.l.Log(L("err_dl_tm"))
			return
		}

		h.l.Log(L("finish_trademacro"))
	}

	ahk := filepath.Join(h.ahkPath, "AutoHotkey.exe")

	exec.Command(ahk, tm).Start()
	h.l.Log(L("tm_executed"))
}

func (h *TMHandler) Handle(data interface{}) {
	go h.InstallTradeMacro()
}

func (h *TMHandler) Key() string {
	return "btn_inst_trademacro"
}
