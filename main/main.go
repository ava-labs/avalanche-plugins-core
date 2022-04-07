package main

import (
	"github.com/hashicorp/go-plugin"

	"github.com/ava-labs/avalanche-plugin/constant"
	avaxPlugin "github.com/ava-labs/avalanche-plugin/plugin"

	"github.com/ava-labs/avalanche-plugins-core/core"
)

func main() {
	repo := core.New()

	var pluginMap = map[string]plugin.Plugin{
		constant.Repository: &avaxPlugin.RPCRepository{Impl: repo},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: core.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
