package view_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *viewHandler) clientsView(c *gin.Context) {
	c.HTML(http.StatusOK, "clients.html", gin.H{
		"Title": "Работа с клиентами",
		"CSS":   "users/users",
		"JS":    "clients/clients",
	})
}

func (h *viewHandler) clientsAddView(c *gin.Context) {
	c.HTML(http.StatusOK, "clients-add.html", gin.H{
		"Title": "Добавление клиента",
		"CSS":   "users/users-edit",
		"JS":    "clients/clients-add",
	})
}

func (h *viewHandler) clientsEditView(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		newRedirecter(c, err)
		return
	}

	client, err := h.services.Clients.GetClientByID(idInt)
	if err != nil {
		newRedirecter(c, err)
		return
	}

	c.HTML(http.StatusOK, "clients-edit.html", gin.H{
		"Title":  "Изменение клиента",
		"CSS":    "users/users-edit",
		"JS":     "clients/clients-edit",
		"Client": client,
	})
}
