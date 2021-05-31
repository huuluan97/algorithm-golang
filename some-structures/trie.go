package some_structures

type NodeTrie struct {
	children map[rune]*NodeTrie
	isLeaf bool
}

func NewNode() *NodeTrie {
	n := &NodeTrie{}
	n.children = make(map[rune]*NodeTrie)
	n.isLeaf = false
	return n
}

func (n *NodeTrie) Insert(s string) {
	current := n
	for _, c := range s {
		next, ok := current.children[c]
		if !ok {
			next = NewNode()
			current.children[c] = next
		}
		current = next
	}
	current.isLeaf = true
}

func (n *NodeTrie) Find(s string) bool {
	current := n
	for _, c := range s {
		next, ok := current.children[c]
		if !ok {
			return false
		}
		current = next
	}
	return true
}