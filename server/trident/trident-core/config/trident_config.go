package config

import (
	"errors"
	nc "nvm-gocore"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type TridentConfig struct {
	Version      string             `yaml:"version"`
	Bind         string             `yaml:"bind"`
	Port         int                `yaml:"port"`
	IsUseExtAuth bool               `yaml:"use-ext-auth"`
	Services     map[string]Service `yaml:"services"`
}

type Service struct {
	ServiceName string   `yaml:"service-name"`
	BaseURL     string   `yaml:"base-url"`
	Scopes      []string `yaml:"scopes"`
	Allow       []string `yaml:"allow"`
	Disallow    []string `yaml:"disallow"`
}

type TridentConfigParser struct {
}

func NewTridentConfig() *TridentConfigParser {
	tc := &TridentConfigParser{}
	return tc
}

func (tc *TridentConfigParser) GetBinaryDir() nc.Option[string] {
	// Get the path to the executable.
	exePath, err := os.Executable()
	if err != nil {
		return nc.None[string]()
	}

	// Determine the directory of the executable.
	exeDir := filepath.Dir(exePath)

	// Ensure the directory ends with a separator.
	if !strings.HasSuffix(exeDir, string(os.PathSeparator)) {
		exeDir += string(os.PathSeparator)
	}

	return nc.Some[string](exeDir)
}

func (tc *TridentConfigParser) OpenConfig(path string) nc.Result[*[]byte] {

	content, err := os.ReadFile(path)
	if err != nil {
		return nc.NewResult[*[]byte](nil, err)
	}

	return nc.NewResult[*[]byte](&content, nil)
}

func (tc *TridentConfigParser) ParseConfig(content *[]byte) nc.Result[*TridentConfig] {
	var config TridentConfig
	err := yaml.Unmarshal(*content, &config)
	if err != nil {
		return nc.NewResult[*TridentConfig](nil, errors.New("error unmarshalling YAML:"+err.Error()))
	}

	return nc.NewResult[*TridentConfig](&config, nil)
}
