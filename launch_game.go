// +build windows

package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gotk3/gotk3/gtk"
	"golang.org/x/sys/windows/registry"
)

func init() {
	// forcing schema souces
	p, _ := filepath.Abs(".")
	os.Setenv("XDG_DATA_DIRS", p)
}

type LaunchHandler struct {
	AbstractHandler
	prog string
}

func (h *LaunchHandler) Init(c buttonInitParam) {
	h.AbstractHandler.Init(c)

	if c.cfg.GameProg != "" {
		h.prog = c.cfg.GameProg
		return
	}

	k, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\GrindingGearGames\Path of Exile`,
		registry.QUERY_VALUE)
	if err != nil {
		h.l.Log(L("not_installed"))
		return
	}

	s, _, err := k.GetStringValue("InstallLocation")
	if err != nil {
		h.l.Log(L("not_installed"))
		return
	}
	s = filepath.Join(s, "PathOfExile_x64.exe")

	h.prog, c.cfg.GameProg = s, s
	c.cfg.Save()
}

func (h *LaunchHandler) Handle(data interface{}) {
	h.l.Log(L("btn_launch"))

	if h.prog == "" {
		h.askProg()
		return
	}

	if err := exec.Command(h.prog).Start(); err != nil {
		h.l.Log(LErr(err))
	}
}

func (h *LaunchHandler) askProg() {
	f, err := gtk.FileFilterNew()
	if err != nil {
		h.l.Log(LErr(err))
		return
	}
	f.AddPattern("*.exe")
	f.SetName(L("filter_exe"))

	dlg, err := gtk.FileChooserDialogNewWith2Buttons(
		L("choose_game_prog"),
		h.root,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		L("ok"),
		gtk.RESPONSE_OK,
		L("cancel"),
		gtk.RESPONSE_CANCEL)
	if err != nil {
		h.l.Log(LErr(err))
		return
	}

	dlg.AddFilter(f)
	dlg.ShowAll()

	go func() {
		defer dlg.Destroy()
		if resp := dlg.Run(); resp != int(gtk.RESPONSE_OK) {
			return
		}

		h.prog = dlg.GetFilename()
		if h.prog == "" {
			return
		}
		h.cfg.GameProg = h.prog
		h.cfg.Save()
		h.Handle(nil)
	}()
}

func (h *LaunchHandler) Key() string {
	return "btn_launch"
}

func init() {
	Handlers = append(Handlers, &LaunchHandler{})
}
