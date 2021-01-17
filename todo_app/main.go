package main

import (
	"fmt"
	"golang_udemy/todo_app/app/controllers"
	"golang_udemy/todo_app/app/models"
	"golang_udemy/todo_app/config"
	"log"
)

func main() {

	controllers.StartMainServer()

	// fmt.Println(models.Db)

	// configPrint()
	// createUser()
	// getUser()
	// updateUser()
	// deleteUser()

	// createTodo()
	// getTodo()
	// getTodos()
	// getTodosByUser()
	// updateTodo()
	// deleteTodo()
}

func deleteTodo() {
	t, _ := models.GetTodo(1)
	fmt.Println(t)

	t.DeleteTodo()

	t, _ = models.GetTodo(1)
	fmt.Println(t)
}

func updateTodo() {
	t, _ := models.GetTodo(1)
	fmt.Println(t)

	t.Content = "update todo"
	t.UserID = 1

	err := t.UpdateTodo()
	if err != nil {
		log.Fatalln(err)
	}

	t, _ = models.GetTodo(1)
	fmt.Println(t)
}

func getTodosByUser() {
	u, _ := models.GetUser(2)
	user, _ := u.GetTodosByUser()

	fmt.Println(user)
}

func getTodos() {
	todos, _ := models.GetTodos()
	fmt.Println(todos)
}

func getTodo() {
	t, _ := models.GetTodo(1)
	fmt.Println(t)
}

func createTodo() {
	user, _ := models.GetUser(1)
	user.CreateTodo("first todo")
}

func deleteUser() {
	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.DeleteUser()

	u, _ = models.GetUser(1)
	fmt.Println(u)
}

func updateUser() {
	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()

	u, _ = models.GetUser(1)
	fmt.Println(u)
}

func getUser() {
	u, _ := models.GetUser(1)
	fmt.Println(u)
}

func createUser() {

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "password"
	fmt.Println(u)

	u.CreateUser()
}

func configPrint() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
