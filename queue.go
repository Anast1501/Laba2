/*
package main 
 import (
	"fmt"
	"errors"
 )

 type QNode struct {
	next *QNode
	val *Node
 }

 type Queue struct {
	front *QNode
	rear *QNode
 }

func (q *Queue) Qpush(value *Node) {
	newNode := &QNode {
		val: value,
		next: nil,
	}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

func (q *Queue) Qpop() (*Node, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	node := q.front.val
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}
	return node, nil
}

func (q *Queue) Qfront() (*Node, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	return q.front.val, nil
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

type Node struct {
	val int
	left *Node
	right *Node
}

func FindRightNeighbors(root *Node) {
	if root == nil {
		return 
	}

	queue := &Queue{}
	queue.Qpush(root)

	for !queue.IsEmpty() {
		node, _ :=queue.Qpop()

		rightNeighbor, _ :=queue.Qfront()

		fmt.Println(node.val, " - ")
		//fmt.Printf("%d - ", node.val)

		if rightNeighbor!=nil {
			fmt.Print(rightNeighbor.val, ", ")
		} else {
			fmt.Print("null, ")
		}
		if node.left != nil {
			queue.Qpush(node.left)
	}
	    if node.right != nil {
			queue.Qpush(node.right)
}
}
fmt.Println()
}

func main() {
	root := &Node{val: 22}
	root.left = &Node{val: 16}
	root.left.left = &Node{val: 7}
	root.left.right = &Node{val: 19}
	root.right = &Node{val: 51}
	root.right.right = &Node{val: 57}
	root.right.left = &Node{val: 43}

	FindRightNeighbors(root)
}
*/


/* !!!!!!!!!!
package main

import (
	"fmt"
	"errors"
)

type QNode struct {
	next *QNode
	val  *Node
}

type Queue struct {
	front *QNode
	rear  *QNode
}

func (q *Queue) Qpush(value *Node) {
	newNode := &QNode{
		val:  value,
		next: nil,
	}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

func (q *Queue) Qpop() (*Node, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	node := q.front.val
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}
	return node, nil
}

func (q *Queue) Qfront() (*Node, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	return q.front.val, nil
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

type Node struct {
	val   int
	left  *Node
	right *Node
}


func FindRightNeighbors(root *Node) {
	if root == nil {
		return
	}

	queue := &Queue{}
	queue.Qpush(root)

	for !queue.IsEmpty() {
		node, _ := queue.Qpop()

		rightNeighbor, _ := queue.Qfront()

		fmt.Print(node.val, " - ")
		
		if rightNeighbor != nil {
			fmt.Print(rightNeighbor.val, ", ")
		} else {
			fmt.Print("null, ")
		}

		if node.left != nil {
			queue.Qpush(node.left)
		}
		if node.right != nil {
			queue.Qpush(node.right)
		}
		
	}
	fmt.Println() // Add a new line after the traversal
}



func main() {
	root := &Node{val: 22}
	root.left = &Node{val: 16}
	root.right = &Node{val: 51}
	root.left.left = &Node{val: 7}
	root.left.right = &Node{val: 19}
	root.right.left = &Node{val: 43} 
	root.right.right = &Node{val: 57}
	

	FindRightNeighbors(root)
}
*/




/* !!!!!!!!!!!
package main

import (
 "errors"
 "fmt"
)

type QNode struct {
 next *QNode
 val  *Node
}

type Queue struct {
 front *QNode
 rear  *QNode
}

func (q *Queue) Qpush(value *Node) {
 newNode := &QNode{
  val:  value,
  next: nil,
 }

 if q.rear == nil {
  q.front = newNode
  q.rear = newNode
 } else {
  q.rear.next = newNode
  q.rear = newNode
 }
}

func (q *Queue) Qpop() (*Node, error) {
 if q.IsEmpty() {
  return nil, errors.New("Queue is empty")
 }

 node := q.front.val
 q.front = q.front.next

 if q.front == nil {
  q.rear = nil
 }
 return node, nil
}

func (q *Queue) Qfront() (*Node, error) {
 if q.IsEmpty() {
  return nil, errors.New("Queue is empty")
 }
 return q.front.val, nil
}

func (q *Queue) IsEmpty() bool {
 return q.front == nil
}

type Node struct {
 val   int
 left  *Node
 right *Node
}

func FindRightNeighbors(root *Node) {
 if root == nil {
  return
 }

 queue := &Queue{}
 queue.Qpush(root)

 for !queue.IsEmpty() {
  currentNode, _ := queue.Qpop()

  nextNode, _ := queue.Qfront()

  fmt.Printf("%d - ", currentNode.val)

  if nextNode != nil {
   fmt.Printf("%d, ", nextNode.val)
  } else {
   fmt.Printf("null, ")
  }

  if currentNode.left != nil {
   queue.Qpush(currentNode.left)
  }
  if currentNode.right != nil {
   queue.Qpush(currentNode.right)
  }

 }
 fmt.Println() // Add a new line after the traversal
}

func main() {
 root := &Node{val: 22}
 root.left = &Node{val: 16}
 root.right = &Node{val: 51}
 root.left.left = &Node{val: 7}
 root.left.right = &Node{val: 19}
 root.right.left = &Node{val: 43}
 root.right.right = &Node{val: 57}

 FindRightNeighbors(root)
}
*/


package main

import (
	"fmt"
)

//Определение структуры узла дерева
type Node struct {
	val   int
	left  *Node
	right *Node
}

//Определение структуры очереди для узлов дерева
type Queue struct {
	nodes []*Node
}

//Метод для добавления узла в очередь
func (q *Queue) Enqueue(node *Node) {
	q.nodes = append(q.nodes, node)
}

//Метод для удаления узла из очереди
func (q *Queue) Dequeue() *Node {
	if len(q.nodes) == 0 {
		return nil
	}
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

//Функция для поиска правых соседей и вывода результатов
func FindRightNeighbors(root *Node) {
	if root == nil {
		return
	}

	queue := Queue{}
	queue.Enqueue(root)
	queue.Enqueue(nil) //Маркер для разделения уровней

	for {
		node := queue.Dequeue()
		if node == nil {
			if len(queue.nodes) == 0 {
				break
			}
			queue.Enqueue(nil) // Маркер для нового уровня 
			continue
	}

	rightNeighbor := queue.nodes[0]

	//Вывод правого соседа или "null" если соседа нет
	if rightNeighbor != nil {
		fmt.Printf("%d - %d, ", node.val, rightNeighbor.val)
	} else {
		fmt.Printf("%d - null, ", node.val)
	}

	//Добавление детоузла в очередь
	if node.left != nil {
		queue.Enqueue(node.left)
	}
	if node.right != nil {
		queue.Enqueue(node.right)
	}
	}

	fmt.Println()
}

func main() {
	//Создание бинарного дерева
	root := &Node{val: 22}
	root.left = &Node{val: 16}
	root.right = &Node{val: 51}
	root.left.left = &Node{val: 7}
	root.left.right = &Node{val: 19}
	root.right.left = &Node{val: 43}
	root.right.right = &Node{val: 57}

	//Вызов функции поиска правых соседей и вывод результатов
	FindRightNeighbors(root)
}




























/*
func main() {
	root := &BSTNode{value: 22}
	root.left = &BSTNode{value: 16}
	root.right = &BSTNode{value: 51}
	root.left.left = &BSTNode{value: 7}
	root.right.left = &BSTNode{value: 19}
	root.right.right = &BSTNode{value: 43}
	root.right.right.right = &BSTNode{value: 57}

	fmt.Println("Right Neighbors:")
	printRightNeighbors(root)
}
*/
