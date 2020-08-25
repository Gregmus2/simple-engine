package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Window struct {
		Title  string
		W      int
		H      int
		Center *Pos
	}
	Physics struct {
		Enable  bool
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
	cfg := &Config{
		Window: struct {
			Title  string
			W      int
			H      int
			Center *Pos
		}{
			Title: "Type my name in config.yaml",
			W:     1024,
			H:     768,
		},
		Physics: struct {
			Enable  bool
			Scale   float64
			Gravity struct {
				X float64
				Y float64
			}
		}{
			Scale: 100,
			Gravity: struct {
				X float64
				Y float64
			}{
				X: 0,
				Y: -5,
			},
		},
		Graphics: struct {
			Scale float32
		}{Scale: 100},
	}

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if !os.IsNotExist(err) {
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(yamlFile, cfg)
		if err != nil {
			return nil, err
		}
	}

	cfg.Window.Center = &Pos{
		X: float32(cfg.Window.W / 2),
		Y: float32(cfg.Window.H / 2),
	}

	return cfg, nil
}
