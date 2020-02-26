package filetree

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	node, _ := Collect("testdata")
	node.Visit(func(node *Node, depth int) {
		for i := 0; i < depth; i++ {
			for j := 0; j < 2; j++ {
				fmt.Print(" ")
			}
		}
		fmt.Println(node.Name)
	})
}
