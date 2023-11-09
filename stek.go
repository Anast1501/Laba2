//Обратная польская запись (вариант 2)
package main

import (
	"errors"
	"fmt"
)

//SNode - это структура, представляющая элемент (узел) стека. Узел содержит два поля: next, которое указывает на следующий узел в стеке, и val, которое хранит целочисленное (строковые) значение элемента стека.
type SNode struct { //стековая структура
	next *SNode //указатель на следующий узел в стеке
	val  string //значения типа string , хранящиеся в элементе стека
}

//Структура Stack представляет стек и содержит указатель на верхний элемент 
type Stack struct {
	head *SNode   //head - одно поле, которое указывает на верхний элемент стека
}

//SPush добавляет новый элемент в стек
func (s *Stack) SPush(value string) {
	newNode := &SNode{
		val:  value,
		next: s.head,
	}
	s.head = newNode  //обновление вершины стека на новый узел
}

//SPop удаляет и возвращает верхний элемент стека
func (s *Stack) SPop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("Stack is empty")  //вызов проверки на пустоту стека
	}
	value := s.head.val //сохранение значение верхнего элемента стека
	s.head = s.head.next  //обновление вершины стека
	return value, nil //возврат значения
}

//IsEmpty возвращает true, если стек пуст, иначе false 
func (s *Stack) IsEmpty() bool {
	return s.head == nil  //если равны, то стек считается пустым => true
}

// evaluateRPN выполняет вычисление выражения в обратной польской записи
func evaluateRPN(expression string) (int, error) {
	stack := Stack{} //создание пустого стека
	tokens := splitString(expression, " ") //разбиение входной строки на токены (части выражения) с помощью функции splitString
	//перебор токенов из входного выражения     Если токен является оператором, то программа извлекает два операнда из стека, выполняет операцию и результат снова помещает в стек. Если токен не является оператором, он просто помещается в стек.
	for _, token := range tokens {
		if isOperator(token) {
			if stack.IsEmpty() {
				return 0, errors.New("Invalid RPN expression")
			}

			operand2, err := sPopInt(&stack)
			if err != nil {
				return 0, err
			}
			operand1, err := sPopInt(&stack)
			if err != nil {
				return 0, err
			}

			result := performOperation(token, operand1, operand2)
			stack.SPush(fmt.Sprintf("%d", result))
		} else {
			stack.SPush(token)
		}
	}

	//
	if stack.IsEmpty() {
		return 0, errors.New("Invalid RPN expression")
	}
	//извлечение результата из стека и преобразования в целое число
	result, err := sPopInt(&stack)
	if err != nil {
		return 0, err
	}

	return result, nil //результат вычислений возвращается как целое число, а при наличии ошибки nil
}

//sPopInt извлекает значение из стека и преобразует его в целое число
func sPopInt(s *Stack) (int, error) {
	value, err := s.SPop()
	if err != nil {
		return 0, err
	}
	var intValue int
	_, err = fmt.Sscanf(value, "%d", &intValue)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

//splitString разбивает строку на подстроку (tokens) с использованием заданного разделителя (пробела)
func splitString(stroka, sep string) []string { //разбиение входной строки s на подстроки (токены) с использованием заданного разделителя (пробела) sep
	tokens := make([]string, 0) //создаём пустой срез строк tokens, который будет содержать разделённые токены  tokens вся строка
	token := "" //инциализируем пустую строку token для хранения текущего токена     token - элемент строки
	for _, char := range stroka { //итерация по каждому символу char во входной stroka
		if char == ' ' { //если текущий символ равен пробелу, то текущий токен завершается
			if token != "" { //если текущий символ не является пробелом, то он добавляется в текущий token, с использованием оператора += 
				tokens = append(tokens, token)
				token = ""
			}
		} else {
			token += string(char)
		}
	}
	if token != "" {
		tokens = append(tokens, token)
	}
	return tokens //возвращение tokens, содержащий все разделённые токены 
}

//isOperator проверяет, является ли подстрока (tokens) оператором
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

//performOperation выполняет операцию на двух операндах в зависимости от оператора
func performOperation(operator string, operand1 int, operand2 int) int {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("Division by zero")
			return 0
		}
		return operand1 / operand2
	default:
		fmt.Println("Unknown operator")
		return 0
	}
}

func main() {
	expression := "6 4 8 * +" //обратная польская запись
	//expression := "17 9 2 * +" //обратная польская запись
	result, err := evaluateRPN(expression)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
