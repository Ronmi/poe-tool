package main

import (
	"io"
	"net/http"
	"os"
)

func DL(uri, dest string) (err error) {
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(dest)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}
