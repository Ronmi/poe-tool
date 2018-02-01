package main

func init() {
	Handlers = append(Handlers, &LangTWHandler{}, &LangENHandler{})
}

type LangTWHandler struct {
	AbstractHandler
}

func (h *LangTWHandler) Handle(data interface{}) {
	switchLocale("tw")
	h.l.Log(L("btn_switch_lang_tw"))
	h.cfg.Locale = "tw"
	if err := h.cfg.Save(); err != nil {
		h.l.Log(L("failed_update_cfg"))
	}
}

func (h *LangTWHandler) Key() string {
	return "btn_switch_lang_tw"
}

type LangENHandler struct {
	AbstractHandler
}

func (h *LangENHandler) Handle(data interface{}) {
	switchLocale("en")
	h.l.Log(L("btn_switch_lang_en"))
	h.cfg.Locale = "en"
	if err := h.cfg.Save(); err != nil {
		h.l.Log(L("failed_update_cfg"))
	}
}

func (h *LangENHandler) Key() string {
	return "btn_switch_lang_en"
}
