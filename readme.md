# filetree
filetree is can search file by recursive.

# install
````
git clone https://github.com/desktopgame/filetree
cd filetree
go install
````

# example
visit all files.
````
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
````

visit to .txt file
````
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
````