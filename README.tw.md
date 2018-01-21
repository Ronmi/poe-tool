一些流氓按到國際服的小工具

# 主要功能

* 下載 `AHK` 與 `POE-TradeMacro`
* 下載 `Path of Building`
* 把物品過濾器備份到 pastebin
* 從 pastebin 把之前備份的過濾器載回來
* 透過 pastebin 鏈結，下載別人上傳的過濾器

# 使用方式

### 申請 pastebin 帳號

你需要有 [pastebin](https://pastebin.com) 的帳號才能使用 poe-tool

### 取得 dev key

登入 pastebin 後，到 [api page](https://pastebin.com/api#1) 取得你的 dev key (一串英文數字)

### 建立設定檔

用記事本 (或其他純文字編輯器) 建立一個 `poe-tool.yml` 檔案:

```yaml
devkey: "你的 dev key"
username: "你的 pastebin 帳號"
password: "你的 pastebin 密碼"
```

### 編譯主程式

C 槽需要大約 1.5G 的空間

* 下載並安裝 [MSYS2](http://www.msys2.org/)，使用預設的設定就好
* 安裝好之後，應該會有一個類似命令提示字元的黑色視窗跑出來
* 在那個黑視窗裡輸入 `/mingw64/mingw64 bash -c "curl -sSL https://raw.githubusercontent.com/Ronmi/poe-tool/master/inst.mingw.sh|bash"` 然後按 enter 執行
* 慢慢等
* 把 `C:\msys2\mingw64\bin` 這個資料夾裡的所有 DLL 檔和 `poe-tool.exe` 備份到其他位置
* 現在可以把 MSYS2 移除來節省你的硬碟空間

### 執行

把設定檔跟剛才編好的主程式放在一起，然後執行 `poe-tool.exe`

# 軟體授權

WTFPL
