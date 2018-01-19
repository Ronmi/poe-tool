package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

var reGithubURL *regexp.Regexp

func init() {
	reGithubURL = regexp.MustCompile(
		`/download/[a-zA-Z0-9./-]+\.zip`,
	)
}

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

func installAHK() (err error) {
	uri := "https://autohotkey.com/download/ahk-install.exe"
	fn := "ahk-install.exe"
	return DL(uri, fn)
}

func InstallTradeMacro() (err error) {
	if err = installAHK(); err != nil {
		return
	}

	dir := filepath.Join(".", "POE-TradeMacro")
	fn := "poe_trademacro.zip"

	uri, err := fetchTradeMacroURL()
	if err != nil {
		return
	}

	if err = DL(uri, fn); err != nil {
		return
	}
	defer os.Remove(fn)

	_, err = Unzip(fn, dir)
	return
}
