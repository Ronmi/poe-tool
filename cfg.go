package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type ConfigFile struct {
	DevKey   string
	UserKey  string
	Username string
	Password string
	Locale   string
}

func loadConfig() *ConfigFile {
	data, err := ioutil.ReadFile("poe-tool.yml")
	if err != nil {
		panic(err)
	}

	var cfg ConfigFile
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	if cfg.Locale == "tw" || cfg.Locale == "en" {
		locale = cfg.Locale
	}

	return &cfg
}
