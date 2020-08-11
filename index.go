// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// // 一番初めに実行 ===================
// func init() {
// 	fmt.Println("Init！")
// }

// func study() {

// 	// 変数宣言 ====================
// 	var a int = 1
// 	var b bool = true

// 	// 関数内のみ
// 	aaa := 100
// 	bbb := false

// 	// 複数宣言(何も入れなければ初期値が表示される)
// 	// var (
// 	// 	i   int
// 	// 	f64 float64
// 	// )

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(aaa)
// 	fmt.Println(bbb)
// 	// fmt.Println(time.Now())
// 	// fmt.Println(user.Current())

// 	// 型変換 ====================
// 	// var x int = 1
// 	xx := float64(1)
// 	fmt.Printf("%T %v %f\n", xx, xx, xx)
// 	// string -> int
// 	s := "14"
// 	i, _ := strconv.Atoi(s)
// 	fmt.Printf("%T %v\n", i, i)

// 	// 配列(同じ型・指定した要素数) =======================
// 	array := [2]string{"test", "配列"}
// 	fmt.Println(array)
// 	// スライス(要素数を変更できる) =====================
// 	n := []int{100, 200, 300, 400, 500, 600, 700, 800}
// 	n = append(n, 900)
// 	fmt.Println(n[2:])

// 	// スライス(多階層) ===============================
// 	var sliceBoard = [][]int{
// 		[]int{1, 2, 3},
// 		[]int{4, 5, 6},
// 		[]int{7, 8, 9},
// 	}
// 	fmt.Println(sliceBoard)

// 	// make([]type, 要素数, 容量), cap ================
// 	slice := make([]int, 3, 5)
// 	fmt.Println(len(slice)) // 3
// 	fmt.Println(cap(slice)) // 5
// 	fmt.Println(slice)

// 	// map (連想配列(辞書), [key]{value}) ======================
// 	mapArray := map[string]int{"apple": 100, "banana": 200, "grape": 300}
// 	mapArray["new"] = 500
// 	fmt.Println(mapArray)
// }

// // 関数 =======================================================
// func add(x int, y int) int {
// 	return x + y
// }
// func duble(x int, y int) (int, int) {
// 	return x + y, x - 10
// }
// func cal(price int, item int) (result int, tax float64) {
// 	result = price * item
// 	return
// }

// // 実行される関数
// func main() {
// 	// study()
// 	rst1, rst2 := duble(10, 20)
// 	fmt.Println(rst1, rst2) // 30, 0
// 	cal1, cal2 := cal(10, 50)
// 	fmt.Println(cal1, cal2) // 500, 0
// }
