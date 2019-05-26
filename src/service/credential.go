package service

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type credential struct {
	ChannelSecret string `yaml:"channel_secret"`
	ChannelToken  string `yaml:"channel_token"`
}

func GetAppCredential() (*credential, error) {
	yamlFile, err := ioutil.ReadFile("../config/app.yml")
	if err != nil {
		return nil, err
	}

	c := credential{}

	if err := yaml.Unmarshal(yamlFile, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
