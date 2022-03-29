package vm

import (
	"github.com/ava-labs/apm/definition"
	"github.com/ava-labs/avalanchego/version"
)

var _ definition.VM = &Spaces{}

type Spaces struct {
}

func (s Spaces) ID() string {
	return "tGas3T58KzdjLHhBDMnH2TvrddhqTji5iZAMZ3RXs2NLpSnhH"
}

func (s Spaces) Alias() string {
	return "spacesvm"
}

func (s Spaces) Homepage() string {
	return "https://tryspaces.xyz"
}

func (s Spaces) Description() string {
	return "Spaces enables authenticated, hierarchical storage of arbitrary " +
		"keys/values using any EIP-712 compatible wallet."
}

func (s Spaces) Maintainers() []string {
	return []string{"patrickogrady@avalabs.org"}
}

func (s Spaces) BeforeInstall() error {
	return nil
}

func (s Spaces) Install() error {
	return nil
}

func (s Spaces) AfterInstall() error {
	return nil
}

func (s Spaces) Version() version.Version {
	return version.NewDefaultVersion(0, 0, 3)
}

func (s Spaces) Repository() string {
	return "https://github.com/ava-labs/spacesvm.git"
}

func (s Spaces) SHA256() string {
	panic("implement me")
}

func (s Spaces) Commit() string {
	return "10f885ffc2770951bb0fd4d2b4b658c3b9cc5caa"
}

func (s Spaces) Branch() string {
	return "master"
}
