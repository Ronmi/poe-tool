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

	return string(ret), nil
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
