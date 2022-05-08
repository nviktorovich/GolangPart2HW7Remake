package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

type CheckSets struct {
	correctAns interface{}
	retErr     error
	innerMap   map[string]interface{}
}

var Sets = []CheckSets{
	{reflect.ValueOf(12), nil, map[string]interface{}{"key": 12}},
	{reflect.ValueOf("abs"), nil, map[string]interface{}{"key": "abs"}},
}

type TestStruct struct {
	key interface{}
}

func TestRemake(t *testing.T) {
	fmt.Println("TableTest")
	var A = TestStruct{"test"}
	for _, set := range Sets {
		err := Remake(&A.key, set.innerMap)
		if err != nil {
			log.Fatalln("ошибка тестирования")
		}
		if set.correctAns != reflect.ValueOf(A.key) {
			fmt.Printf("\nошибка. Тестируемый набор: %v\nОжидалось: %v\nПолучено: %v\n\n",
				set,
				set.correctAns,
				reflect.ValueOf(A.key),
			)
		}

	}

}

func TestRemakeErrorsInverse(t *testing.T) {
	fmt.Println("Test inverse error (empty map)")
	var A = TestStruct{"test"}
	err := Remake(&A.key, map[string]interface{}{})
	if err == nil {
		fmt.Printf("\nошибка. При передаче пустой мапы не была возвращена ошибка\n")
	}
}

func TestRemakeErrors(t *testing.T) {
	fmt.Println("Test error (empty map)")
	var A = TestStruct{"test"}
	err := Remake(&A.key, map[string]interface{}{})
	if err != ErrorNilMap {
		fmt.Printf("\nошибка. Ожидалась ошибка: \t\t%v\nПолучена ошибка: \t\t%v\n", ErrorNilMap, err)
	}
}
