package controller

import (
	"github.com/PedroMoreiraDev/projeto-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota errada")
	c.JSON(err.Code, err)
}
