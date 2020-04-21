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
	search := se.NewGroup("comment")
	fmt.Println(search.Instance)
}
