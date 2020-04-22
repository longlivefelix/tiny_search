/*
@Author: Felix <https://github.com/longlivefelix>
@Date:   2020-04-21 18:41
*/
package main

import (
	"fmt"
	"github.com/longlivefelix/tiny_search/se"
)
func main(){
	commentSearch := se.NewDirectory("comment")
	fmt.Println(len(commentSearch.BlockList))
	commentSearch.AddFile("1121","又是迷茫的1天哦")
	commentSearch.AddFile("1122","又是迷茫的一天哦")
	commentSearch.AddFile("1123","又是迷茫的yi天哦")
	fmt.Println(len(commentSearch.BlockList))
	commentSearch.Search("yi天")
}
