package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bawazy/rest_todo/pkg/models"
	"github.com/bawazy/rest_todo/pkg/utils"
	"github.com/gorilla/mux"
)

var NewTodo models.Todo

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	newTodos := models.GetAllTodos()
	res, _ := json.Marshal(newTodos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTodobyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("Error whilst Parsing")
	}
	todoDetails, _ := models.GetTodobyId(ID)
	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	CreateTodo := &models.Todo{}
	utils.ParseBody(r, CreateTodo)
	t := CreateTodo.CreateTodo()
	res, _ := json.Marshal(t)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("Error whilst Parsing")
	}
	todo := models.DeleteTodo(ID)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var updateTodo = &models.Todo{}
	utils.ParseBody(r, updateTodo)

	vars := mux.Vars(r)
	todoId := vars["todoId"]

	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("Error whilst Parsing")
	}

	todoDetails, db := models.GetTodobyId(ID)

	if updateTodo.Task != "" {
		todoDetails.Task = updateTodo.Task
	}

	if updateTodo.Completed {
		todoDetails.Completed = updateTodo.Completed
	}

	db.Save(&todoDetails)

	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
