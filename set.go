package main

import (
	"fmt"
)

//Определение структуры Set
type Set struct {
	elements      map[int]bool //множество, представленное в виде map для уникальных значений
	sortedElements []int  //срез для хранния отсортированных элементов
}

//Создание нового множества Set
func NewSet() *Set {
	s := &Set{}
	s.elements = make(map[int]bool)
	return s
}

//Добавление элемента в множество 
func (s *Set) SetAdd(element int) {
	s.elements[element] = true
}

//Удаление элемента из множества
func (s *Set) SetRemove(element int) {
	delete(s.elements, element)
}

//Проверка наличия элемента в множестве
func (s *Set) SetContains(element int) bool {
	return s.elements[element]
}

//Получение размера множества
func (s *Set) SetSize() int {
	return len(s.elements)
}

//Вывод элементов множества
func (s *Set) SetPrint() {
	for element := range s.elements {
		fmt.Print(" ", element)
	}
}

//Вычисление суммы элементов множества
func (s *Set) Sum() int {
	sum := 0
	for element := range s.elements {
		sum += element
	}
	return sum
}

//Сортировка элементов множества с помощью сортировки пузырьком
func (s *Set) BubbleSort() []int {
	elements := make([]int, 0, len(s.elements))
	for element := range s.elements {
		elements = append(elements, element)
	}

	n := len(elements)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if elements[i-1] < elements[i] {
				elements[i-1], elements[i] = elements[i], elements[i-1]
				swapped = true
			}
		}
	}

	return elements
}


func main() {
	set0 := NewSet()//множество
	set1 := NewSet()//подмножество
	set2 := NewSet()//подмножество 

	//Добавление элементов в множество set0
	set0.SetAdd(5)
	set0.SetAdd(8)
	set0.SetAdd(1)
	set0.SetAdd(14)
	set0.SetAdd(7)
	set0.BubbleSort()

	//Сортировка элементов множества set0 и вывод их
	//set0.BubbleSort()
	fmt.Println("set0:")
	set0.SetPrint()
	fmt.Println("")

	//Получение отсортированных элементов и разделение их на два подмножества set1 и set2
	sortedElements := set0.BubbleSort()
	//fmt.Println("Sorted Elements:", sortedElements)
	for i := range sortedElements {
		element:=sortedElements[i]
		//fmt.Println(element)
		if set1.Sum() > set2.Sum() {
			set2.SetAdd(element)
		} else {
			set1.SetAdd(element)
		}
	}

	// Получаем отсортированные элементы и выводим их

	//Вывод сумм элементов в подмножествах set1 и set2
	fmt.Println("sum set1", set1.Sum())
	fmt.Println("sum set2", set2.Sum())
}
