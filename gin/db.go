package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

// gorm: デベロッパーフレンドリーを目指した、Go言語のORMライブラリ
// [GORMガイド](https://gorm.io/ja_JP/docs/index.html)

const sqlite_filename = "test.sqlite3"

// モデル
type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// DB初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", sqlite_filename)
	if err != nil {
		panic(fmt.Sprintf("db open filed [dbInit] %s", err))
	}
	// Auto Migration https://gorm.io/ja_JP/docs/migration.html
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

// インサート
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", sqlite_filename)
	if err != nil {
		panic(fmt.Sprintf("db open filed [dbInsert] %s", err))
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

// アップデート
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", sqlite_filename)
	if err != nil {
		panic(fmt.Sprintf("db open filed [dbUpdate] %s", err))
	}
	// 該当レコードをidで取得して更新
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// 削除
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", sqlite_filename)
	if err != nil {
		panic(fmt.Sprintf("db open filed [dbDelete] %s", err))
	}
	// 該当レコードをidで取得して削除
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

// 全て取得
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", sqlite_filename)
	if err != nil {
		panic(fmt.Sprintf("db open filed [dbGetAll] %s", err))
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// １つ取得
func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", sqlite_filename)
	if err != nil {
		panic(fmt.Sprintf("db open filed [dbGetOne] %s", err))
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
