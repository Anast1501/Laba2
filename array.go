//Различные подмассивы (вариант 3)

//Необходимо реализовать алгоритм, который выводит все различные
//подмассивы массива.

package main 

import (
	"fmt"
)

//Массив 
type Array struct {
	data []string //хранение данных массива
}

//NewArray создаёт и возвращает новый пустой объект массива (создание нового пустого массива)
func NewArray(size int) Array {
	return Array {
		data: make([]string, size), 
	}
}

//Сортировка подмассивов по их длине (сортировка вставками)
func InsertionSortSubarrays(subarrays [][]string) { //отсортировка подмассивов по длине в порядке возрастания
	for i := 1; i<len(subarrays); i++ {
		key := subarrays[i] //текущий подмассив
		j := i-1
		for j>=0 && len(subarrays[j])>len(key) { //сравнение длины подмассивов и производим сдвиг элементов, чтоб вставить key на правильное место
			subarrays[j+1] = subarrays[j]
			j--
		}
		subarrays[j+1] = key
	}
}


//ArAdd добавляет элемент в конец массива
func (arr *Array) ArAdd(value string) string {
	for i := 0; i<len(arr.data); i++ {
		if arr.data[i] == "" {
			arr.data[i] = value
			return ""
	}
}
return "Error: array is full!"
}

//AGet возвращает элемент по индексу 
func (arr *Array) AGet(index int) (string, error) {
	if index<0 || index>=len(arr.data) {
		return "0", fmt.Errorf("Index out of range")
	}
	return arr.data[index], nil
}

//ASet устанавливает значение элемента по индексу 
func (arr *Array) ASet(index int, value string) error {
	if index<0 || index>=len(arr.data) {
		return fmt.Errorf("Index out of range")
}
arr.data[index] = value
return nil
}

//ALength возвращает текущую длину массива
func (arr *Array) ALength() int {
	return len(arr.data)
}

//ADel нахождение элемента, удаление и смещение
func (arr *Array) ADel(index int) string {
	if index>=0 && index<len(arr.data) {
		for i:=index; i<len(arr.data)-1; i++ {
			arr.data[i]=arr.data[i+1] //смещение
		}
		arr.data[len(arr.data)-1] = ""
		return ""
	}
	return "Error: Index out of range"
}

//Функция вывода 
func (arr Array) ArrPrint() {
	for i := 0; i<len(arr.data); i++ {
		if arr.data[i] != "" {
			fmt.Printf("%s ", arr.data[i])
		}
	}
}

//Функция SubArrays генерирует и возвращает все подмассивы массива
func SubArrays(arr Array) [][]string {
	subarrays := make([][]string, 0) //создание пустого среза срезов строк для хранения подмассивов

	for i := 0; i < (1 << int(len(arr.data))); i++ {  //битовый сдвиг влево для создания числа (от 0 до 2^n-1) //Эта строка использует   для приведения длины массива к целому числу без знака и выполнения побитовых операций в цикле.
		subarray := make([]string, 0) //создание текущего среза строк для текущего подмассива
		for j, elem := range arr.data { //перебор элементов массива вместе с их индексами
			if (i & (1<<int(j))) != 0 {
				subarray = append(subarray, elem) //добавление текущего подмассива в общий список подмассивов
			}
		}
		subarrays = append(subarrays, subarray)
	}
	return subarrays
}


func main() {
	//Создание нового массива и добавление элементов
	array := NewArray(3)
	array.ArAdd("x")
	array.ArAdd("y")
	array.ArAdd("z")

	//Генерация и сортировка подмассивов
	subarrays := SubArrays(array)
	InsertionSortSubarrays(subarrays)

	//Вывод результатов
	fmt.Println("AllSubarrays")
	for _, subarray := range subarrays {
		fmt.Printf("%v", subarray)
}
}

//не конкретные удовлетворяющие индексы а подмассив

//проверка 