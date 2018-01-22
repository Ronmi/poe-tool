package main

import "testing"

func TestLangMapping(t *testing.T) {
	expect := lang["tw"]

	for x, m := range lang {
		if x == "tw" {
			continue
		}

		for k, _ := range expect {
			if m[k] == "" {
				t.Errorf("key %s not translated to lang %s", k, x)
			}
		}

		for k, _ := range m {
			if expect[k] == "" {
				t.Errorf("key %s in lang %s is not used", k, x)
			}
		}
	}
}
