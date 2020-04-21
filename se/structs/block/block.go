/*
@Author: Felix <https://github.com/longlivefelix>
@Date:   2020-04-21 18:54
*/
package block

import "github.com/longlivefelix/tiny_search/se/structs/index"

type BlockInfo struct{
	BlockID int64
	LargestPrimaryId int64

	BlockCharSet map[rune]index.RuneIndex

	Files []int64
}