package handler

import (
	"crm_admin/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) getPlots(c *gin.Context) {
	plots, err := h.services.Plots.GetPlots()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string][]domain.Plot{
		"data": plots,
	})
}

func (h *handler) addPlot(c *gin.Context) {
	var plot domain.Plot
	if err := c.Bind(&plot); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Plots.CreatePlot(plot)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *handler) editPlot(c *gin.Context) {
	var plot domain.Plot
	if err := c.Bind(&plot); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Plots.EditPlot(plot)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
