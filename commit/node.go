package commit

type Node struct {
	Parent   *Node
	Children []*Node
	Hash     []byte
	Change   Change
}
