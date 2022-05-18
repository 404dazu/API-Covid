package admin

import (
	"go-covid/business/admin"
	adminBusiness "go-covid/business/admin"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service adminBusiness.Service
}

func NewController(service adminBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Get All admin
// @description Get all admin with data
// @tags AdminJWT
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} []admin.Admin
// @Router /admin [get]
func (Controller *Controller) GetAdmins(c echo.Context) error {
	admins, err := Controller.service.GetAdmins()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    admins,
	})
}

// Create godoc
// @Summary Get Admin By ID
// @description Get Admin By ID
// @tags admin
// @Accept json
// @Produce json
// @Param id path int true "anything id"
// @Success 200 {object} admin.Admin
// @Router /admin/{id} [get]
func (Controller *Controller) GetAdminByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	admin, err := Controller.service.GetAdminByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    admin,
	})
}

// Create godoc
// @Summary Create admin
// @description create admin with data
// @tags admin
// @Accept json
// @Produce json
// @Param admin body admin.AdminSwagger true "admin"
// @Success 201 {object} admin.Admin
// @Failure 400 {object} map[string]interface{}
// @Router /admin [post]
func (Controller *Controller) CreateAdmin(c echo.Context) error {
	admin := admin.Admin{}
	c.Bind(&admin)
	admins, err := Controller.service.CreateAdmin(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"Error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create admin",
		"Data":    admins,
	})
}

// Create godoc
// @Summary Get Token
// @description Get token for admin
// @tags admin
// @Accept json
// @Produce json
// @Param admin body admin.InputAdminToken true "admin"
// @Success 200 {object} map[string]interface{}
// @Failure 400
// @Router /admin/token [post]
func (Controller *Controller) GetToken(c echo.Context) error {
	var request adminBusiness.Admin

	c.Bind(&request)
	token, err := Controller.service.GetToken(&request)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get token",
		"token":   token,
	})
}

// Create godoc
// @Summary Delete Admin
// @description delete data admin
// @tags admin
// @Accept json
// @Produce json
// @Param id path int true "id admin"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Failure 400
// @Router /admin/{id} [delete]
func (Controller *Controller) DeleteAdmin(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := Controller.service.DeleteAdmin(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"messages": "success delete admin",
	})
}

// Create godoc
// @Summary Update Admin
// @description update data admin
// @tags AdminJWT
// @Accept json
// @Produce json
// @Param id path int true "id admin"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param admin body admin.AdminSwagger true "admin"
// @Success 200 {object} map[string]interface{}
// @Failure 400
// @Router /admin/{id} [PUT]
func (Controller *Controller) UpdateAdmin(c echo.Context) error {
	var admin *admin.Admin
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&admin)
	admin, err := Controller.service.UpdateAdmin(id, admin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update admin",
		"Data":    admin,
	})
}
