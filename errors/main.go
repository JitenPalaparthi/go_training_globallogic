package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		//	return
	} else {
		fmt.Println(f.Name())
		buf := make([]byte, 1)
		for {
			n, err := f.Read(buf)

			if err != nil && n == 0 {
				// if err == io.EOF {
				// 	fmt.Println("err")
				// }
				break
			} else if err != nil {
				fmt.Println("Some error ", err)
			}
			fmt.Print(string(buf))
		}
	}

	var a, b float64
	a, b = 10.23, 20.34
	fmt.Println()
	fmt.Println(add(a, b))
	fmt.Println(addM(123.45, 145.578))
	fmt.Println(addM("GlobalLogic", " Golang Training"))
	fmt.Println(addM(true, false))
}

var (
	ERROR_TYPES = errors.New("Invalid type")
)

func addM(a, b interface{}) (interface{}, error) {
	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)
	if t1 != t2 {
		return nil, errors.New("a and b types does not match")
	}
	switch a.(type) {
	case int:
		return a.(int) + b.(int), nil
	case int16:
		return a.(int16) + b.(int16), nil
	case int32:
		return a.(int32) + b.(int32), nil
	case int64:
		return a.(int64) + b.(int64), nil
	case uint:
		return a.(uint) + b.(uint), nil
	case uint16:
		return a.(uint16) + b.(uint16), nil
	case uint32:
		return a.(uint32) + b.(uint32), nil
	case uint64:
		return a.(uint64) + b.(uint64), nil
	case float32:
		return a.(float32) + b.(float32), nil
	case float64:
		return a.(float64) + b.(float64), nil
	case string:
		return a.(string) + b.(string), nil
	default:
		return nil, ERROR_TYPES
	}
}

func add(a, b interface{}) (interface{}, error) {
	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)
	if t1 != t2 {
		return nil, errors.New("a and b types does not match")
	} else if t1.Name() != "int" && t1.Name() != "float32" {
		return nil, fmt.Errorf("a and b are invalid types")
	}
	if t1.Name() == "int" {
		c := a.(int) + b.(int)
		return c, nil
	} else if t1.Name() == "float32" {
		c := a.(float32) + b.(float32)
		return c, nil
	}
	return nil, ERROR_TYPES
}
