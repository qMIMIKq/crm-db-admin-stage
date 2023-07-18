package handler

import (
	"crm_admin/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) getFilters(c *gin.Context) {
	filters, err := h.services.Filters.GetFilters()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": filters,
	})
}

func (h *handler) addFilter(c *gin.Context) {
	var filter domain.FilterInfo
	if err := c.Bind(&filter); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Filters.CreateFilter(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *handler) editFilter(c *gin.Context) {
	var filter domain.FilterInfo
	if err := c.Bind(&filter); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Filters.EditFilter(filter); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *handler) editFilterPosition(c *gin.Context) {
	var filters []domain.FilterInfo
	if err := c.Bind(&filters); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Filters.UpdatePosition(filters); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
