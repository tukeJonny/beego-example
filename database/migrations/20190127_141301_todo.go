package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Todo_20190127_141301 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Todo_20190127_141301{}
	m.Created = "20190127_141301"

	migration.Register("Todo_20190127_141301", m)
}

// Run the migrations
func (m *Todo_20190127_141301) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Todo_20190127_141301) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
