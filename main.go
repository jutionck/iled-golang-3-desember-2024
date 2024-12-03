package main

import (
	"fmt"

	"enigmacamp.com/iled-golang-3-desember-2024/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("failed to connect database")
	}

	// ini sebenernya cukup 1 kali aja di jalankan
	// db.AutoMigrate(&model.Todo{})

	// INSERT INTO todos (name, is_done) ...
	// db.Create(&model.Todo{
	// 	Name:   "Makan",
	// 	IsDone: false,
	// })

	var todo model.Todo
	// SELECT * FROM todos WHERE id = 1
	db.First(&todo, 1)

	fmt.Println("Todo:")
	fmt.Println("ID:", todo.ID)
	fmt.Println("Name:", todo.Name)
	fmt.Println("IsDone:", todo.IsDone)
}
