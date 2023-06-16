package view_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *viewHandler) plotsView(c *gin.Context) {
	c.HTML(http.StatusOK, "plots.html", gin.H{
		"Title": "Участки",
		"CSS":   "plots/plots",
		"JS":    "plots/plots",
	})
}

func (h *viewHandler) plotsAddView(c *gin.Context) {
	c.HTML(http.StatusOK, "plots-add.html", gin.H{
		"Title": "Добавить участок",
		"CSS":   "users/users-edit",
		"JS":    "plots/plots-add",
	})
}

func (h *viewHandler) plotsEditView(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/users/view")
		return
	}

	plot, err := h.services.Plots.GetPlotByID(idInt)
	if err != nil {
		c.Redirect(http.StatusFound, "/users/view")
		return
	}

	c.HTML(http.StatusOK, "plots-edit.html", gin.H{
		"Title": "Изменить участок",
		"CSS":   "users/users-edit",
		"JS":    "plots/plots-edit",
		"Plot":  plot,
	})
}
