// Написать функцию, которая принимает на вход структуру in (struct или кастомную
// struct) и values map[string]interface{} (key - название поля структуры,
// которому нужно присвоить value этой мапы). Необходимо по значениям из мапы
// изменить входящую структуру in с помощью пакета reflect. Функция может
// возвращать только ошибку error. Написать к данной функции тесты (чем больше,
// тем лучше - зачтется в плюс).

package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type CustomStruct struct {
	key interface{}
}

type TestCustomStruct struct {
	a int
}

var ErrorNilMap = errors.New("ошибка. Мапа не может быть пустой\n")

func main() {

	var a = CustomStruct{12}
	var SuperMap = make(map[string]interface{})
	SuperMap["one"] = TestCustomStruct{2222}

	fmt.Println("before: ", a)

	if err := Remake(&a.key, SuperMap); err != nil {
		log.Fatalf("ошибка, %s", err)
	}

	fmt.Println("after: ", a)

}

func Remake(in *interface{}, m map[string]interface{}) error {
	if len(m) == 0 {
		err := ErrorNilMap
		return err
	}
	inVal := reflect.ValueOf(in)
	for _, v := range m {
		mVal := reflect.ValueOf(v)
		inVal.Elem().Set(mVal)
		break
	}
	return nil
}
