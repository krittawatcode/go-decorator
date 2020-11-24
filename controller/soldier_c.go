package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-decorator/service"
)

type SoldierController interface {
	Promote(c *gin.Context)
}

type soldierController struct {
	soldierService service.SoldierService
}

func SoldierControllerHandler(soldierService *service.SoldierService) SoldierController {
	return &soldierController{
		soldierService: *soldierService,
	}
}

func (s *soldierController) Promote(c *gin.Context) {
	s.soldierService.GetPromote(c)
}
