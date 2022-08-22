package models

type Player struct {
	Id   int    `form:"id" binding:"required"`
	Name string `form:"name" binding:"required"`
	Age  int8   `form:"age" binding:"required"`
}
