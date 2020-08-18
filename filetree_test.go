package filetree

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	node, _ := Collect("testdata", nil)
	node.Visit(func(node *Node, depth int) {
		for i := 0; i < depth; i++ {
			for j := 0; j < 2; j++ {
				fmt.Print(" ")
			}
		}
		fmt.Println(node.Name)
	})
}

func TestWalk2(t *testing.T) {
	node, _ := Collect("testdata\\c", ExtensionFilter([]string{"txt"}))
	node.Visit(func(node *Node, depth int) {
		for i := 0; i < depth; i++ {
			for j := 0; j < 2; j++ {
				fmt.Print(" ")
			}
		}
		fmt.Println(node.Name)
	})
}

func TestWalk3(t *testing.T) {
	node, _ := CollectLimited("testdata", nil, 1)
	node.Visit(func(node *Node, depth int) {
		for i := 0; i < depth; i++ {
			for j := 0; j < 2; j++ {
				fmt.Print(" ")
			}
		}
		fmt.Println(node.Name)
	})
}

func TestWalk4(t *testing.T) {
	node, _ := CollectLimited("testdata", nil, 0)
	node.Visit(func(node *Node, depth int) {
		for i := 0; i < depth; i++ {
			for j := 0; j < 2; j++ {
				fmt.Print(" ")
			}
		}
		fmt.Println(node.Name)
	})
}
func TestWalk5(t *testing.T) {
	node, _ := CollectLimited("testdata", nil, -1)
	node.Visit(func(node *Node, depth int) {
		for i := 0; i < depth; i++ {
			for j := 0; j < 2; j++ {
				fmt.Print(" ")
			}
		}
		fmt.Println(node.Name)
	})
}
func TestFlatten(t *testing.T) {
	node, _ := Collect("testdata", nil)
	for _, v := range node.Flatten() {
		fmt.Println(v.Path)
	}
}
