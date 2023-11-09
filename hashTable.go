//Метод открытой адресации (вариант 3)
package main

import (
	"errors"
	"fmt"
)

// Структура HashMap представляет хеш-таблицу с фиксированным размером массива table
type HashMap struct {
	table [512]*Pair
}

// Структура Pair представляет пару ключ-значение, которую храним в таблице
type Pair struct {
	key   string
	value string
}

// Функция calcHash вычисляет хеш ключа и возвращает индекс в массиве table
func calcHash(key string, size int) (int, error) {
	if len(key) == 0 {
		return 0, errors.New("no value")
	}
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % size, nil
}

// Метод Insert вставляет пару ключ-значение в хеш-таблицу
func (hmap *HashMap) Insert(key string, value string) error {
	p := &Pair{key, value}
	hash, err := calcHash(key, len(hmap.table))
	if err != nil {
		return errors.New("unacceptable key")
	}
	if hmap.table[hash] == nil {
		hmap.table[hash] = p
		return nil
	}
	if hmap.table[hash].key == key {
		return errors.New("this key already exists")
	}
	for i := (hash + 1) % len(hmap.table); i != hash; i = (i + 1) % len(hmap.table) {
		if hmap.table[i] == nil {
			hmap.table[i] = p
			return nil
		}
		if hmap.table[i].key == key {
			return errors.New("this key already exists")
		}
	}
	return errors.New("table is full")
}

// Метод Get получает значение по ключу из хеш-таблицы
func (hmap *HashMap) HGet(key string) (string, error) {
	hash, err := calcHash(key, len(hmap.table))
	if err != nil {
		return "", errors.New("unacceptable key")
	}
	if hmap.table[hash] != nil && hmap.table[hash].key == key {
		return hmap.table[hash].value, nil
	}
	for i := (hash + 1) % len(hmap.table); i != hash; i = (i + 1) % len(hmap.table) {
		if hmap.table[i] != nil && hmap.table[i].key == key {
			return hmap.table[i].value, nil
		}
	}
	return "", errors.New("no such key")
}

// Метод Del удаляет элемент из хеш-таблицы по ключу
func (hmap *HashMap) HDel(key string) error {
	hash, err := calcHash(key, len(hmap.table))
	if err != nil {
		return errors.New("unacceptable key")
	}
	if hmap.table[hash] == nil {
		return errors.New("nothing to delete")
	}
	if hmap.table[hash].key == key {
		hmap.table[hash] = nil
		return nil
	}
	for i := (hash + 1) % len(hmap.table); i != hash; i = (i + 1) % len(hmap.table) {
		if hmap.table[i] != nil && hmap.table[i].key == key {
			hmap.table[i] = nil
			return nil
		}
	}
	return errors.New("no such key")
}

// Метод OpenAddressInsert вставляет пару ключ-значение в хеш-таблицу с использованием открытой адресации.
// func (hmap *HashMap) OpenAddressInsert(key string, value string) error {
// 	p := &Pair{key, value}
// 	hash, err := calcHash(key, len(hmap.table))
// 	if err != nil {
// 		return errors.New("unacceptable key")
// 	}

// 	// Попробуйте вставить элемент с использованием открытой адресации
// 	for i := 0; i < len(hmap.table); i++ {
// 		newHash := (hash + i) % len(hmap.table)
// 		if hmap.table[newHash] == nil {
// 			hmap.table[newHash] = p
// 			return nil
// 		}
// 		if hmap.table[newHash].key == key {
// 			return errors.New("this key already exists")
// 		}
// 	}

// 	return errors.New("table is full")
// }

func main() {
	hmap := HashMap{}

	// Вставляем пару ключ-значение в хеш-таблицу
	err := hmap.Insert("1", "A")
	if err != nil {
		println(err)
	}

	// Получаем значение по ключу
	value, err := hmap.HGet("1")
	if err != nil {
		println(err)
	} else {
		println( value)
	}

	// Удаляем элемент по ключу
	err = hmap.HDel("1")
	if err != nil {
		println(err)
	}

	// Попытка получить удаленный элемент
	fmt.Println(hmap.HGet("1"))
}