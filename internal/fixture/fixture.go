package fixture

import "embed"

//go:embed fixture.yml
var EmbedFS embed.FS
