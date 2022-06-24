package main

import (
	"fmt"
	"reflect"
)

type myInt int

func (m myInt) ToString() string {
	return fmt.Sprint(m)
}

func (m myInt) ToByte() []byte {
	return []byte(fmt.Sprint(m))
}

type myMap map[string]interface{}

func (m myMap) Insert(key string, val interface{}) {
	if m != nil {
		m[key] = val
	}
}

func (m myMap) ToString() string {
	if m != nil {
		return fmt.Sprint(m)
	}
	return ""
}

func (m myMap) ToByte() []byte {
	if m != nil {
		return []byte(fmt.Sprint(m))
	}
	return []byte("")
}

func main() {
	//  var num int = 100
	//	fmt.Println(num.ToString())
	var num myInt = 100
	str := num.ToString()

	fmt.Println("Value of num", num, "Type of num", reflect.TypeOf(num))
	fmt.Println("Value of str", str, "Type of str", reflect.TypeOf(str))

	//var num2 int = num this cannot be done becasue both of them are two different types
	var num2 int = int(num)
	fmt.Println("Value of num2", num2, "Type of num2", reflect.TypeOf(num2))

	var num3 int = 200

	str3 := myInt(num3).ToString()
	fmt.Println("Value of str3", str3, "Type of str3", reflect.TypeOf(str3))
	fmt.Println("num3 to byte", myInt(num3).ToByte())

	// map related stuff

	var mp myMap
	mp = make(map[string]interface{})
	mp.Insert("52213", "Guntur")
	mp.Insert("52214", "Vijayawada")
	mp.Insert("560086", "Yeshvantpur Bangalore")
	mp.Insert("Bangalore", 35)
	fmt.Println("Calling convert to string", mp.ToString())
	fmt.Println("Calling convert to []byte", mp.ToByte())
}
