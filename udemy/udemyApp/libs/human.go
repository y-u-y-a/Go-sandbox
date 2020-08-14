package libs

import "fmt"

/* PublicとPrivate -------------------------------------------
他のファイルから呼び出していいものは大文字で, ダメのものは小文字で
Public：大文字
private：小文字
------------------------------------------------------------*/

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human！")
}
