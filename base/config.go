package base

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	OpenRange []Range `yaml:openrange`
}

type Range struct {
	Name   string         `yaml:name`
	Range  map[string]int `yaml:range`
	Regexp string         `yaml:regexp`
}

func ReadConfig(path string) (*Config, error) {
	conf := Config{}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if _, err = os.Create(path); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return &conf, err
	}
	return &conf, nil
}

func WriteConf(conf *Config, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if _, err = os.Create(path); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
}
