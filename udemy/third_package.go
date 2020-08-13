package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	renamefile "github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := renamefile.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
