/*
@Author: Felix <https://github.com/longlivefelix>
@Date:   2020-04-21 18:42
*/
package se

import (
	"github.com/longlivefelix/tiny_search/se/structs/group"
)

type Search struct {
	Instance *group.Group
}

func NewGroup(groupName string) *Search{
	return &Search{}
}