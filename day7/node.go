package main

type Node struct {
	Name   string
	Parent *Node
	IsDir  bool
	Bytes  int
	Sub    []*Node
}

func (n *Node) Size() int {
	if !n.IsDir {
		return n.Bytes
	}
	if n.Bytes > 0 {
		return n.Bytes
	}
	for _, node := range n.Sub {
		n.Bytes += node.Size()
	}
	return n.Bytes
}

func (n *Node) Mkdir(name string) {
	// fmt.Println("Making DIR", name, "inside", n.Name)
	n.Sub = append(n.Sub, &Node{Name: name, Parent: n, IsDir: true})
}

func (n *Node) Touch(name string, size int) {
	// fmt.Println("Adding FILE", name, "inside", n.Name, "size", size)
	n.Sub = append(n.Sub, &Node{Name: name, Bytes: size})
}

func (n *Node) Search(max int) []*Node {
	res := []*Node{}
	if n.IsDir {
		if n.Size() < max {
			res = append(res, n)
		}
		for _, child := range n.Sub {
			res = append(res, child.Search(max)...)
		}
	}
	return res
}

func (n *Node) Dirsizes() []int {
	res := []int{}
	if n.IsDir {
		res = append(res, n.Size())
		for _, child := range n.Sub {
			res = append(res, child.Dirsizes()...)
		}
	}
	return res
}
