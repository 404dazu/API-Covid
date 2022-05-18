package api

import (
	"go-covid/api/admin"
	"go-covid/api/covid"
	auth "go-covid/api/middleware"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminControlller *admin.Controller
	CovidController  *covid.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	//covid
	k := e.Group("/covid")
	k.GET("", controller.CovidController.GetPasien, auth.SetupAuthenticationJWT())
	k.GET("/:id", controller.CovidController.GetPasienById)
	k.POST("", controller.CovidController.CreatePasien)
	k.PUT("/:id", controller.CovidController.UpdatePasien, auth.SetupAuthenticationJWT())
	k.DELETE("/:id", controller.CovidController.DeletePasien, auth.SetupAuthenticationJWT())
	//rumah sakit
	h := e.Group("/hospital")
	h.GET("", controller.CovidController.GetRumahSakit, auth.SetupAuthenticationJWT())
	h.GET("/:id", controller.CovidController.GetRumahSakitById)
	h.POST("", controller.CovidController.CreateRumahSakit, auth.SetupAuthenticationJWT())
	h.PUT("/:id", controller.CovidController.UpdateRumahSakit, auth.SetupAuthenticationJWT())
	h.DELETE("/:id", controller.CovidController.DeleteRumahSakit, auth.SetupAuthenticationJWT())
	//admin
	g := e.Group("/admin")
	g.POST("", controller.AdminControlller.CreateAdmin)
	g.POST("/token", controller.AdminControlller.GetToken)
	g.PUT("/:id", controller.AdminControlller.UpdateAdmin)
	g.GET("/:id", controller.AdminControlller.GetAdminByID)
	// admin using jwt
	g.DELETE("/:id", controller.AdminControlller.DeleteAdmin, auth.SetupAuthenticationJWT())
	g.GET("", controller.AdminControlller.GetAdmins, auth.SetupAuthenticationJWT())
}
