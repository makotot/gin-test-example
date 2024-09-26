package main

import "github.com/gin-gonic/gin"

type User struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
