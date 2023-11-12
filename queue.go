//Поиск правого соседа BST (вариант 2)
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
	node := q.nodes[0] //извлечение первого элемента (голову) из очереди, так как в очереди первый пришёл, первый вышел. Он сохраняется в переменной node
	q.nodes = q.nodes[1:] //обновление среза q.nodes, удаляя первый элемент (голову), т о, уменьшая размер очереди на 1. Теперь в q.nodes остались остальные элементы, и первый элемент удалён из очереди
	return node
}

//Функция для поиска правых соседей и вывода результатов
func FindRightNeighbors(root *Node) {
	if root == nil {
		return
	}
	
	//обход дерева в ширину
	queue := Queue{} //инциализация очереди
	queue.Enqueue(root) //добавление первого уровня дерева в очередь
	queue.Enqueue(nil) //для разделения уровней

	//обход уровней дерева
	for {
		node := queue.Dequeue() //извлечение узла из очереди
		if node == nil { //проверка, является ли извлечённый узел nil
			if len(queue.nodes) == 0 {
				break
			}
			queue.Enqueue(nil) // добавление nil для нового уровня 
			continue
	} 
	//определение правый соседа текущего узла Правый сосед - это первый элемент в очереди, т к(обход идёт слева направо по уровню)
	rightNeighbor := queue.nodes[0]

	//Вывод правого соседа или "null" если соседа нет
	if rightNeighbor != nil {
		fmt.Printf("%d - %d, ", node.val, rightNeighbor.val)
	} else {
		fmt.Printf("%d - null, ", node.val)
	}

	//Добавление детоузла в очередь
	if node.left != nil { //если у текущего узла есть левый дочерний узел, то он добавляется в конец очереди
		queue.Enqueue(node.left)
	}
	if node.right != nil {
		queue.Enqueue(node.right) //добавление в конец очереди
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