package structs

type Directory struct {
	DirectoryName       string
	BlockList           []*BlockInfo
	BlockSizeLimit      int32
	BlockListCountLimit int32
	Files               []*File
}

func (g *Directory) addBlock() *BlockInfo {
	if int32(len(g.BlockList)) >= g.BlockListCountLimit {
		return nil
	}
	blockId := int32(len(g.BlockList))
	tmp := &BlockInfo{
		BlockID:      blockId,
		ContentIndex: RuneIndex{},
		MemSize:      0,
	}
	g.BlockList = append(g.BlockList, tmp)
	return tmp
}

func (g *Directory) AddFile(primaryId string, content string) error {
	var currentBlock *BlockInfo = nil
	for _, block := range g.BlockList {
		if block.MemSize >= g.BlockSizeLimit {
			continue
		}
		currentBlock = block
		break
	}
	if currentBlock == nil {
		currentBlock = g.addBlock()
	}
	if currentBlock == nil {
		return SearchErr{level: "FAIL", msg: "create block failed!"}
	}
	g.Files = append(g.Files, &File{
		PrimaryId:primaryId,
		Content:content,
	})
	fileIdx := FileIdx(len(g.Files)-1)
	return currentBlock.AddFile(fileIdx, content)
}

func (g *Directory) Search(content string) []*File{
	result := []*File{}
	for _, block := range g.BlockList{
		for _,fileIdx := range block.Search(content){
			if fileIdx>=0 && int(fileIdx)<len(g.Files) && g.Files[int(fileIdx)]!= nil{
				result = append(result,g.Files[int(fileIdx)])
			}
		}
	}
	return result
}