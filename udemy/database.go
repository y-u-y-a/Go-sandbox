package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)
	defer row.Close()

	// スライスで宣言
	var pp []Person

	for rows.Next() {
		var p Person
		// // エラーを１つずつ確認
		// err := rows.Scan(&p.Name, &p.Age)
	}
	// まとめてエラーチェック
	err = rows.Err()
	// エラーハンドリング
	if err != nil {
		log.Fatalln("エラー", err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}

// func main() {
// 	// データベースとなるファイルを作成
// 	DbConnection, _ := sql.Open("sqlite3", "./db/sample.sql")
// 	defer DbConnection.Close()
// 	cmd := `CREATE TABLE IF NOT EXISTS person(
// 				name STRING,
// 				age INT)`

// 	_, err := DbConnection.Exec(cmd)
// 	if err != nil {
// 		log.Fatalln("コネクションエラー", err)
// 	}

// 	// SQLインジェクション対策で？を使用
// 	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
// 	_, err = DbConnection.Exec(cmd, "Nancy", 24)
// 	if err != nil {
// 		log.Fatalln("インサートエラー", err)
// 	}
// 	// "sqlite3 SQLファイル名"で作成したデータベースを開く
// }
