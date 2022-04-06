package main

import (
	"fmt"

	"github.com/hashicorp/go-plugin"

	avaxPlugin "github.com/ava-labs/avalanche-plugin/plugin"

	"github.com/ava-labs/avalanche-plugins-core/core"
)

func main() {
	// pluginMap is the map of plugins we can dispense.
	repo := core.New()

	var pluginMap = map[string]plugin.Plugin{
		"repository": &avaxPlugin.RPCRepository{Impl: repo},
	}

	fmt.Println("Core loaded.")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: core.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
