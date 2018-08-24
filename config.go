package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// config - config
type config struct {
	BindAddr     string
	ServAddr     string
	RootServAddr string
	NodeName     string
	runpath      string
}

// loadConfig - load config
func loadConfig(filename string) (*config, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err1 := ioutil.ReadAll(fi)
	if err1 != nil {
		return nil, err1
	}

	cfg := config{}

	err2 := yaml.Unmarshal(fd, &cfg)
	if err2 != nil {
		return nil, err2
	}

	return &cfg, nil
}
