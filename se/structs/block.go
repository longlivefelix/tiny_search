/*
@Author: Felix <https://github.com/longlivefelix>
@Date:   2020-04-21 18:54
*/
package structs

type BlockInfo struct {
	BlockID      int32
	ContentIndex RuneIndex
	MemSize      int32
}

func (b BlockInfo) AddFile(fileIdx FileIdx, content string) error {
	contentBytes := []byte(content)
	b.ContentIndex.BuildIndex(contentBytes,fileIdx)
	b.MemSize+=int32(len(contentBytes))
	return nil
}

func (b BlockInfo) Search(content string) []FileIdx{
	return b.ContentIndex.Search(content)
}