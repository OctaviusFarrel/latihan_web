package models

type Player struct {
	Name string `form:"name" binding:"required"`
	Age  int8   `form:"age" binding:"required"`
}
