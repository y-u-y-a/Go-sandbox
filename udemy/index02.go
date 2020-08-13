// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// // クロージャー ===============================================
// func incrementGenerator() func() int {
// 	// numがクロージャー
// 	num := 0
// 	return func() int {
// 		num++
// 		return num
// 	}
// }

// // 円周率円の面積を返す関数
// func circleArea(pi float64) func(radius float64) float64 {
// 	return func(radius float64) float64 {
// 		return pi * radius * radius
// 	}
// }

// // 可長変引数 ==================================================
// func args(params ...int) {
// 	fmt.Println(len(params), params)
// }

// // ログファイルの作成・書き出し ===================================
// func LoggingSettings(targetFile string) {
// 	// os.OpenFile(ファイル名, 用途)(なければ作成)
// 	logfile, _ := os.OpenFile(targetFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	multiLogFile := io.MultiWriter(os.Stdout, logfile)
// 	// ロギングする項目
// 	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
// 	log.SetOutput(multiLogFile)
// 	fmt.Println(logfile)
// }

// // 色々書き方 ==================================================
// func how() {
// 	// for文書き方-1
// 	for i := 0; i < 10; i++ {
// 		if i == 3 {
// 			fmt.Println("continue")
// 			continue
// 		}
// 		if i > 5 {
// 			fmt.Println("break")
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// 	// for文書き方-2
// 	sum := 1
// 	for sum < 10 {
// 		sum += sum
// 		fmt.Println(sum)
// 	}
// 	// range(foreachみたいな挙動) ==================================================
// 	list := []string{"python", "go", "php", "ruby"}
// 	for index, value := range list {
// 		fmt.Println(index, value)
// 	}
// 	// indexを使いたくない場合は_を代用
// 	for _, val := range list {
// 		fmt.Println(val)
// 	}
// 	// mapをrangeで
// 	mapList := map[string]int{"apple": 100, "banana": 200, "grape": 300}
// 	for key, value := range mapList {
// 		fmt.Println(key, value)
// 	}
// 	// keyだけ
// 	for key := range mapList {
// 		fmt.Println(key)
// 	}
// 	// valueだけ
// 	for _, value := range mapList {
// 		fmt.Println(value)
// 	}
// 	// switch =====================================================================
// 	os := "mac"
// 	switch os {
// 	case "mac":
// 		fmt.Println("Mac！")
// 	case "windows":
// 		fmt.Println("Windows！")
// 	default:
// 		fmt.Println("Default！")
// 	}
// 	// defer(遅延実行) =============================================================
// 	// 自関数が終了したら実行("World",1,2,3の順で実行される)
// 	// defer fmt.Println(3)
// 	// defer fmt.Println(2)
// 	// defer fmt.Println(1)
// 	// defer fmt.Println("World")
// 	fmt.Println("Hello")
// 	// log(ロギング) ===============================================================
// 	log.Printf("Logging！")
// 	log.Printf("%T %v", "test", "test")
// 	// コードが終了する
// 	// log.Fatalln("error！")
// 	// 例
// 	// _, err := os.Open("fnsofneo")
// 	// if err != nil {
// 	// 	log.Fatalln("Exit", err)
// 	// }
// }

// func main() {
// 	LoggingSettings("test.log")
// 	how()
// 	// args(10, 20)
// 	// args(10, 20, 30)
// 	// c1 := circleArea(3.14)
// 	// fmt.Println(c1(2))
// 	// c2 := circleArea(3)
// 	// fmt.Println(c2(2))
// }
