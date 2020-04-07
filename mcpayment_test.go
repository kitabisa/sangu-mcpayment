package mcpayment

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Configs struct {
	BaseURLToken    string
	XSignKey        string
	ReturnURL       string
	RegisteredID    string
	RegisteredToken string
}

// GetConfig get config for test
func GetConfig() (Configs, error) {
	theToml, err := ioutil.ReadFile("credential_test.toml")
	if err != nil {
		return Configs{}, err
	}

	var conf Configs
	if _, err := toml.Decode(string(theToml), &conf); err != nil {
		return Configs{}, err
	}

	return conf, nil
}
