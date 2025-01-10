package main

import (
	"todo-go/data"
	"todo-go/routes"
	"todo-go/tools"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	g := gin.Default()
	g.LoadHTMLGlob("templates/*")
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	tools.Crush(err)
	db.AutoMigrate(&data.Task{})

	g.GET("/", routes.GetTasksHandlerHtml(db))
	g.GET("/tasks", routes.GetTasksHandler(db))
	g.POST("/tasks", routes.CreateTaskHandler(db))
	g.DELETE("/tasks/:id", routes.DeleteTaskHandler(db))

	g.Run(":8080")
}
