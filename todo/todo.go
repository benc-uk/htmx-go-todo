package todo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID       int
	Title    string
	Done     bool
	Priority int
	Details  string
	Tags     []string
}

var todoList = []Todo{
	{1, "Buy milk", false, 1, "Low fat", []string{"shopping", "personal"}},
	{2, "Buy bread", true, 1, "Whole grain", []string{"shopping", "personal"}},
	{3, "Buy cheese", false, 2, "Swiss", []string{"shopping", "personal"}},
	{4, "Buy beer", true, 3, "IPA", []string{"shopping", "personal"}},
	{5, "Buy wine", false, 1, "Merlot", []string{"shopping", "personal"}},
	{6, "Buy eggs", false, 1, "Free range", []string{"shopping", "personal"}},
	{7, "Buy butter", false, 2, "Salted", []string{"shopping", "personal"}},
	{8, "Buy jam", true, 1, "Strawberry", []string{"shopping", "personal"}},
}

func AddHandlers(e *echo.Echo) {
	e.GET("/data/todos", func(c echo.Context) error {
		err := c.Render(http.StatusOK, "data/todo/list", todoList)
		return err
	})

	e.GET("/data/todos/:id", func(c echo.Context) error {
		todo := findTodoByID(c)
		if todo == nil {
			return c.HTML(http.StatusNotFound, "")
		}

		return c.Render(http.StatusOK, "data/todo/single", todo)
	})

	e.GET("/data/todos/:id/edit", func(c echo.Context) error {
		todo := findTodoByID(c)
		if todo == nil {
			return c.HTML(http.StatusNotFound, "")
		}

		return c.Render(http.StatusOK, "data/todo/edit", todo)
	})

	e.PUT("/data/todos/:id", func(c echo.Context) error {
		todo := findTodoByID(c)
		if todo == nil {
			return c.HTML(http.StatusNotFound, "")
		}

		p, _ := c.FormParams()
		log.Println("### Edit todo:", p)

		done := c.FormValue("done")
		title := c.FormValue("title")
		details := c.FormValue("details")
		priority := c.FormValue("priority")

		if done != "" {
			doneBool, err := strconv.ParseBool(done)
			if err == nil {
				todo.Done = doneBool
			}
		}

		if title != "" {
			todo.Title = title
		}

		if details != "" {
			todo.Details = details
		}

		if priority != "" {
			priorityInt, err := strconv.Atoi(priority)
			if err == nil {
				todo.Priority = priorityInt
			}
		}

		todoList[getTodoIndexByID(todo.ID)] = *todo

		return c.Render(http.StatusOK, "data/todo/single", todo)
	})

	e.DELETE("/data/todos/:id", func(c echo.Context) error {

		// delete from todoList
		todo := findTodoByID(c)
		if todo == nil {
			return c.HTML(http.StatusNotFound, "")
		}

		todoList = append(todoList[:getTodoIndexByID(todo.ID)], todoList[getTodoIndexByID(todo.ID)+1:]...)

		return c.HTML(http.StatusOK, "")
	})
}

func findTodoByID(c echo.Context) *Todo {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}

	for _, todo := range todoList {
		if todo.ID == idInt {
			return &todo
		}
	}

	return nil
}

func getTodoIndexByID(id int) int {
	for i, todo := range todoList {
		if todo.ID == id {
			return i
		}
	}

	return -1
}