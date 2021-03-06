package main

import (
	"io/ioutil"
	"log"

	"github.com/Ronmi/pastebin"
	yaml "gopkg.in/yaml.v2"
)

type ConfigFile struct {
	DevKey   string `yaml:"devkey,omitempty"`
	UserKey  string `yaml:"userkey,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Locale   string `yaml:"locale,omitempty"`
	GameProg string `yaml:"gameprog,omitempty"`
}

func (c *ConfigFile) Save() (err error) {
	data, err := yaml.Marshal(c)
	if err != nil {
		return
	}

	return ioutil.WriteFile("poe-tool.yml", data, 0600)
}

func loadConfig() *ConfigFile {
	data, err := ioutil.ReadFile("poe-tool.yml")
	if err != nil {
		panic(err)
	}

	var cfg ConfigFile
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.DevKey == "" && cfg.Username != "" && cfg.Password != "" {
		devkey, err := pastebin.GetDevKey(cfg.Username, cfg.Password, nil)
		if err != nil {
			log.Fatal(err)
		}

		cfg.DevKey = devkey
		cfg.Save()
	}

	if cfg.Locale == "tw" || cfg.Locale == "en" {
		locale = cfg.Locale
	}

	return &cfg
}
