package main

import "fmt"

func pointer() {

	// ポインタ ====================================================================
	// 変数(値を指した箱)
	num := 100
	fmt.Println(&num) // アドレスを取得

	// ポインタ変数(アドレスを指した箱)
	p := &num       // アドレスを格納
	fmt.Println(*p) // アドレス内に格納されている値取得

	// new(空のメモリを確保) =========================================================
	newNum := new(int)
	fmt.Println(&newNum)

	var newNum2 *int
	fmt.Println(newNum2) // nill
}

func main() {
	pointer()
}
