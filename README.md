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

### Get your user key (optional, recommended)

You'll need a user key to upload/delete/list your pastes.

Visits [this page](https://pastebin.com/api/api_user_key.html) to generate one after logging in to pastebin.com.

**CAUSION**: Previous user key exipres when you generate new one.

### Prepare config file

create `poe-tool.yml` with your favorite plain text editor (notepad for example)

```yaml
devkey: "your dev key"
userkey: "your user key"
```

If you are brave enough, use paste account and password, let `poe-tool` allocate a new user key for you:

```yaml
devkey: "your dev key"
username: "your pastebin account"
password: "your pastebin password"
```

All supported configuration items:

- `devkey`: pastebin dev key.
- `userkey`: pasetbin user key. It has precedence over `account` and `password`.
- `username`: pastebin account name.
- `password`: pastebin password.
- `locale`: language selection. Currently only `tw` and `en` are supported. Default to `tw`.

### Obtain binaries

Build on you own is suggested! See [build instructions](https://github.com/Ronmi/poe-tool/blob/master/build.md).

You may refer to release page to download prebuilt binaries.

### Run it!

Put the config file along with the binaries, and run `poe-tool.exe`

# License

WTFPL
