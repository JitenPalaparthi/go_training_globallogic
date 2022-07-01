package main

import "fmt"

func finished(i int) {
	fmt.Println("Finished this function-->", i)
}
func largest(nums []int) {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func main() {
	// defer finished(1)
	// defer finished(2)
	// nums := []int{174, 136, 39, 368, 876, 343, 432, 562, 235}
	// largest(nums)

	/*a := 10
	defer fmt.Println("The deffered one-->", a)
	a = 20
	fmt.Println("normal one-->", a)*/

	str := "Hello, 世界!"
	// normal loop it takes byte
	// range loop it takes rune

	// for i := 0; i < len(str); i++ {
	// 	fmt.Print(str[i], " ")
	// }

	fmt.Println()
	for _, v := range str {
		defer fmt.Print(string(v))
	}

}

/*break , default , func , interface , select , case , defer , go ,
map , struct , chan , else , goto , package , switch , const ,
fallthrough , if , range , type , continue , for , import , return , var */
