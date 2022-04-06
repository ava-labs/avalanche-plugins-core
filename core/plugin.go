package core

import "github.com/hashicorp/go-plugin"

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "CORE_PLUGINS",
	MagicCookieValue: "dynamic",
}
