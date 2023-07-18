package view_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *viewHandler) constructorView(c *gin.Context) {

	c.HTML(http.StatusOK, "constructor.html", gin.H{
		"Title": "Конструктор фильтров",
		"CSS":   "users/users",
		"JS":    "constructor/constructor",
	})
}
