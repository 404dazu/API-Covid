package covid

import (
	"go-covid/business/covid"
	"go-covid/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) covid.Repository {
	covidRepo := NewMysqlRepository(dbCon.Mysql)
	return covidRepo
}
