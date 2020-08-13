package main

import "fmt"

func pointer() {

	// 変数（値を入れる箱） =========================
	num := 100
	// デフォルトが値
	fmt.Println("変数（値）", num)
	fmt.Println("変数（アドレス）", &num)

	// ポインタ変数(アドレスを入れる箱) ===============
	// defaultがアドレス
	// pointerがポインタ変数, *intがポインタ型
	var pointer *int = &num
	fmt.Println("ポインタ変数(アドレス)：", pointer)
	fmt.Println("ポインタ変数(値)：", *pointer)
}

// struct(構造体) =================================
// データ型の１つで複数の値を格納できる型
type Person struct {
	name string
	age  int
}
type Animal struct {
	name string
	size int
}

// 通常のインスタンス生成
func newUser(get_name string, get_age int) Person {
	// user := Person{get_name, get_age}
	var user Person
	user.name = get_name
	user.age = get_age
	return user
}

// ポインタ型で構造体からインスタンス生成
func newAnimal(get_name string, get_size int) *Animal {
	// newの戻り値はポインタ変数
	pet := new(Animal)
	pet.name = get_name
	pet.size = get_size
	return pet
}

// メソッドと値レシーバとポインタレシーバ ============================
type Vertex struct {
	width, height int
}

// ただの関数の場合
func Area(v Vertex) int {
	return v.width * v.height
}

// メソッドの場合(structと関数を結びつき)
// 関数の前にstructを()で入れる
func (v Vertex) Area() int {
	return v.width * v.height
}

/* ポインタレシーバ ------------------------------------------------------
・値レシーバ：元の値とは別のコピーした値が関数に渡されるので元の値は変更できない
・ポインタレシーバ：ポインタとして引数に渡すので関数内で元の値を変更できる
--------------------------------------------------------------------- */
func (v *Vertex) Scale(i int) {
	v.width = v.width * i
	v.height = v.height * i
}

func main() {
	// pointer()
	// firstUser := newUser("yuya", 24)
	// firstPet := newAnimal("dog", 12)
	// fmt.Println(p_firstUser, firstPet)
	v := Vertex{3, 4}
	fmt.Println(Area(v)) // 関数として
	v.Scale(10)
	fmt.Println(v.Area()) // メソッドとして
}
