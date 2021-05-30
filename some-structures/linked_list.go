package some_structures

import "fmt"

type Node struct {
	nextNode *Node
	property rune
}

func CreateLinkedList() *Node {
	var headNode *Node
	headNode = &Node{nil, 'a'}

	var currNode *Node
	currNode = headNode

	var i rune
	for i = 'b'; i <= 'z'; i++ {
		var node *Node
		node = &Node{nil, i}
		currNode.nextNode = node
		currNode = node
	}

	return headNode
}

func StringifyList(nodeList *Node) {
	var currentNode *Node
	currentNode = nodeList

	for {
		fmt.Printf("%c", currentNode.property)
		if currentNode.nextNode != nil {
			currentNode = currentNode.nextNode
		} else {
			break
		}
	}

	fmt.Println("")
}

func ReverseLinkedList(nodeList *Node) *Node {

	var currentNode *Node
	currentNode = nodeList
	var topNode *Node = nil
	for {
		if currentNode == nil {
			break
		}
		var tempNode *Node
		tempNode = currentNode.nextNode
		currentNode.nextNode = topNode
		topNode = currentNode
		currentNode = tempNode
	}

	return topNode
}

func handleRunLinkedListTest() {

	var linkedList = CreateLinkedList()
	StringifyList(linkedList)
	StringifyList(ReverseLinkedList(linkedList))
}