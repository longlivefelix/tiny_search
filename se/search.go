/*
@Author: Felix <https://github.com/longlivefelix>
@Date:   2020-04-21 18:42
*/
package se

import (
	"github.com/longlivefelix/tiny_search/se/structs/block"
	"github.com/longlivefelix/tiny_search/se/structs/group"
)

type Search struct {
	Instance *group.Group
}

var searchInstance map[string]*Search

func NewGroup(groupName string) *Search {
	if instance, ok := searchInstance[groupName]; ok {
		return instance
	}
	searchInstance[groupName] = &Search{
		Instance: &group.Group{
			GroupName:           groupName,
			BlockList:           []block.BlockInfo{},
			BlockListCountLimit: 64,               // 10m*64
			BlockSizeLimit:      1024 * 1024 * 10, // 10m
		},
	}
	return searchInstance[groupName]
}
