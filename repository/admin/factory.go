package admin

import (
	"go-covid/business/admin"
	"go-covid/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) admin.Repository {
	adminRepo := NewMysqlRepository(dbCon.Mysql)
	return adminRepo
}
