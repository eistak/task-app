package pkg

import (
	"log"
	"net/http"
	"time"

	"go-ws/task-app/models"
	"go-ws/task-app/tools"
	"go-ws/task-app/viewmodel"

	"github.com/labstack/echo/v4"
)

var tasks []models.Task

// GetTasks finds all books
func GetTasks() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		db := tools.ConnectDB()
		tasks = []models.Task{}
		result := db.Find(&tasks)
		if result.Error != nil {
			log.Println("failed to get tasks.")
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		// change tasks from array to struct
		vmTasks := viewmodel.MakeTaskViewModel(tasks)

		return c.JSON(http.StatusOK, vmTasks)
	}
}

// GetTask finds a book by id
func GetTask() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		db := tools.ConnectDB()
		var task models.Task
		err = c.Bind(&task)
		if err != nil {
			log.Printf("failed to bind task. %+v", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		result := db.First(&task)
		if result.RowsAffected == 0 {
			log.Printf("row was not found. %+v", result)
			return c.JSON(http.StatusBadRequest, result.Error)
		}
		if result.Error != nil {
			log.Println(result.Error)
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusOK, &task)
	}
}

// AddTask adds a task
func AddTask() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		db := tools.ConnectDB()
		var task models.Task
		err = c.Bind(&task)
		log.Println(task)
		if err != nil {
			return err
		}

		now := time.Now()
		task.CreatedAt = now
		task.UpdatedAt = now

		result := db.Create(&task)
		if result.Error != nil {
			log.Println(result.Error)
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusCreated, task.ID)
	}
}

// UpdateTask updates a task by id
func UpdateTask() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		db := tools.ConnectDB()
		var task models.Task
		if err := c.Bind(&task); err != nil {
			log.Println("failed to bind task. err.Error()", err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if task.ID == 0 || task.Name == "" {
			errMsg := "id and name are required."
			log.Println(errMsg, task)
			return c.JSON(http.StatusBadRequest, errMsg)
		}

		task.UpdatedAt = time.Now()

		result := db.Save(&task)
		if result.Error != nil {
			log.Println(result.Error)
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		log.Println("updated task.", task)
		return c.JSON(http.StatusAccepted, task.ID)
	}
}

// RemoveTask deletes a task by id
func RemoveTask() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		db := tools.ConnectDB()
		var task models.Task
		err = c.Bind(&task)
		if err != nil {
			log.Println("failed to bind task to delete.")
			return c.JSON(http.StatusBadRequest, nil)
		}
		result := db.First(&task)
		if result.RowsAffected == 0 {
			log.Printf("row was not found. %+v", result)
			return c.JSON(http.StatusBadRequest, result.Error)
		}
		result = db.Delete(&task)
		if result.Error != nil {
			log.Printf("failed to delete book. task:%+v\n", task)
			return c.JSON(http.StatusInternalServerError, task)
		}
		log.Printf("successfully deleted book. task:%+v\n", task)
		return c.JSON(http.StatusAccepted, nil)
	}
}
