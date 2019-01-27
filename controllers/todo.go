package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/tukejonny/beego-example/models"
)

// NOTE: For `@Param`, please refer https://beego.me/docs/advantage/docs.md

type TodoController struct {
	beego.Controller
}

// @Title Create
// @Description create a todo
// @Param   body body models.Todo true "created todo"
// @Success 200 {int} models.Todo.Id
// @Failure 403 body is empty
// @router / [post]
func (t *TodoController) Create() {
	var todo *models.Todo

	if err := json.Unmarshal(t.Ctx.Input.RequestBody, &todo); err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	id, err := models.AddTodo(todo)
	if err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	t.Data["json"] = map[string]int64{
		"ID": id,
	}
	t.ServeJSON()
}

// @Title List
// @Description list all todos
// @Success 200 {object} models.Todo
// @router / [get]
func (t *TodoController) List() {
	todos, err := models.ListTodo()
	if err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	t.Data["json"] = todos
	t.ServeJSON()
}

// @Title Retrieve
// @Description retrieve a todo
// @Success 200 {object} models.Todo
// @Failure 403 :todoId is empty
// @router /:todoId [get]
func (t *TodoController) Retrieve() {
	_todoId := t.Ctx.Input.Param(":todoId")
	if _todoId == "" {
		t.Data["json"] = "todoId is empty"
		t.ServeJSON()
	}

	todoId, err := strconv.ParseInt(_todoId, 10, 64)
	if err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	todo, err := models.RetrieveTodo(todoId)
	if err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	t.Data["json"] = todo
	t.ServeJSON()
}

// @Title SearchByTitle
// @Description search todos by title
// @Param title query string true "todo title"
// @Success 200 {object} models.Todo
// @Failure 403 title is empty
// @router / [get]
func (t *TodoController) SearchByTitle(title *string) {
	if title == nil || *title == "" {
		t.Data["json"] = "title is Empty"
		t.ServeJSON()
	}

	todos, err := models.SearchTodosByTitle(*title)
	if err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	t.Data["json"] = todos
	t.ServeJSON()
}

// @Title Update
// @Description update a todo
// @Param todoId path string true "The todoID you want to update"
// @Param body body models.Todo true "todo"
// @Success 200 {object} models.Todo
// @Failure 403 :todoId is empty
// @router /:todoId [put]
func (t *TodoController) Update() {
	var (
		_todoId = t.Ctx.Input.Param(":todoId")
		todo    *models.Todo
	)

	// id(string) -> id(int64)
	todoId, err := strconv.ParseInt(_todoId, 10, 64)
	if err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	if err := json.Unmarshal(t.Ctx.Input.RequestBody, todo); err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	if err := models.UpdateTodo(todoId, todo.Title, todo.Description); err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	t.Data["json"] = "update success"
	t.ServeJSON()
}

// @Title Delete
// @Description delete a todo
// @Param todoId path string true "The todoId you want to delete"
// @Success 200 {string} delete success
// @Failure 403 todoId is empty
// @router /:todoId [delete]
func (t *TodoController) Delete() {
	_todoId := t.Ctx.Input.Param(":todoId")

	todoId, err := strconv.ParseInt(_todoId, 10, 64)
	if err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	if err := models.DeleteTodo(todoId); err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
	}

	t.Data["json"] = "delete success"
	t.ServeJSON()
}
