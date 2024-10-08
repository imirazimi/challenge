package manager

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	g := e.Group("/manager")
	g.POST("/laptops", h.IntegrateLaptop)
	g.GET("/laptops", h.GetAllLaptop)
}
