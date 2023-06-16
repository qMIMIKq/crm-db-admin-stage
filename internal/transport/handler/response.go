package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type err struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, status int, msg string) {
	log.Warn().Interface("error:", msg).Caller().Msg("Error is")
	c.AbortWithStatusJSON(status, err{msg})
}
