package handler

import (
	"crm_admin/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) getGroups(c *gin.Context) {
	groups, err := h.services.Groups.GetGroups()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string][]domain.Group{
		"data": groups,
	})
}

func (h *handler) editGroup(c *gin.Context) {
	var group domain.Group
	if err := c.Bind(&group); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Groups.EditGroup(group)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
