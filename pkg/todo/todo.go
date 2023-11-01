package todo

import (
	"htmx-go-todo/pkg/middleware"
	"math/rand"
	"net/http"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Todo is a struct representing a todo item
type Todo struct {
	ID       int
	Title    string
	Done     bool
	Priority int
	Details  string
	Tags     []string
}

// Our in memory todo list
var todoList = []Todo{
	{1, "Buy a goat", false, 1, "One that's nice and friendly", []string{"shopping", "personal"}},
	{2, "Learn to juggle", false, 2, "Buy some balls and start practicing", []string{"personal", "skills"}},
	{3, "Conquer the world", false, 1, "Start by taking over a small country", []string{"ambitious"}},
	{4, "Become a ninja", false, 3, "Train in the art of stealth and combat", []string{"personal", "skills"}},
	{5, "Learn to play the guitar", false, 2, "Start with some basic chords", []string{"personal", "skills"}},
	{6, "Build a robot", false, 1, "Gather materials and start building", []string{"hobby", "technology"}},
	{7, "Write a book", false, 2, "Start with an outline", []string{"personal", "skills"}},
	{8, "Cook a 3 course meal", false, 3, "Start with a main course", []string{"personal", "skills"}},
	{9, "Learn to fly", false, 1, "Start with a small plane", []string{"personal", "skills"}},
	{10, "Visit the moon", false, 1, "Start by building a rocket", []string{"ambitious"}},
	{11, "Learn to draw", false, 2, "Start with a pencil and paper", []string{"personal", "skills"}},
}

const pageSize = 10

func AddHandlers(e *echo.Echo) {
	// The base path for all todo handlers
	g := e.Group("/data/todos")

	// This blocks users from directlyaccessing these data endpoints
	g.Use(middleware.HTMXGuard())

	//
	// List all todo using GET
	//
	g.GET("", func(c echo.Context) error {
		offset := c.QueryParam("offset")

		offsetInt := 0
		if offset != "" {
			offsetInt, _ = strconv.Atoi(offset)
		}

		hasMore := true
		upperOffset := offsetInt + pageSize
		if upperOffset >= len(todoList) {
			upperOffset = len(todoList)
			hasMore = false
		}

		err := c.Render(http.StatusOK, "todo/list", map[string]any{
			"todos":   todoList[offsetInt:upperOffset],
			"offset":  offsetInt + pageSize,
			"hasMore": hasMore,
		})

		return err
	})

	//
	// Fetch todo using GET for viewing
	//
	g.GET("/:id", func(c echo.Context) error {
		todo := findTodoByID(c)
		if todo == nil {
			return c.HTML(http.StatusNotFound, "")
		}

		return c.Render(http.StatusOK, "todo/single", todo)
	})

	//
	// Fetch todo using GET for editing
	//
	g.GET("/:id/edit", func(c echo.Context) error {
		todo := findTodoByID(c)
		if todo == nil {
			return c.HTML(http.StatusNotFound, "")
		}

		return c.Render(http.StatusOK, "todo/edit", todo)
	})

	//
	// Update todo using PUT
	//
	g.PUT("/:id", func(c echo.Context) error {
		todo := findTodoByID(c)
		if todo == nil {
			return c.HTML(http.StatusNotFound, "")
		}

		formValues, err := c.FormParams()
		if err != nil {
			return c.HTML(http.StatusInternalServerError, "")
		}

		if formValues.Has("done") {
			doneBool, err := strconv.ParseBool(c.FormValue("done"))
			if err == nil {
				todo.Done = doneBool
			}
		}

		if formValues.Has("title") {
			title := c.FormValue("title")
			if title == "" {
				return c.HTML(http.StatusBadRequest, "Title cannot be empty")
			}

			todo.Title = c.FormValue("title")
		}

		if formValues.Has("details") {
			todo.Details = c.FormValue("details")
		}

		if formValues.Has("priority") {
			priorityInt, err := strconv.Atoi(c.FormValue("priority"))
			if err == nil {
				if priorityInt < 1 || priorityInt > 3 {
					return c.HTML(http.StatusBadRequest, "Priority must be between 1 and 3")
				}

				todo.Priority = priorityInt
			}
		}

		// Mutate todoList with updated todo
		todoList[getTodoIndexByID(todo.ID)] = *todo

		return c.Render(http.StatusOK, "todo/single", todo)
	})

	//
	// Delete a todo using DELETE
	//
	g.DELETE("/:id", func(c echo.Context) error {
		// delete from todoList
		todo := findTodoByID(c)
		if todo == nil {
			return c.HTML(http.StatusNotFound, "")
		}

		todoList = append(todoList[:getTodoIndexByID(todo.ID)], todoList[getTodoIndexByID(todo.ID)+1:]...)

		return c.HTML(http.StatusOK, "")
	})

	g.POST("", func(c echo.Context) error {
		done := c.FormValue("done")
		title := c.FormValue("title")
		details := c.FormValue("details")
		priority := c.FormValue("priority")

		doneBool, err := strconv.ParseBool(done)
		if err != nil {
			doneBool = false
		}

		priorityInt, err := strconv.Atoi(priority)
		if err != nil {
			return c.HTML(http.StatusBadRequest, "Priority must be a number")
		}

		if priorityInt < 1 || priorityInt > 3 {
			return c.HTML(http.StatusBadRequest, "Priority must be between 1 and 3")
		}

		if title == "" {
			return c.HTML(http.StatusBadRequest, "Title cannot be empty")
		}

		todo := Todo{
			ID:       rand.Intn(80000),
			Title:    title,
			Done:     doneBool,
			Priority: priorityInt,
			Details:  details,
		}

		// Add to todoList at the start
		todoList = append([]Todo{todo}, todoList...)

		// sort todoList by priority (highest first)
		sort.Slice(todoList, func(i, j int) bool {
			return todoList[i].Priority < todoList[j].Priority
		})
		// Note that we render the list view here, not one of the todo views
		return c.Render(http.StatusOK, "view/list-todos", nil)
	})
}

// Helper function to find a todo by ID
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

// Helper function to find a todo index by ID
func getTodoIndexByID(id int) int {
	for i, todo := range todoList {
		if todo.ID == id {
			return i
		}
	}

	return -1
}
