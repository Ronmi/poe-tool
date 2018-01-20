package main

import (
	"errors"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/Ronmi/pastebin"
)

const filterSuffix = ".poe.filter"

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

type remote struct {
	Key  string
	Name string
}

type filterHnd struct {
	DevKey  string
	UserKey string
	api     *pastebin.API
}

func FilterHnd(cfg *ConfigFile) *filterHnd {
	ret := &filterHnd{
		DevKey:  cfg.DevKey,
		UserKey: cfg.UserKey,
		api:     &pastebin.API{Key: cfg.DevKey},
	}

	if ret.UserKey == "" {
		if cfg.Username == "" || cfg.Password == "" {
			panic(errors.New("you must provide pastebin username/password or user key"))
		}

		k, err := ret.api.UserKey(cfg.Username, cfg.Password)
		if err != nil {
			panic(err)
		}

		ret.UserKey = k
	}

	return ret
}

func (h *filterHnd) dlRemoteFile(dir string, remote remote) (err error) {
	data, err := h.api.UserPaste(h.UserKey, remote.Key)
	if err != nil {
		return
	}
	fn := filepath.Join(dir, remote.Name)
	return ioutil.WriteFile(fn, data, 0664)
}

func (h *filterHnd) rmRemoteFile(key string) (err error) {
	return h.api.Delete(h.UserKey, key)
}

func (h *filterHnd) listRemoteFilters() (ret []remote, err error) {
	files, err := h.api.List(h.UserKey, 1000)
	if err != nil {
		return
	}

	ret = make([]remote, 0)
	for _, f := range files {
		if !strings.HasSuffix(f.Title, filterSuffix) {
			continue
		}

		ret = append(ret, remote{
			Key:  f.Key,
			Name: f.Title[:len(f.Title)-len(filterSuffix)] + ".filter",
		})
	}

	return
}

func (h *filterHnd) listLocalFilters(dir string) (ret []string, err error) {
	return filepath.Glob(filepath.Join(dir, "*.filter"))
}

func (h *filterHnd) ulLocalFilter(fn string) (uri string, err error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return
	}

	title := filepath.Base(fn)
	p := &pastebin.Paste{
		Title:   title[:len(title)-len(".filter")] + ".poe.filter",
		Content: string(data),
		UserKey: h.UserKey,
	}

	return h.api.Post(p)
}

// leaving Filename, Key and Finished in zero value to indicate unexpected error
type DLState struct {
	Filename string
	Key      string
	Finished bool
	Err      error
}

func (h *filterHnd) DownloadFilters(dir string) (ch chan DLState) {
	ch = make(chan DLState)

	go func(ch chan DLState) {
		defer close(ch)

		remotes, err := h.listRemoteFilters()
		if err != nil {
			ch <- DLState{Err: err}
			return
		}

		for _, x := range remotes {
			st := DLState{
				Filename: x.Name,
				Key:      x.Key,
			}
			ch <- st

			st.Err = h.dlRemoteFile(dir, x)
			st.Finished = true
			ch <- st
		}
	}(ch)

	return
}

func (h *filterHnd) UploadFilters(dir string) (ch chan DLState) {
	ch = make(chan DLState)

	go func(ch chan DLState) {
		defer close(ch)

		files, err := h.listLocalFilters(dir)
		if err != nil {
			ch <- DLState{Err: err}
			return
		}

		remotes, err := h.listRemoteFilters()
		if err != nil {
			ch <- DLState{Err: err}
			return
		}

		m := make(map[string]bool)
		for _, x := range files {
			fn := filepath.Base(x)
			m[fn] = true
		}

		for _, x := range remotes {
			if !m[x.Name] {
				continue
			}

			if err := h.rmRemoteFile(x.Key); err != nil {
				ch <- DLState{Err: err}
				return
			}
		}

		for _, x := range files {
			st := DLState{
				Filename: x,
			}
			ch <- st

			uri, err := h.ulLocalFilter(x)
			if err == nil {
				st.Key = uri
			}
			st.Err = err
			st.Finished = true
			ch <- st
		}
	}(ch)

	return
}
