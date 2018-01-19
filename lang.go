package main

import "fmt"

var lang = map[string]map[string]string{
	"tw": map[string]string{
		"":                           "未預期的錯誤: ",
		"btn_switch_lang_en":         "Switch to English",
		"btn_switch_lang_tw":         "切換至正體中文",
		"btn_inst_trademacro":        "下載 TradeMacro/AHK",
		"btn_inst_pob":               "下載 PoB",
		"btn_caps2ctrl":              "把 CapsLock 綁定成 Ctrl",
		"btn_restore_caps":           "還原 CapsLock",
		"btn_download_filters":       "下載物品過濾器",
		"btn_upload_filters":         "上傳物品過濾器",
		"btn_import_filter":          "下載指定的過濾器",
		"dling_trademacro":           "下載 TradeMacro 中",
		"finish_trademacro":          "下載完成",
		"trademacro_no_matching_url": "無法在 GitHub 上找到 TradeMacro 的下載位址",
		"dling_pob":                  "下載 PoB 中",
		"finish_pob":                 "下載完成",
		"bind_caps":                  "已將 CapsLock 綁定為 Ctrl",
		"restore_caps":               "已將 CapsLock 還原",
		"log_state_begin":            "正在處理 %[1]s",
		"log_state_ok":               "%[1]s 完成: %[2]s",
		"log_state_fail":             "%[1]s 失敗: %[2]s",
		"cannot_get_user_dir":        "找不到過濾器的目錄!?",
		"ask_pastebin_key":           "你必須提供 Pastebin 的 API Key 才能使用過濾器上傳下載功能",
		"not_implemented":            "目前尚未提供此功能"},
	"en": map[string]string{
		"":                           "Unexpected error: ",
		"btn_switch_lang_en":         "Switch to English",
		"btn_switch_lang_tw":         "切換至正體中文",
		"btn_inst_trademacro":        "Download TradeMacro/AHK",
		"btn_inst_pob":               "Download PoB",
		"btn_caps2ctrl":              "Rebind CapsLock to Ctrl",
		"btn_restore_caps":           "Restore CapsLock",
		"btn_download_filters":       "Download item filters",
		"btn_upload_filters":         "Upload item filters",
		"btn_import_filter":          "Import from pastebin URL",
		"dling_trademacro":           "Downloading TradeMacro...",
		"finish_trademacro":          "Download success",
		"trademacro_no_matching_url": "Cannot find URL to download TradeMacro on GitHub",
		"dling_pob":                  "Downloading PoB...",
		"finish_pob":                 "Download success",
		"bind_caps":                  "Successfully binding CapsLock to Ctrl",
		"restore_caps":               "CapsLock functionality restored",
		"log_state_begin":            "Processing %[1]s",
		"log_state_ok":               "%[1]s success: %[2]s",
		"log_state_fail":             "%[1]s failed: %[2]s",
		"cannot_get_user_dir":        "No filters found!?",
		"ask_pastebin_key":           "You have to provide Pastebin API Keys to upload/download item filters",
		"not_implemented":            "This feature is not implemented yet",
	},
}

var locale string

func LErr(err error) string {
	l, ok := lang[locale]
	if !ok {
		l = lang["tw"]
	}

	msg, ok := l[err.Error()]
	if !ok {
		msg = l[""] + err.Error()
	}

	return msg
}

func L(str string) string {
	l, ok := lang[locale]
	if !ok {
		l = lang["tw"]
	}

	msg, ok := l[str]
	if !ok {
		panic(fmt.Errorf("no lang setting for [%s]", str))
	}

	return msg
}
