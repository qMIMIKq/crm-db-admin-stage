package handler

import (
	"crm_admin/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) getClients(c *gin.Context) {
	clients, err := h.services.Clients.GetClients()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": clients,
	})
}

func (h *handler) addClient(c *gin.Context) {
	var client domain.Client
	if err := c.Bind(&client); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Clients.CreateClient(client)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *handler) editClient(c *gin.Context) {
	var client domain.Client
	if err := c.Bind(&client); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Clients.UpdateClient(client); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
