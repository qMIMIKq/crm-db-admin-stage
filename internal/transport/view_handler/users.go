package view_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *viewHandler) usersView(c *gin.Context) {
	c.HTML(http.StatusOK, "users.html", gin.H{
		"Title": "Работа с пользователями",
		"CSS":   "users/users",
		"JS":    "users/users",
	})
}

func (h *viewHandler) usersAddView(c *gin.Context) {
	c.HTML(http.StatusOK, "users-add.html", gin.H{
		"Title": "Добавление пользователей",
		"CSS":   "users/users-edit",
		"JS":    "users/users-add",
	})
}

func (h *viewHandler) editUserView(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/users/view")
		return
	}

	user, err := h.services.Users.GetUserByID(idInt)
	if err != nil {
		newRedirecter(c, err)
		return
	}

	c.HTML(http.StatusOK, "users-edit.html", gin.H{
		"Title": "Изменить пользователя",
		"CSS":   "users/users-edit",
		"JS":    "users/users-edit",
		"User":  user,
	})
}
