[中文說明](https://github.com/Ronmi/poe-tool/blob/master/README.tw.md)

Some simple tool for PoE international realm

# Features

* Download `AHK` installer and `POE-TradeMacro`.
* Download `Path of Building`.
* Upload your item filters to pastebin.
* Download previously uploaded filters.
* Import filters via pastebin link.

# How to use

### Apply pastebin account

First, you need a [pastebin](https://pastebin.com) account.

### Get your dev key

After logging in, visit [api page](https://pastebin.com/api#1) for your dev key.

### Prepare config file

create `poe-tool.yml` with your favorite plain text editor (notepad for example)

```yaml
devkey: "your dev key"
username: "your pastebin account"
password: "your pastebin password"
```

### Obtain binaries

Build on you own is suggested! See [build instructions](https://github.com/Ronmi/poe-tool/blob/master/build.md).

You may refer to release page to download prebuilt binaries.

### Run it!

Put the config file along with the binaries, and run `poe-tool.exe`

# License

WTFPL
