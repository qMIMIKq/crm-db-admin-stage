package view_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *viewHandler) filtersView(c *gin.Context) {
	c.HTML(http.StatusOK, "filters.html", gin.H{
		"Title": "Управление фильтрами",
		"CSS":   "users/users",
		"JS":    "filters/filters",
	})
}

func (h *viewHandler) filtersAddView(c *gin.Context) {

	c.HTML(http.StatusOK, "filters-add.html", gin.H{
		"Title": "Добавлени фильтра",
		"CSS":   "users/users-edit",
		"JS":    "filters/filters-add",
	})
}

func (h *viewHandler) filtersEditView(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		newRedirecter(c, err)
		return
	}

	filter, err := h.services.Filters.GetFilterByID(idInt)
	if err != nil {
		newRedirecter(c, err)
		return
	}

	c.HTML(http.StatusOK, "filters-edit.html", gin.H{
		"Title":  "Изменение фильтра",
		"CSS":    "users/users-edit",
		"JS":     "filters/filters-edit",
		"Filter": filter,
	})
}
