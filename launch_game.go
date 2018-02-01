// +build windows

package main

import (
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

type LaunchHandler struct {
	AbstractHandler
	prog string
}

func (h *LaunchHandler) Init(l Logger, cfg *ConfigFile) {
	h.AbstractHandler.Init(l, cfg)

	k, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		`Software\GrindingGearGames\Path of Exile`,
		registry.QUERY_VALUE)
	if err != nil {
		// not installed, disable it
		h.GetObject().SetSensitive(false)
		h.l.Log(L("not_installed"))
		return
	}

	s, _, err := k.GetStringValue("InstallLocation")
	if err != nil {
		h.l.Log(L("not_installed"))
		return
	}
	h.prog = s
}

func (h *LaunchHandler) Handle(data interface{}) {
	h.l.Log(L("btn_launch"))

	exec.Command(h.prog).Start()
}

func (h *LaunchHandler) Key() string {
	return "btn_launch"
}

func init() {
	Handlers = append(Handlers, &LaunchHandler{})
}
