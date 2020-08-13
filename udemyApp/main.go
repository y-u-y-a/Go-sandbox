package main

import (
	// 標準
	"fmt"
	// ローカル
	"github.com/y-u-y-a/Go-sandbox/sampleApp/libs"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(libs.Average(list))

	person := libs.Person{Name: "yuya", Age: 24}
	fmt.Println(person)
	libs.Say()
}
