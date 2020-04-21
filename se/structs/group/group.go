
package group

import (
	"github.com/longlivefelix/tiny_search/se/structs/block"
)
type Group struct {
	GroupName string
	BlockList []block.BlockInfo
	BlockSizeLimit  int64
	BlockListCountLimit int64
}