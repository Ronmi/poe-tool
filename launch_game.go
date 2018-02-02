// +build windows

package main

import (
	"os/exec"

	"github.com/gotk3/gotk3/gtk"
	"golang.org/x/sys/windows/registry"
)

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
		registry.LOCAL_MACHINE,
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

	h.prog, c.cfg.GameProg = s, s
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
		if resp := dlg.Run(); resp == int(gtk.RESPONSE_CANCEL) {
			return
		}

		h.prog = dlg.GetFilename()
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
