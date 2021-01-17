package main

import (
	"fmt"
	"golang_udemy/todo_app/app/models"
	"golang_udemy/todo_app/config"
	"log"
)

func main() {
	fmt.Println(models.Db)

	// configPrint()
	// createUser()
	// getUser()
	// updateUser()
	// deleteUser()

	createTodo()
}

func createTodo() {
	user, _ := models.GetUser(2)
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
