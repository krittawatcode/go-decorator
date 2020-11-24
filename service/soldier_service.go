package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-decorator/model"
)

type SoldierService interface {
	GetPromote(c *gin.Context)
}

type soldierService struct {
	soldier model.Soldier
}

func SoldierServiceHandler(c *gin.Context) SoldierService {
	var soldier model.Soldier
	if err := c.ShouldBind(&soldier); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	return &soldierService{
		model.Soldier{
			Rank:       soldier.Rank,
			Wife:       soldier.Wife,
			Salary:     soldier.Salary,
			Home:       soldier.Home,
			Car:        soldier.Car,
			Corruption: soldier.Corruption,
		},
	}
}

func (s *soldierService) GetPromote(c *gin.Context) {
	s.soldier.Rank = c.Param("rank")
	c.JSON(http.StatusOK, gin.H{"resp": s.soldier})
}
