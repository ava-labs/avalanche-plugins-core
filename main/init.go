package main

import (
	"encoding/gob"

	"github.com/ava-labs/avalanche-plugins-core/plugin/subnet"
	"github.com/ava-labs/avalanche-plugins-core/plugin/vm"
)

func init() {
	gob.Register(&subnet.Spaces{})
	gob.Register(subnet.Spaces{})
	gob.Register(&vm.Spaces{})
	gob.Register(vm.Spaces{})
}
