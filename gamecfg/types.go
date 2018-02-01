package gamecfg

import (
	"bufio"
	"bytes"
	"io"
	"reflect"
)

var bom = []byte{0xef, 0xbb, 0xbf}

func write(w io.Writer, data []byte, err error) (e error) {
	if err != nil {
		return err
	}
	_, e = w.Write(data)
	return
}

func Merge(to, from map[string][]string) (ret map[string][]string) {
	ret = to
	for k, v := range from {
		ret[k] = v
	}
	return
}

func Write(w io.Writer, cfg map[string][]string) (err error) {
	// write BOM header
	err = write(w, bom, err)
	for k, vals := range cfg {
		err = write(w, []byte(k), err)
		err = write(w, []byte{'\n'}, err)
		for _, l := range vals {
			err = write(w, []byte(l), err)
			err = write(w, []byte{'\n'}, err)
		}
	}

	return
}

func Read(reader io.Reader) (cfg map[string][]string, err error) {
	r := bufio.NewReader(reader)
	cfg = make(map[string][]string)

	begin := true
	var cursec string
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return cfg, nil
			}

			return cfg, err
		}
		if begin {
			// detect BOM
			if reflect.DeepEqual(line[:3], bom) {
				line = line[3:]
			}
			begin = false
		}

		l := bytes.TrimSpace(line)
		if l[0] == '[' {
			// section
			cursec = string(l)
			cfg[cursec] = make([]string, 0)
			continue
		}

		cfg[cursec] = append(cfg[cursec], string(l))
	}
}
