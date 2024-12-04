package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (loc *Node) insertNode(node Node) {
	if loc.next == nil {
		loc.next = &node
	} else {
		if loc.next.data > node.data {
			node.next = loc.next
			loc.next = &node
		} else {
			loc.next.insertNode(node)
		}
	}
}

func (linkedList *LinkedList) traverseList() {
	if linkedList.head != nil {
		fmt.Println(linkedList.head.data)
		if linkedList.head.next != nil {
			linkedList.head.next.printRecursive()
		}
	}
	
}

func (node *Node) printRecursive() {
	fmt.Println(node.data)
	if node.next != nil {
		node.next.printRecursive()
	}
}


func sumDistances(leftNode *Node, rightNode *Node, sum *int) {
	var dist int
	if leftNode.data > rightNode.data {
		dist = leftNode.data - rightNode.data 
	} else {
		dist = rightNode.data - leftNode.data
	}
	
	*sum += dist
	if leftNode.next != nil {
		sumDistances(leftNode.next, rightNode.next, sum)
	}
}

func findNumberOfAppearances(data int, loc *Node, result *int)  {
	if data == loc.data {
		*result += 1	
		if loc.next != nil {
			findNumberOfAppearances(data, loc.next, result)
		}
	} else {
		if data > loc.data {
			if loc.next != nil {
				findNumberOfAppearances(data, loc.next, result)
			}
		}
	}
}

func calculateSimilarityScore(leftNode *Node, rightNode *Node, score *int) {
	appearances := 0
	findNumberOfAppearances(leftNode.data, rightNode, &appearances) 

	nodeScore := appearances * leftNode.data
	*score += nodeScore
	if leftNode.next != nil {
		calculateSimilarityScore(leftNode.next, rightNode, score)
	}
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var leftList, rightList LinkedList
	first := true
	for scanner.Scan() {
		digits := strings.Split(scanner.Text(), "   ")	
		if len(digits) == 2 {
			leftNum, err := strconv.Atoi(digits[0])
			rightNum, err := strconv.Atoi(digits[1])
			leftNode := Node{data : leftNum} 
			rightNode := Node{data : rightNum} 
			if err != nil {
				fmt.Println("failed to parse number")
			}
			if first {
				leftList.head = &leftNode
				rightList.head = &rightNode
				first = false
			} else {
				if leftList.head.data > leftNum {
					leftNode.next = leftList.head
					leftList.head = &leftNode 
				} else {
					leftList.head.insertNode(leftNode)
				}
				if rightList.head.data > rightNum {
					rightNode.next = rightList.head
					rightList.head = &rightNode 
				} else {
					rightList.head.insertNode(rightNode)
				}
				
			}
		}
	}

	sum := 0
	sumDistances(leftList.head, rightList.head, &sum)


	score := 0
	calculateSimilarityScore(leftList.head, rightList.head, &score)

	fmt.Printf("sum: %d\n", sum)
	fmt.Printf("score: %d\n", score)


}
