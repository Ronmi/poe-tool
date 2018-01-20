#!/bin/bash

export GOROOT=/mingw64/lib/go
export GOPATH=/mingw64
export PATH="$PATH:/mingw64/bin"
export PKG_CONFIG_PATH=/mingw64/lib/pkgconfig

pacman -Ss
pacman -S --needed --noconfirm mingw-w64-x86_64-go mingw-w64-x86_64-gtk3 mingw-w64-x86_64-gtk2 git mingw-w64-x86_64-toolchain base-devel

go get -v github.com/Ronmi/poe-tool
