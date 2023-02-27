package common

import (
	"github.com/creasty/defaults"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Config *ConfigModel

func DefineConfig() error {
	Config = &ConfigModel{}

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if !os.IsNotExist(err) {
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(yamlFile, Config)
		if err != nil {
			return err
		}
	}

	Config.Window.Center = &Pos{
		X: float32(Config.Window.W / 2),
		Y: float32(Config.Window.H / 2),
	}

	return nil
}

type ConfigModel struct {
	Window struct {
		Title  string `default:"Type my name in config.yaml"`
		W      int    `default:"1024"`
		H      int    `default:"768"`
		Center *Pos
	}
	Graphics struct {
		FPS      int     `default:"60"`
		Scale    float32 `default:"1"`
		Debug    bool    `default:"false"`
		Font     map[string]string
		Textures string
	}
	Debug struct {
		FPS bool `default:"false"`
	}
}

func (c *ConfigModel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(c); err != nil {
		return errors.Wrap(err, "error setting defaults")
	}

	type plain ConfigModel
	if err := unmarshal((*plain)(c)); err != nil {
		return errors.Wrap(err, "error unmarshalling")
	}

	return nil
}
