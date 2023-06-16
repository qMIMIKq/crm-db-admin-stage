package view_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *viewHandler) panelView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Панель администратора",
		"CSS":   "base/index",
		"JS":    "panel/index",
	})
}
