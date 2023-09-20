package main

import (
	"examples/go-crud/controlers"
	"examples/go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controlers.PostsCreate)
	r.GET("/posts", controlers.PostIndex)
	r.GET("/posts/:id", controlers.PostSingle)
	r.PUT("/posts/:id", controlers.PostUpdate)
	r.DELETE("/posts/:id", controlers.PostDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
