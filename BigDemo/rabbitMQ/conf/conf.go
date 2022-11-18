package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host   string `yaml:"Host"`
	Rabbit struct {
		Url string `yaml:"Url"`
	} `yaml:"Rabbit"`
}

func Init(path string) (*Config, error) {
	c := &Config{}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(content, c); err != nil {
		return nil, err
	}

	return c, nil
}
