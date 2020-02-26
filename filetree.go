package filetree

import (
	"io/ioutil"
	"path/filepath"
)

// Node is element of file tree
type Node struct {
	Path     string
	Name     string
	IsDir    bool
	Children []*Node
}

// Visitor is callback function for node recursive.
type Visitor func(node *Node, depth int)

// Visit is visit all element for recursive in node
func (node *Node) Visit(v Visitor) {
	node._visit(v, 0)
}

func (node *Node) _visit(v Visitor, depth int) {
	v(node, depth)
	for _, child := range node.Children {
		child._visit(v, depth+1)
	}
}

// Collect is create file tree node.
func Collect(dir string) (*Node, error) {
	node := &Node{
		Path:     dir,
		Name:     dir,
		IsDir:    true,
		Children: nil,
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		filename := file.Name()
		if file.IsDir() {
			child, err := Collect(filepath.Join(dir, filename))
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, child)
		} else {
			node.Children = append(node.Children, &Node{
				Path:     filepath.Join(dir, filename),
				Name:     filename,
				IsDir:    file.IsDir(),
				Children: nil,
			})
		}
	}
	return node, nil
}
