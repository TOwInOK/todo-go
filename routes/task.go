package routes

import (
	"net/http"
	"time"
	"todo-go/data"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// get tasks
func GetTasksHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tasks []data.Task
		db.Find(&tasks)
		ctx.JSON(http.StatusOK, tasks)
	}
}

// create task
func CreateTaskHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var task data.Task
		if err := ctx.ShouldBindJSON(&task); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		task.CreatedAt = time.Now()
		db.Create(&task)
		ctx.JSON(http.StatusOK, task)
	}
}

// delete task
func DeleteTaskHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		// Находим запись перед удалением
		var task data.Task
		if err := db.First(&task, id).Error; err != nil {
			// Если запись не найдена
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
				return
			}

			// Обработка других ошибок
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
			return
		}

		// Удаляем запись
		if err := db.Delete(&task).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
			return
		}

		// Возвращаем удалённый объект
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Task deleted",
			"task":   task,
		})
	}
}

func GetTasksHandlerHtml(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Получаем задачи из базы данных
		var tasks []data.Task
		db.Find(&tasks)

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"tasks": &tasks,
		})
	}
}
