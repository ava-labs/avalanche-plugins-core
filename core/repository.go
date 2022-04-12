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
	avaxPlugins *avalanche.Plugins
}

func New() *Repository {
	subnets := []avalanche.Subnet{
		&subnet.Spaces{},
	}
	vms := []avalanche.VM{
		&vm.Spaces{},
	}

	plugins := &avalanche.Plugins{
		Subnets: subnets,
		VMs:     vms,
	}

	return &Repository{
		avaxPlugins: plugins,
	}
}

func (r Repository) Plugins() (avalanche.Plugins, error) {
	return *r.avaxPlugins, nil
}
