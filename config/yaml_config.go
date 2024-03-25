package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg6/go-defaults"
	"github.com/pkg6/httpgo/t"
	"gopkg.in/yaml.v3"
	"os"
)

type YAMLConfig struct {
	mapdata map[string]any
	config  any
}

func (this *YAMLConfig) SetConfig(config any) {
	this.config = config
}

func (this *YAMLConfig) ReadFile(file string) error {
	var (
		bytes []byte
		err   error
	)
	bytes, err = os.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, &this.mapdata)
	if err != nil {
		return err
	}
	return nil
}

func (this *YAMLConfig) Decode() error {
	return mapstructure.Decode(this.mapdata, this.config)
}

func (this *YAMLConfig) Defaults() error {
	return defaults.Set(this.config)
}

func (this *YAMLConfig) Validate() error {
	return t.ValidateStruct(this.config)
}
