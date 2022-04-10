package config

import (
	"io/ioutil"

	"github.com/pelletier/go-toml/v2"
)

func Read(path string) (MatrixConfig, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return MatrixConfig{}, err
	}

	var cfg MatrixConfig
	err = toml.Unmarshal(content, &cfg)
	return cfg, err
}
