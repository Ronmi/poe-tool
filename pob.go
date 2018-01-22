package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

type POBHandler struct {
	AbstractHandler
}

func (h *POBHandler) Handle(data interface{}) {
	h.l.Log(L("dling_pob"))
	go func() {
		err := InstallPoB()
		if err != nil {
			h.l.Log(LErr(err))
		} else {
			h.l.Log(L("finish_pob"))
		}
	}()
}

func (h *POBHandler) Key() string {
	return "btn_inst_pob"
}
