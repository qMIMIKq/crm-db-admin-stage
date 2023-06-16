package view_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func newRedirecter(c *gin.Context, err error) {
	log.Info().Caller().Err(err).Msg("Error")
	c.Redirect(http.StatusFound, "/users/view")
}
