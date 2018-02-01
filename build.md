Although complex and time consuming, it is much safer to use your own binary instead of downloading pre-built binaries.

# Howto

You'll need a fast network connection and ~1.5Gb disk space in C driver (with default settings)

* Grab and install [MSYS2](http://www.msys2.org/), default settings should be fairly enough.
* After it is installed, there should be a window poped up, which looks similair with windows command prompt, but more colorful.
* Type `/mingw64.exe bash -c "curl -sSL https://raw.githubusercontent.com/Ronmi/poe-tool/master/inst.mingw.sh|bash"` in it, and press enter.
* Wait till it complete.
* Move `C:\msys2\go\bin\poe-tool.exe` and all DLL files from `C:\msys2\mingw64\bin` to somewhere else.
* Now it's safe to uninstall MSYS2 to save your disk space.
