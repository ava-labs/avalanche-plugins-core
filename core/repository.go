package core

import (
	"github.com/ava-labs/avalanche-plugin/avalanche"
	"github.com/ava-labs/avalanche-plugin/plugin"

	"github.com/ava-labs/avalanche-plugins-core/plugin/subnet"
	"github.com/ava-labs/avalanche-plugins-core/plugin/vm"
)

const (
	Name = "avalanche-plugins-core"
	URL  = "https://github.com/ava-labs/avalanche-plugins-core.git"
)

var _ plugin.Repository = &Repository{}

type Repository struct {
	plugins avalanche.Plugins
}

func New() *Repository {
	subnets := []avalanche.Subnet{
		&subnet.Spaces{},
	}

	vms := []avalanche.VM{
		&vm.Spaces{},
	}

	plugins := avalanche.Plugins{
		Subnets: subnets,
		VMs:     vms,
	}

	return &Repository{
		plugins: plugins,
	}
}

func (r Repository) Plugins() (avalanche.Plugins, error) {
	return r.plugins, nil
}
