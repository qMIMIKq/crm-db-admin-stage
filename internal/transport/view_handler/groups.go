package view_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *viewHandler) groupsView(c *gin.Context) {
	c.HTML(http.StatusOK, "groups.html", gin.H{
		"Title": "Группы пользователей",
		"CSS":   "groups/groups",
		"JS":    "groups/groups",
	})
}

func (h *viewHandler) groupsAddView(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "В разработке",
	})
}

func (h *viewHandler) groupsEditView(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		newRedirecter(c, err)
		return
	}

	group, err := h.services.Groups.GetGroupByID(idInt)
	if err != nil {
		newRedirecter(c, err)
		return
	}

	c.HTML(http.StatusOK, "groups-edit.html", gin.H{
		"Title": "Изменить группу",
		"CSS":   "users/users-edit",
		"JS":    "groups/groups-edit",
		"Group": group,
	})
}
