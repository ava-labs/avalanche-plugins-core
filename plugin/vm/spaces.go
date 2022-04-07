package vm

import (
	"os/exec"

	"github.com/ava-labs/avalanchego/version"

	"github.com/ava-labs/avalanche-plugin/avalanche"
)

var _ avalanche.VM = &Spaces{}

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
	cmd := exec.Command("./scripts/build.sh")
	return cmd.Run()
}

func (s Spaces) AfterInstall() error {
	return nil
}

func (s Spaces) Version() version.Version {
	return version.NewDefaultVersion(0, 0, 3)
}

func (s Spaces) URL() string {
	return "https://github.com/ava-labs/spacesvm/archive/refs/tags/v0.0.3.tar.gz"
}

func (s Spaces) SHA256() string {
	return "1ac250f6c40472f22eaf0616fc8c886078a4eaa9b2b85fbb4fb7783a1db6af3f"
}
