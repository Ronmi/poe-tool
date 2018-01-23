package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func fetchPoBURL() (uri string, err error) {
	urlPoB := "https://github.com/Openarl/PathOfBuilding/releases/latest"
	resp, err := http.Get(urlPoB)
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

	return "https://github.com/Openarl/PathOfBuilding/releases" + string(ret), nil
}

func InstallPoB() (err error) {
	dir := filepath.Join(".", "PathOfBuilding")
	fn := "pob.zip"

	uri, err := fetchPoBURL()
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

var reGithubURL *regexp.Regexp

func init() {
	reGithubURL = regexp.MustCompile(
		`/download/[a-zA-Z0-9./-]+\.zip`,
	)

	Handlers = append(Handlers, &POBHandler{})
}

type POBHandler struct {
	AbstractHandler
}

func (h *POBHandler) Handle(data interface{}) {
	go func() {
		fn := filepath.Join("PathOfBuilding", "Path of Building.exe")
		h.l.Log(L("detect_pob"))
		if _, err := os.Stat(fn); err != nil {
			h.l.Log(L("dling_pob"))
			err := InstallPoB()
			if err != nil {
				h.l.Log(L("err_dl_pob"))
				return
			}

			h.l.Log(L("finish_pob"))
		}

		exec.Command(fn).Start()
		h.l.Log(L("pob_executed"))
	}()
}

func (h *POBHandler) Key() string {
	return "btn_inst_pob"
}
