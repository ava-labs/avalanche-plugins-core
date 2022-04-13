package core

import (
	"github.com/ava-labs/avalanche-plugin/avalanche"
	"github.com/ava-labs/avalanche-plugin/plugin"
	"github.com/ava-labs/avalanche-plugins-core/plugin/subnet"
	"github.com/ava-labs/avalanche-plugins-core/plugin/vm"
)

const (
	Alias = "ava-labs/avalanche-plugins-core"
	URL   = "https://github.com/ava-labs/avalanche-plugins-core.git"
)

var _ plugin.Repository = &Repository{}

type Repository struct {
	vms     []avalanche.VM
	subnets []avalanche.Subnet
}

func New() *Repository {
	subnets := []avalanche.Subnet{
		&subnet.Spaces{},
	}
	vms := []avalanche.VM{
		&vm.Spaces{},
	}

	return &Repository{
		vms:     vms,
		subnets: subnets,
	}
}

func (r *Repository) GetSubnets() ([]avalanche.Subnet, error) {
	return r.subnets, nil
}

func (r *Repository) GetVMs() ([]avalanche.VM, error) {
	return r.vms, nil
}
