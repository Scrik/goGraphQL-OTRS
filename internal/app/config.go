package app

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// TypeConfig Config
type TypeConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	ClientURL   string `yaml:"client_url"`
	GraphiQLURL string `yaml:"graphiql_url"`
	Schema      string `yaml:"schema"`
}

// NewConfig NewConfig
func NewConfig(data []byte) (*TypeConfig, error) {
	config := &TypeConfig{}
	err := yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, err
}

// DefaultConfig Default
func DefaultConfig() TypeConfig {
	C := TypeConfig{}
	C.Host = "localhost"
	C.Port = 3000
	C.GraphiQLURL = "/graphql"
	C.ClientURL = "/cli"
	C.Schema = "schema.gql"
	return C
}

// GetURL GetURL
func (C *TypeConfig) GetURL() string {
	return fmt.Sprintf("%s:%d", C.Host, C.Port)
}

// Generate Generate
func Generate(filename string) {
	bytes, err := yaml.Marshal(DefaultConfig())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
	return
}
