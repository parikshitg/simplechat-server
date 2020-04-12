package main

import (
	"github.com/gin-gonic/gin"
	"github.com/parikshitg/simplechat-server/server"
)

func main() {
	r := gin.Default()

	r.GET("/ping", server.PingHandler)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
