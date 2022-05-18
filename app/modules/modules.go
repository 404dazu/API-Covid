package modules

import (
	"go-covid/api"
	adminApi "go-covid/api/admin"
	covidApi "go-covid/api/covid"
	adminBusiness "go-covid/business/admin"
	covidBusiness "go-covid/business/covid"
	"go-covid/config"
	adminRepo "go-covid/repository/admin"
	covidRepo "go-covid/repository/covid"
	"go-covid/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {
	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)

	covidPermitRepository := covidRepo.RepositoryFactory(dbCon)
	covidPermitService := covidBusiness.NewService(covidPermitRepository)
	covidPermitController := covidApi.NewController(covidPermitService)

	controller := api.Controller{
		AdminControlller: adminPermitController,
		CovidController:  covidPermitController,
	}
	return controller
}
