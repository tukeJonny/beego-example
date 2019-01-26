package models

import "github.com/astaxie/beego/orm"

// getDefaultTodoQuery returns a default query for many todos
func getDefaultTodoQuery() orm.QuerySeter {
	var (
		db = orm.NewOrm()
		qs = db.QueryTable(todoTable)
	)

	return qs.OrderBy("id")
}
