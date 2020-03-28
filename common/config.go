package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Window struct {
		Title  string
		H      int
		W      int
		Center *Pos
	}
	Physics struct {
		Scale   float64
		Gravity struct {
			X float64
			Y float64
		}
	}
	Graphics struct {
		Scale float32
	}
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, err
	}

	cfg.Window.Center = &Pos{
		X: float32(cfg.Window.W / 2),
		Y: float32(cfg.Window.H / 2),
	}

	return cfg, nil
}
