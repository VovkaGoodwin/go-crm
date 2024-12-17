package usecase

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthcheckUseCase struct {
}

func NewHealthCheckUseCase() *HealthcheckUseCase {
	return &HealthcheckUseCase{}
}

func (u *HealthcheckUseCase) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
