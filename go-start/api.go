package main

import (
	"go/api/go-start/infrastructure"
	"go/api/go-start/router"

	"github.com/gin-gonic/gin"
)

func init() {
	infrastructure.ConnectToDB()
}
func main() {

	r := gin.Default()
	router.DefineRoute(r)
	r.Run("localhost:8080")
}
