#!/bin/bash

export dir="${PWD}/travis"
export deploy="${dir}/build"
export PKG_CONFIG_PATH="${dir}/mingw64/lib/pkgconfig"
export CGO_ENABLED=1
export GOOS=windows
export CC=x86_64-w64-mingw32-gcc
export C_INCLUDE_PATH="${dir}/mingw64/include/gtk-3.0:${dir}/mingw64/include/atk-1.0:${dir}/mingw64/include:${dir}/mingw64/include/cairo:${dir}/mingw64/include/gdk-pixbuf-2.0:${dir}/mingw64/include/glib-2.0:${dir}/mingw64/lib/glib-2.0/include:${dir}/mingw64/include/pango-1.0:${dir}/mingw64/include/pixman-1:${dir}/mingw64/include/freetype2:${dir}/mingw64/include/libpng16:${dir}/mingw64/include/harfbuzz"
export LD_LIBRARY_PATH="${dir}/mingw64/lib"

pkgs="mingw-w64-x86_64-adwaita-icon-theme-3.26.1-1  mingw-w64-x86_64-atk-2.26.1-1  mingw-w64-x86_64-bzip2-1.0.6-6  mingw-w64-x86_64-ca-certificates-20170211-2 \
    mingw-w64-x86_64-cairo-1.15.10-1  mingw-w64-x86_64-expat-2.2.5-1  mingw-w64-x86_64-fontconfig-2.12.6-1  mingw-w64-x86_64-freeglut-3.0.0-4 \
    mingw-w64-x86_64-freetype-2.9-1  mingw-w64-x86_64-gcc-libs-7.2.0-2  mingw-w64-x86_64-gdk-pixbuf2-2.36.11-1  mingw-w64-x86_64-gettext-0.19.8.1-2 \
    mingw-w64-x86_64-glib2-2.54.2-1  mingw-w64-x86_64-gmp-6.1.2-1  mingw-w64-x86_64-graphite2-1.3.10-1  mingw-w64-x86_64-harfbuzz-1.7.4-1 \
    mingw-w64-x86_64-hicolor-icon-theme-0.15-2  mingw-w64-x86_64-jasper-2.0.14-1  mingw-w64-x86_64-json-glib-1.4.2-2  mingw-w64-x86_64-libcroco-0.6.12-1 \
    mingw-w64-x86_64-libepoxy-1.4.3-1  mingw-w64-x86_64-libffi-3.2.1-4  mingw-w64-x86_64-libiconv-1.15-1  mingw-w64-x86_64-libjpeg-turbo-1.5.3-1 \
    mingw-w64-x86_64-libpng-1.6.34-1  mingw-w64-x86_64-librsvg-2.40.20-1  mingw-w64-x86_64-libsystre-1.0.1-4  mingw-w64-x86_64-libtasn1-4.12-1 \
    mingw-w64-x86_64-libtiff-4.0.9-1  mingw-w64-x86_64-libtre-git-r128.6fb7206-2  mingw-w64-x86_64-libwinpthread-git-5.0.0.4850.d1662dc7-1 \
    mingw-w64-x86_64-libxml2-2.9.7-1  mingw-w64-x86_64-lzo2-2.10-1  mingw-w64-x86_64-mpc-1.1.0-1  mingw-w64-x86_64-mpfr-4.0.0-1 \
    mingw-w64-x86_64-ncurses-6.0.20170916-1  mingw-w64-x86_64-openssl-1.0.2.n-1  mingw-w64-x86_64-p11-kit-0.23.9-1  mingw-w64-x86_64-pango-1.40.11-1 \
    mingw-w64-x86_64-pcre-8.41-1  mingw-w64-x86_64-pixman-0.34.0-3  mingw-w64-x86_64-python2-2.7.14-3  mingw-w64-x86_64-readline-7.0.003-1 \
    mingw-w64-x86_64-shared-mime-info-1.9-1  mingw-w64-x86_64-tcl-8.6.8-1  mingw-w64-x86_64-termcap-1.3.1-3  mingw-w64-x86_64-tk-8.6.8-1 \
    mingw-w64-x86_64-wineditline-2.201-1  mingw-w64-x86_64-xz-5.2.3-1  mingw-w64-x86_64-zlib-1.2.11-1  mingw-w64-x86_64-gtk2-2.24.32-1 \
    mingw-w64-x86_64-gtk3-3.22.26-1"

function inst {
    curl -sSL "http://repo.msys2.org/mingw/x86_64/${1}-any.pkg.tar.xz" | tar Jxvf - -C travis
}

if [[ $1 == "dep" ]]
then
    echo "preparing mingw-gtk"
    if [[ ! -d "${dir}/mingw64/bin" ]]
    then
	for x in $pkgs
	do
	    inst "$x"
	done
    fi

    exit 0
fi

sed -i "s#prefix=/mingw64#prefix=${dir}/mingw64#" travis/mingw64/lib/pkgconfig/*.pc

go build -v -ldflags="-H=windowsgui"

mv poe-tool.exe "$deploy"
cp poe-tool.yml "$deploy"

dlls="libatk-1.0-0.dll libgdk-3-0.dll libpangocairo-1.0-0.dll \
libbz2-1.dll libgio-2.0-0.dll libpangoft2-1.0-0.dll \
libcairo-2.dll libglib-2.0-0.dll libpangowin32-1.0-0.dll \
libcairo-gobject-2.dll libgmodule-2.0-0.dll libpcre-1.dll \
libepoxy-0.dll libgobject-2.0-0.dll libpixman-1-0.dll \
libexpat-1.dll libgraphite2.dll libpng16-16.dll \
libffi-6.dll libgtk-3-0.dll libstdc++-6.dll \
libfontconfig-1.dll libharfbuzz-0.dll libwinpthread-1.dll \
libfreetype-6.dll libiconv-2.dll zlib1.dll \
libgcc_s_seh-1.dll libintl-8.dll \
libgdk_pixbuf-2.0-0.dll libpango-1.0-0.dll"

for x in $dlls
do
    cp "${dir}/mingw64/bin/${x}" "$deploy"
done

tar czf poe-tool.tar.gz -C "$deploy" --exclude=.gitignore .
