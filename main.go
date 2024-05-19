package main

import (
	"fmt"
	"goreads/routes"
	"goreads/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello goreads")

	app := gin.Default()
	routes.RegisterRoutes(app)

	app.SetTrustedProxies(nil)
	err := app.Run(":2324")

	utils.PrintIfErr(err, "SERVER exception")
}
