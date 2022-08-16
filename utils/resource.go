package utils

import "github.com/gin-gonic/gin"

func CreateResource(r *gin.Engine, endpoint string) *ResourceBuilder {
	instance := ResourceBuilder{r.Group(endpoint), make(map[string]func(), 5)}

	return &instance
}

type ResourceBuilder struct {
	router *gin.RouterGroup
	routes map[string]func()
}

func (resource ResourceBuilder) Build() *gin.RouterGroup {
	return resource.router
}

func (resource ResourceBuilder) AddIndexRoute(function func(*gin.Context)) *ResourceBuilder {
	resource.router.GET("/", function)

	return &resource
}

func (resource ResourceBuilder) AddStoreRoute(function func(*gin.Context)) *ResourceBuilder {
	resource.router.POST("/", function)

	return &resource
}

func (resource ResourceBuilder) AddUpdateRoute(function func(*gin.Context)) *ResourceBuilder {
	resource.router.PUT("/:index", function)

	return &resource
}

func (resource ResourceBuilder) AddDestroyRoute(function func(*gin.Context)) *ResourceBuilder {
	resource.router.DELETE("/:index", function)

	return &resource
}

func (resource ResourceBuilder) AddShowRoute(function func(*gin.Context)) *ResourceBuilder {
	resource.router.GET("/:index", function)

	return &resource
}
