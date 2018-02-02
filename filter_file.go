package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func failULDL(l Logger, fn string, err error) {
	l.Logf(L("log_state_fail"), fn, err)
}
func doneULDL(l Logger, fn, msg string) {
	l.Logf(L("log_state_ok"), fn, msg)
}

type uploader struct {
	buf bytes.Buffer
	L   Logger
}

func (h *uploader) Data() []byte {
	return h.buf.Bytes()
}

func (h *uploader) Tar(files []string) (err error) {
	b := base64.NewEncoder(
		base64.StdEncoding,
		&h.buf)
	defer b.Close()
	z := gzip.NewWriter(b)
	defer z.Close()
	t := tar.NewWriter(z)

	for _, fn := range files {
		if err = h.write(t, fn); err != nil {
			t.Close()
			return
		}
	}

	return t.Close()
}

func (h *uploader) write(t *tar.Writer, full string) (err error) {
	fn := filepath.Base(full)
	h.L.Logf(L("log_state_begin"), fn)

	info, err := os.Stat(full)
	if err != nil {
		failULDL(h.L, fn, err)
		return
	}

	data, err := ioutil.ReadFile(full)
	if err != nil {
		failULDL(h.L, fn, err)
		return
	}

	hdr := &tar.Header{
		Name: fn,
		Mode: int64(info.Mode()),
		Size: info.Size(),
	}
	if err = t.WriteHeader(hdr); err != nil {
		failULDL(h.L, fn, err)
		return
	}

	_, err = t.Write(data)
	if err != nil {
		failULDL(h.L, fn, err)
	} else {
		doneULDL(h.L, fn, strconv.FormatInt(hdr.Size, 10))
	}
	return
}

type downloader struct {
	R   io.Reader
	Dir string
	L   Logger
}

func (d *downloader) Extract() (err error) {
	r, err := gzip.NewReader(base64.NewDecoder(
		base64.StdEncoding, d.R))
	if err != nil {
		return
	}
	t := tar.NewReader(r)

	for {
		h, err := t.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if err = d.read(t, h); err != nil {
			return err
		}
	}
}

func (d *downloader) read(t *tar.Reader, h *tar.Header) (err error) {
	d.L.Logf(L("log_state_begin"), h.Name)
	f, err := os.Create(filepath.Join(d.Dir, h.Name))
	if err != nil {
		failULDL(d.L, h.Name, err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, t)
	if err != nil {
		failULDL(d.L, h.Name, err)
	} else {
		doneULDL(d.L, h.Name, strconv.FormatInt(h.Size, 10))
	}
	return
}
