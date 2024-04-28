package mcdata

import "embed"

var (
	SubmodulePath = "minecraft-data"
	//go:embed minecraft-data
	embeddedMinecraftData embed.FS
)

type MCData struct {
	ProtocolVersion string // bedrock | pc
	Version         string
}
