package addons

import (
	"github.com/stader-labs/stader-node/addons/graffiti_wall_writer"
	"github.com/stader-labs/stader-node/shared/types/addons"
)

func NewGraffitiWallWriter() addons.SmartnodeAddon {
	return graffiti_wall_writer.NewGraffitiWallWriter()
}
