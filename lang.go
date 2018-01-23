package main

import "fmt"

var lang = map[string]map[string]string{
	"tw": map[string]string{
		"":                           "未預期的錯誤: ",
		"btn_switch_lang_en":         "Switch to English",
		"btn_switch_lang_tw":         "切換至正體中文",
		"btn_inst_trademacro":        "執行 TradeMacro",
		"btn_inst_pob":               "執行 PoB",
		"btn_caps2ctrl":              "把 CapsLock 綁定成 Ctrl",
		"btn_restore_caps":           "還原 CapsLock",
		"btn_download_filters":       "下載物品過濾器",
		"btn_upload_filters":         "上傳物品過濾器",
		"btn_import_filter":          "下載指定的過濾器",
		"detect_ahk":                 "偵測系統是否已安裝 AutoHotkey",
		"ahk_path":                   "AutoHotkey 已安裝在 %s",
		"try_dl_ahk":                 "嘗試下載 AutoHotkey...",
		"inst_ahk":                   "執行 AutoHotkey 安裝程式",
		"err_dl_ahk":                 "下載 AutoHotkey 失敗，請重試一次",
		"dling_trademacro":           "下載 TradeMacro 中",
		"err_dl_tm":                  "下程 TradeMacro 失敗，請重試一次",
		"finish_trademacro":          "TradeMacro 已下載完成",
		"tm_executed":                "已成功執行 TradeMacro",
		"trademacro_no_matching_url": "無法在 GitHub 上找到 TradeMacro 的下載位址",
		"detect_pob":                 "偵測是否已下載 PoB...",
		"dling_pob":                  "下載 PoB 中",
		"err_dl_pob":                 "下載 PoB 失敗，請重試一次",
		"finish_pob":                 "PoB 已下載完成",
		"pob_executed":               "已成功執行 PoB",
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
		"btn_inst_trademacro":        "Run TradeMacro",
		"btn_inst_pob":               "Run PoB",
		"btn_caps2ctrl":              "Rebind CapsLock to Ctrl",
		"btn_restore_caps":           "Restore CapsLock",
		"btn_download_filters":       "Download item filters",
		"btn_upload_filters":         "Upload item filters",
		"btn_import_filter":          "Import from pastebin URL",
		"detect_ahk":                 "Detecting if AutoHotkey is installed",
		"ahk_path":                   "AutoHotkey is installed in %s",
		"try_dl_ahk":                 "Downloading AutoHotkey...",
		"inst_ahk":                   "Executing AutoHotkey install",
		"err_dl_ahk":                 "Failed to download AutoHotkey, you have to try again",
		"dling_trademacro":           "Downloading TradeMacro...",
		"err_dl_tm":                  "Failed to download TradeMacro, you have to try again",
		"finish_trademacro":          "Download success",
		"tm_executed":                "TradeMacro executed",
		"trademacro_no_matching_url": "Cannot find URL to download TradeMacro on GitHub",
		"detect_pob":                 "Detecting if PoB is downloaded...",
		"dling_pob":                  "Downloading PoB...",
		"err_dl_pob":                 "Failed to download PoB, you have to try again",
		"finish_pob":                 "Download success",
		"pob_executed":               "PoB executed",
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
