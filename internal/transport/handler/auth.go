package handler

import (
	"crm_admin/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *handler) signIn(c *gin.Context) {
	var user domain.UserAuth
	if err := c.Bind(&user); err != nil {
		log.Warn().Err(err).Msg("error binding")
		return
	}

	log.Info().Interface("user", user).Msg("user")
}
