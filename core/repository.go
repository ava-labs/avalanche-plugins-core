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
	AvaxPlugins avalanche.Plugins
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
		AvaxPlugins: plugins,
	}
}

func (r Repository) Plugins() (avalanche.Plugins, error) {
	return r.AvaxPlugins, nil
}
