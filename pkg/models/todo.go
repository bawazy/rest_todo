package models

import (
	"github.com/bawazy/rest_todo/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	Task      string `gorm:"" json:"task"`
	Completed bool   `json:"completed"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}

func (t *Todo) CreateTodo() *Todo {
	db.NewRecord(t)
	db.Create(&t)
	return t
}

func GetAllTodos() []Todo {
	var Todos []Todo
	db.Find(&Todos)
	return Todos
}
func GetTodobyId(Id int64) (*Todo, *gorm.DB) {
	var getTodo Todo
	db.Where("ID=?", Id).Find(&getTodo)
	return &getTodo, db
}

func DeleteTodo(Id int64) Todo {
	var todo Todo
	db.Where("ID=?", Id).Delete(todo)
	return todo
}
