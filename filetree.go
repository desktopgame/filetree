package filetree

import (
	"io/ioutil"
	"path/filepath"
	"strings"
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

// Filter is file filter.
type Filter func(node *Node) bool

// Visit is visit all element for recursive in node
func (node *Node) Visit(v Visitor) {
	node._visit(v, 0)
}

// Flatten is node tree convert to array
func (node *Node) Flatten() []*Node {
	return node._flatten(nil)
}

func (node *Node) _visit(v Visitor, depth int) {
	v(node, depth)
	for _, child := range node.Children {
		child._visit(v, depth+1)
	}
}

func (node *Node) _flatten(out []*Node) []*Node {
	out = append(out, node)
	for _, v := range node.Children {
		out = v._flatten(out)
	}
	return out
}

// CollectLimited is create file tree node.
func CollectLimited(dir string, filter Filter, limit int) (*Node, error) {
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
		if limit == 0 {
			break
		}
		filename := file.Name()
		if file.IsDir() {
			child, err := CollectLimited(filepath.Join(dir, filename), filter, limit-1)
			if err != nil {
				return nil, err
			}
			if filter == nil || filter(child) {
				node.Children = append(node.Children, child)
			}
		} else {
			child := &Node{
				Path:     filepath.Join(dir, filename),
				Name:     filename,
				IsDir:    file.IsDir(),
				Children: nil,
			}
			if filter == nil || filter(child) {
				node.Children = append(node.Children, child)
			}
		}
	}
	return node, nil
}

// Collect is create file tree node.
func Collect(dir string, filter Filter) (*Node, error) {
	return CollectLimited(dir, filter, -1)
}

// ExtensionFilter is returns new filter by file extension.
func ExtensionFilter(extensions []string) Filter {
	return func(node *Node) bool {
		if node.IsDir {
			return true
		}
		for _, ext := range extensions {
			if ext[0] != '.' {
				ext = "." + ext
			}
			if strings.HasSuffix(node.Name, ext) {
				return true
			}
		}
		return false
	}
}
