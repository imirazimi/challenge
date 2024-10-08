package manager

import (
	"challenge/service/manager"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Handler struct {
	managerSvc manager.Service
}

func New(managerSvc manager.Service) Handler {
	return Handler{managerSvc: managerSvc}
}

// Add godoc
// @Summary      Integrate Laptop
// @Tags         manager Laptops
// @Accept       json
// @Produce      json
// @Param        Request body   manager.IntegrateLaptopRequest true "Integrate Laptop"
// @Success      200  {object}  manager.IntegrateLaptopResponse
// @Failure      400  {string}  "Bad request"
// @Failure      500  {string}  "Internal Server Error"
// @Router       /manager/laptops [post].
func (h Handler) IntegrateLaptop(c echo.Context) error {
	var req manager.IntegrateLaptopRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	resp, err := h.managerSvc.IntegrateLaptop(c.Request().Context(), req)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}

// Add godoc
// @Summary      GetAll Laptop
// @Tags         manager Laptops
// @Accept       json
// @Produce      json
// @Success      200  {object}  manager.GetAllLaptopResponse
// @Failure      400  {string}  "Bad request"
// @Failure      500  {string}  "Internal Server Error"
// @Router       /manager/laptops [get].
func (h Handler) GetAllLaptop(c echo.Context) error {
	var req manager.GetAllLaptopRequest
	resp, err := h.managerSvc.GetAllLaptop(c.Request().Context(), req)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}
