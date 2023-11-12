//К-тая максимальная сумма (вариант 5)

//Задано двоичное дерево и целочисленное число k. Необходимо найти сумму
//каждого уровня дерева, и вывести k-тое из них.

package main

import (
    "fmt"
)

//BSTNode представляет узел двоичного дерева поиска
type BSTNode struct {
    Key   int
    Left  *BSTNode
    Right *BSTNode
}

//NewNode создаёт новый узел с заданным ключом
func NewNode(key int) *BSTNode {
    return &BSTNode{Key: key, Left: nil, Right: nil}
}

//BSTadd добавляет узел с заданным ключом в двоичное дерево поиска
func BSTadd(root *BSTNode, key int) *BSTNode {
    if root == nil {
        return NewNode(key)
    }

    if key < root.Key {
        root.Left = BSTadd(root.Left, key)
    } else if key > root.Key {
        root.Right = BSTadd(root.Right, key)
    }

    return root
}

//FindLevelSums находит суммы значений на каждом уровне двоичного дерева
func FindLevelSums(root *BSTNode) []int {
    levelSums := make([]int, 0)
    computeLevelSums(root, 1, &levelSums)
    return levelSums
}

//computeLevelSums рекурсивно вычисляет суммы значений на каждом уровне дерева
func computeLevelSums(root *BSTNode, currentLevel int, levelSums *[]int) {
    if root == nil {
        return
    }

    // Если текущий уровень превышает длину среза levelSums, добавляем новый элемент
    if currentLevel > len(*levelSums) {
        *levelSums = append(*levelSums, root.Key)
    } else {
        // Иначе, увеличиваем сумму текущего уровня
        (*levelSums)[currentLevel-1] += root.Key
    }

    // Рекурсивно обходим левое и правое поддерево
    computeLevelSums(root.Left, currentLevel+1, levelSums)
    computeLevelSums(root.Right, currentLevel+1, levelSums)
}

func main() {
    // Создаем двоичное дерево
    root := NewNode(5)
    BSTadd(root, 3)
    BSTadd(root, 6)
    BSTadd(root, 4)
    BSTadd(root, 2)
    BSTadd(root, 7)

    // Находим суммы каждого уровня и выводим результат
    levelSums := FindLevelSums(root)
    for i, sum := range levelSums {
        fmt.Printf("%d уровень – %d\n", i+1, sum)
    }
}

//если потомков нет как обрабатывается