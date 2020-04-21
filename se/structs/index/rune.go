/*
@Author: Felix <https://github.com/longlivefelix>
@Date:   2020-04-21 18:19
*/
package index

type Offset uint16
type OffsetList struct {
	Count   int8
	Offsets []Offset
}
type FileIdx uint16                   // File serial number in block
type RuneField map[FileIdx]OffsetList // Rune’s offsets
type RuneIndex map[rune]RuneField  // The index of all runes' offsets in every file

