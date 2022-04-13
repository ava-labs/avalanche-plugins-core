package main

import (
	"github.com/hashicorp/go-plugin"

	"github.com/ava-labs/avalanche-plugin/constant"
	"github.com/ava-labs/avalanche-plugin/grpc"
	"github.com/ava-labs/avalanche-plugins-core/core"
)

func main() {
	repo := core.New()

	var pluginMap = map[string]plugin.Plugin{
		constant.Repository: &grpc.PluginRepository{PluginRepository: repo},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: core.HandshakeConfig,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}
