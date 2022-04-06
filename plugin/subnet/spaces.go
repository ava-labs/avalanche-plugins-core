package subnet

import (
	"github.com/ava-labs/avalanchego/chains"
	consensus "github.com/ava-labs/avalanchego/snow/consensus/avalanche"
	"github.com/ava-labs/avalanchego/snow/consensus/snowball"
	"github.com/ava-labs/avalanchego/snow/networking/sender"

	"github.com/ava-labs/avalanche-plugin/avalanche"
	"github.com/ava-labs/avalanche-plugins-core/plugin/vm"
)

var _ avalanche.Subnet = &Spaces{}

type Spaces struct {
}

func (s Spaces) ID() string {
	return "Ai42MkKqk8yjXFCpoHXw7rdTWSHiKEMqh5h8gbxwjgkCUfkrk"
}

func (s Spaces) Alias() string {
	return "spaces"
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

func (s Spaces) VMs() []avalanche.VM {
	return []avalanche.VM{
		&vm.Spaces{},
	}
}

func (s Spaces) SubnetConfig() *chains.SubnetConfig {
	return &chains.SubnetConfig{
		GossipConfig: sender.GossipConfig{
			AcceptedFrontierSize:      10,
			OnAcceptSize:              10,
			AppGossipNonValidatorSize: 10,
			AppGossipValidatorSize:    10,
		},
		ValidatorOnly: false,
		ConsensusParameters: consensus.Parameters{
			Parameters: snowball.Parameters{
				K:                     10,
				Alpha:                 10,
				BetaVirtuous:          10,
				BetaRogue:             10,
				ConcurrentRepolls:     10,
				OptimalProcessing:     10,
				MaxOutstandingItems:   10,
				MaxItemProcessingTime: 10,
			},
			Parents:   10,
			BatchSize: 10,
		},
	}
}
