package models

import "github.com/astaxie/beego/orm"

const todoTable = "todos"

type Todo struct {
	*ModelBase
	ID          int64  `orm:"column(id);pk"`
	Title       string `orm:"column(title);size(256)"`
	Description string `orm:"column(description);size(256)"`
}

func init() {
	orm.RegisterModel(new(Todo))
}

// TableName returns Todo's tablename
func (t *Todo) TableName() string {
	return todoTable
}

// AddTodo adds a todo
func AddTodo(todo *Todo) (int64, error) {
	var (
		db = orm.NewOrm()
	)

	id, err := db.Insert(todo)
	if err != nil {
		return -1, err
	}

	return id, nil
}

// RetrieveTodo gets a todo
func RetrieveTodo(id int64) (*Todo, error) {
	var (
		db   = orm.NewOrm()
		todo = &Todo{ID: id}
	)

	if err := db.Read(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// ListTodo gets all todos
func ListTodo() ([]*Todo, error) {
	var (
		qs    = getDefaultTodoQuery()
		todos []*Todo
	)

	_, err := qs.All(&todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// SearchTodosByTitle searches todos by `title` column
func SearchTodosByTitle(title string) ([]*Todo, error) {
	var (
		qs    = getDefaultTodoQuery()
		todos []*Todo
	)

	_, err := qs.Filter("title", title).All(&todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// UpdateTodo updates a todo
func UpdateTodo(id int64, title, desc string) error {
	var (
		qs = getDefaultTodoQuery()
	)

	_, err := qs.Filter("id", id).Update(orm.Params{
		"title":       title,
		"description": desc,
	})
	if err != nil {
		return err
	}

	return nil
}

// DeleteTodo deletes a todo
func DeleteTodo(id int64) error {
	var (
		qs = getDefaultTodoQuery()
	)

	_, err := qs.Filter("id", id).Delete()
	if err != nil {
		return err
	}

	return nil
}
