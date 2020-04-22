/*
@Author: Felix <https://github.com/longlivefelix>
@Date:   2020-04-21 18:42
*/
package se

import (
	"github.com/longlivefelix/tiny_search/se/structs"
)

type Search struct {
	Groups map[string]*structs.Directory
}

var searchInstance *Search = nil

func init() {
	searchInstance = &Search{
		Groups: map[string]*structs.Directory{},
	}
}

func NewDirectory(directoryName string) *structs.Directory {
	if instance, ok := searchInstance.Groups[directoryName]; ok {
		return instance
	}
	searchInstance.Groups[directoryName] = &structs.Directory{
		DirectoryName:       directoryName,
		BlockList:           []*structs.BlockInfo{},
		BlockListCountLimit: 64,               // 10m*64
		BlockSizeLimit:      1024 * 1024 * 10, // 10m
	}
	return searchInstance.Groups[directoryName]
}
