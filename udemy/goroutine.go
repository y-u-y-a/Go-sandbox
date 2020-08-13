package main

import (
	"fmt"
	"sync"
)

/* goroutine --------------------------------------------
goroutine：処理を仮想スレッドで起動する機能
-> 作業員を増やして仕事を分担するイメージ
-> goキーワードで呼び出し, 戻り値は取得できない
-> goroutineの処理が終わらなくても他の処理が終われば終了してしまう
対処法１：sync.WaitGroup
-> 戻り値がない
-> 処理が完了されると破棄される
------------------------------------------------------- */

func nomal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	// wgの終了を宣言
	wg.Done()
}

/* channel -----------------------------------------------
channel：プロセス間で値をやり取りするための機能
-> goroutineとの値をやり取りを可能にする
-> 「すれ違い」をなくすための待ち合わせを可能にする
make関数：初期化された領域への参照を返す(map, slice, channelのみ)
------------------------------------------------------- */
func gorCh(list []int, ch chan int) {
	sum := 0
	for _, v := range list {
		sum += v
	}
	ch <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	// channel生成(バッファのデフォルトは0)
	ch := make(chan int)
	go gorCh(s, ch)
	// 受け取り
	getCh := <-c
	fmt.Println(getCh)
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1) // goroutine数を指定
// 	go goroutine("world", &wg)
// 	nomal("Hello")
// 	// wgの終了を待つ
// 	wg.Wait()
// }
