package main

import (
	"fakeeyes_client/client"
	"fakeeyes_client/config"

	"github.com/goodaye/wire"
)

// fakeeyes client
var fesclient *client.Client

func init() {
	wire.Append(SVC{})
}

type SVC struct {
	wire.BaseService
}

func (s SVC) Init() error {
	var err error
	fesclient, err = client.NewClient(config.GlobalConfig.Fakeeyes.Server)
	if err != nil {
		return err
	}
	return nil
}

func (s SVC) Start() error {
	return nil
}
