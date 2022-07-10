package common

import (
	"github.com/creasty/defaults"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Window struct {
		Title  string `default:"Type my name in config.yaml"`
		W      int    `default:"1024"`
		H      int    `default:"768"`
		Center *Pos
	}
	Graphics struct {
		Scale float32 `default:"1"`
		Debug bool    `default:"false"`
	}
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(c); err != nil {
		return errors.Wrap(err, "error setting defaults")
	}

	type plain Config
	if err := unmarshal((*plain)(c)); err != nil {
		return errors.Wrap(err, "error unmarshalling")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

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
