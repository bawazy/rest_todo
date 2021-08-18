package routes

import (
	"github.com/bawazy/rest_todo/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterTodoRoutes = func(router *mux.Router) {
	router.HandleFunc("/todo/", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/todo/{todoId}", controllers.GetTodobyId).Methods("GET")
	router.HandleFunc("/todo/{todoId}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{todoId}", controllers.DeleteTodo).Methods("DELETE")
}
