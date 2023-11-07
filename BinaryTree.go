/*package main

import (
    "fmt"
)

type BSTNode struct {
    Key   int
    Left  *BSTNode
    Right *BSTNode
}

func NewNode(key int) *BSTNode {
    return &BSTNode{Key: key, Left: nil, Right: nil}
}

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

func LevelSum(root *BSTNode, level int) int {
    if root == nil {
        return 0
    }

    if level == 1 {
        return root.Key
    }

    return LevelSum(root.Left, level-1) + LevelSum(root.Right, level-1)
}

func FindKthLevelSum(root *BSTNode, k int) int {
    maxLevel := 0
    result := 0

    for level := 1; ; level++ {
        sum := LevelSum(root, level)
        if sum == 0 {
            break
        }
        if level == k {
            result = sum
        }
        maxLevel = level
    }

    if k > maxLevel {
        return 0 // Уровень k больше, чем максимальный уровень в дереве
    }

    return result
}

func main() {
    // Создаем двоичное дерево
    root := NewNode(1)
    BSTadd(root, 2)
    BSTadd(root, 3)
    BSTadd(root, 4)
    BSTadd(root, 5)
    BSTadd(root, 6)
    BSTadd(root, 7)
	BSTadd(root, 8)
	BSTadd(root, 9)

    k := 9
    result := FindKthLevelSum(root, k)
    fmt.Printf("%d уровень – %d\n", k, result)
}

*/


/*- !!!!!!!
package main

import (
    "fmt"
)

type BSTNode struct {
    Key   int
    Left  *BSTNode
    Right *BSTNode
}

func NewNode(key int) *BSTNode {
    return &BSTNode{Key: key, Left: nil, Right: nil}
}

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

func LevelSum(root *BSTNode, level int) int {
    if root == nil {
        return 0
    }

    if level == 1 {
        return root.Key
    }

    return LevelSum(root.Left, level-1) + LevelSum(root.Right, level-1)
}

func FindKthLevelSum(root *BSTNode, k int) int {
    if root == nil || k <= 0 {
        return 0
    }

    if k == 1 {
        return root.Key
    }

    return FindKthLevelSum(root.Left, k-1) + FindKthLevelSum(root.Right, k-1)
}

func main() {
    // Создаем двоичное дерево
    root := NewNode(1)
    BSTadd(root, 2)
    BSTadd(root, 3)
    BSTadd(root, 4)
    BSTadd(root, 5)
    BSTadd(root, 6)
    BSTadd(root, 7)
    BSTadd(root, 8)
    BSTadd(root, 9)

    k := 3 //порядковый номер уровня в двоичном дереве
    result := FindKthLevelSum(root, k)
    fmt.Printf("%d уровень – %d\n", k, result)
}
*/



package main

import (
    "fmt"
)

type BSTNode struct {
    Key   int
    Left  *BSTNode
    Right *BSTNode
}

func NewNode(key int) *BSTNode {
    return &BSTNode{Key: key, Left: nil, Right: nil}
}

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

func FindKthLevelSum(root *BSTNode, k int, currentLevel int) int {
    if root == nil {
        return 0
    }

    if currentLevel == k {
        return root.Key
    }

    leftSum := FindKthLevelSum(root.Left, k, currentLevel+1)
    rightSum := FindKthLevelSum(root.Right, k, currentLevel+1)

    return leftSum + rightSum
}


func main() {
    // Создаем двоичное дерево
    //var sum int
    root := NewNode(5)
    BSTadd(root, 3)
    BSTadd(root, 6)
    BSTadd(root, 4)
    BSTadd(root, 2)
    BSTadd(root, 7)
    k := 2 // Порядковый номер уровня в двоичном дереве
    result := FindKthLevelSum(root, k, 1) // Начинаем с первого уровня
    fmt.Printf("%d уровень – %d\n", k, result)
}
