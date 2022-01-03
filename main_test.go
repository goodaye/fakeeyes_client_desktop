package main

import (
	"fakeeyes_client/config"
	"fmt"
	"testing"

	"github.com/goodaye/wire"
)

func TestRegister(t *testing.T) {

	configFilePath := "./bin/config.toml"
	config.SetConfigFile(configFilePath)
	err := wire.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	Register()

}
