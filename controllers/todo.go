package controllers

import "github.com/astaxie/beego"

// NOTE: For `@Param`, please refer https://beego.me/docs/advantage/docs.md

type TodoController struct {
	beego.Controller
}

// @Title Create
// @Description create a todo
// @Param title
