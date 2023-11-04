package main
import( //импортирование 1 пакета: errors для создания и возврата ошибок
	"fmt"
	"errors" 
)

//Массив
type Array struct{
	data []string //хранение данных массива
	length int //хранение текущей длины массива
}

//NewArray создаёт и возвращает новый пустой объект массива (создание нового пустого массива)
func NewArray(size int) Array {
	return Array{
		data: make([]string, size),
	}
}

//Append добавляет элемент в конец массива
func (arr *Array) ArAdd(value string) string{
	for i:=0; i<len(arr.data); i++{
		if arr.data[i] == ""{
			arr.data[i] = value
			return ""
		}
	}
	return "Error: array is full!"
}

//Get возвращает элемент по индексу
func (arr *Array) AGet(index int) (string, error){
	if index<0 || index>= len(arr.data){
		return "0", errors.New("Index out of range")
}
return arr.data[index], errors.New("")
}

//Set устанавливает значение элемента по индексу 
func (arr *Array) ASet(index int, value string) error{
	if index < 0|| index >= len(arr.data){
		return errors.New("Index out of range")
}
arr.data[index] = value
return nil
}

//Length возвращает текущую длину массива
func (arr *Array) ALength() int{
	return arr.length
}

//Нахождение элемента, удаление и смещение
func (arr *Array) ADel(index int) string{
	if index>=0 && index<len(arr.data){
		for i:=index; i<len(arr.data)-1; i++{
			arr.data[i]=arr.data[i+1]   //смещение
		}
		arr.data[len(arr.data)-1]=""
		return ""
	}
	return "Error: Index out of range"
}

//Функция вывода
func (arr Array) ArrPrint(){
	for i:=0; i<len(arr.data); i++{
		if arr.data[i]!=""{
			fmt.Printf("%s ",arr.data[i])
		}
	} 
}



// SubArrays возвращает различные подмассивы массива
func SubArrays(arr Array) [][]string {
	subarrays := [][]string{{}} // Пустой подмассив всегда присутствует
	for i := 1; i < (1 << uint(len(arr.data))); i++ {
		subarray := make([]string, 0)
		for j, elem := range arr.data {
			if (i & (1 << uint(j))) != 0 {
				subarray = append(subarray, elem)
			}
		}
		// Проверка на уникальность подмассива
		unique := true
		for _, existing := range subarrays {
			if equalSlices(existing, subarray) {
				unique = false
				break
			}
		}
		if unique {
			subarrays = append(subarrays, subarray)
		}
	}
	return subarrays
}

// Функция для проверки на равенство двух срезов
func equalSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func main() {
	set0 := NewArray(3)
	set0.ArAdd("x")
	set0.ArAdd("y")
	set0.ArAdd("z")

	fmt.Println("All Subarrays:")
	subarrays := SubArrays(set0)
	for _, subarray := range subarrays {
		fmt.Printf("%v", subarray)
	}
}
