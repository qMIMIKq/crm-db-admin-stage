package view_handler

import (
	"crm_admin/internal/services"
	"github.com/gin-gonic/gin"
)

type ViewHandler interface {
	InitViewRoutes(router *gin.Engine)
}

type viewHandler struct {
	services *services.Services
}

func (h *viewHandler) InitViewRoutes(router *gin.Engine) {
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("./web/templates/**/*.html")

	router.GET("/panel", h.panelView)

	users := router.Group("/users")
	{
		users.GET("/view", h.usersView)
		users.GET("/add", h.usersAddView)
		users.GET("/edit/:id", h.editUserView)
	}

	groups := router.Group("/groups")
	{
		groups.GET("/view", h.groupsView)
		groups.GET("/add", h.groupsAddView)
		groups.GET("/edit/:id", h.groupsEditView)
	}

	plots := router.Group("/plots")
	{
		plots.GET("/view", h.plotsView)
		plots.GET("/add", h.plotsAddView)
		plots.GET("/edit/:id", h.plotsEditView)
	}

	filters := router.Group("/filters")
	{
		filters.GET("/view", h.filtersView)
		filters.GET("/add", h.filtersAddView)
		filters.GET("/edit/:id", h.filtersEditView)
	}

	clients := router.Group("/clients")
	{
		clients.GET("/view", h.clientsView)
		clients.GET("/add", h.clientsAddView)
		clients.GET("/edit/:id", h.clientsEditView)
	}

	constructor := router.Group("/constructor")
	{
		constructor.GET("/view", h.constructorView)
	}
}

func NewViewHandler(services *services.Services) ViewHandler {
	return &viewHandler{
		services: services,
	}
}
