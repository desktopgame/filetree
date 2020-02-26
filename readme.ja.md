# filetree
filetree は再帰的にファイルを検索できます。

# install
````
git clone https://github.com/desktopgame/filetree
cd filetree
go install
````

# example
全てのファイルにアクセスします。
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

拡張子が .txt のファイルにアクセスします。
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