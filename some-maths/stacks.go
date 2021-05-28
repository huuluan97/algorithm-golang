package some_maths

import "fmt"

/// struct of type
type Node struct {
	Value int
}

type Stack struct {
	nodes []*Node
	count int
}

type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// handle logic
func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count ++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count --
	return s.nodes[s.count]
}

func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size: size,
	}
}

func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes) + q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes) - q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}

	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}

	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return  node
}