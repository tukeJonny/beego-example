package models

import "time"

// ModelBase is abstract base model of all models
type ModelBase struct {
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
}
