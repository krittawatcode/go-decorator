package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-decorator/controller"
	"github.com/krittawatcode/go-decorator/service"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/v1")
	{
		v1.POST("/promote/:rank", func(c *gin.Context) {
			var soldierService service.SoldierService = service.SoldierServiceHandler(c)
			var soldierController controller.SoldierController = controller.SoldierControllerHandler(&soldierService)
			soldierController.Promote(c)
		})
	}
	port := "8080"
	server.Run(":" + port)
}
