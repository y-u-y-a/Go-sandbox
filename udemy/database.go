package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// データベースとなるファイルを作成
	DbConnection, _ := sql.Open("sqlite3", "./db/sampleApp.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
				name STRING,
				age INT)`

	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln("コネクションエラー", err)
	}

	// SQLインジェクション対策で？を使用
	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err = DbConnection.Exec(cmd, "Nancy", 24)
	if err != nil {
		log.Fatalln("インサートエラー", err)
	}
	// "sqlite3 SQLファイル名"で作成したデータベースを開く
}
