package covid

import (
	"fmt"
	"go-covid/business/covid"

	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *MysqlRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (repo *MysqlRepository) FindPasien() (detail_covids []covid.Detail_covid, err error) {
	err = repo.db.Preload("Rumah_sakit").Find(&detail_covids).Error
	if err != nil {
		return nil, err
	}
	return detail_covids, nil
}

func (repo *MysqlRepository) FindPasienById(id int) (*covid.Detail_covid, error) {
	var covid *covid.Detail_covid
	err := repo.db.Where("ID = ?", id).Preload("Rumah_sakit").First(&covid).Error
	if err != nil {
		return nil, err
	}
	return covid, nil
}

func (repo *MysqlRepository) InsertPasien(detail_covid *covid.Detail_covid) (*covid.Detail_covid, error) {
	err := repo.db.Create(&detail_covid).Preload("Rumah_sakit").Error
	if err != nil {
		return nil, fmt.Errorf("Your data is Failed!")
	}
	return detail_covid, nil
}

func (repo *MysqlRepository) RenewPasien(id int, detail_covid *covid.Detail_covid) (*covid.Detail_covid, error) {
	err := repo.db.Model(*detail_covid).Where("ID = ?", id).Updates(detail_covid).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Where("ID = ?", id).First(detail_covid).Error
	if err != nil {
		return nil, err
	}
	return detail_covid, nil
}

func (repo *MysqlRepository) RemovePasien(id int) (*covid.Detail_covid, error) {
	var detail_covid *covid.Detail_covid
	err := repo.db.Where("ID = ?", id).First(&detail_covid).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Delete(detail_covid, id).Error
	if err != nil {
		return nil, err
	}
	return detail_covid, nil
}

func (repo *MysqlRepository) FindRumahSakit() (rumah_sakit []covid.Rumah_sakit, err error) {
	err = repo.db.Find(&rumah_sakit).Error
	if err != nil {
		return nil, err
	}
	return rumah_sakit, nil
}

func (repo *MysqlRepository) FindRumahSakitById(id int) (*covid.Rumah_sakit, error) {
	var rumah_sakit *covid.Rumah_sakit
	err := repo.db.Where("ID = ?", id).First(&rumah_sakit).Error
	if err != nil {
		return nil, err
	}
	return rumah_sakit, nil
}

func (repo *MysqlRepository) InsertRumahSakit(rumah_sakit *covid.Rumah_sakit) (*covid.Rumah_sakit, error) {
	err := repo.db.Create(&rumah_sakit).Error
	if err != nil {
		return nil, fmt.Errorf("Your data is Failed!")
	}
	return rumah_sakit, nil
}

func (repo *MysqlRepository) RenewRumahSakit(id int, rumah_sakit *covid.Rumah_sakit) (*covid.Rumah_sakit, error) {
	err := repo.db.Model(*rumah_sakit).Where("ID = ?", id).Updates(rumah_sakit).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Where("ID = ?", id).First(rumah_sakit).Error
	if err != nil {
		return nil, err
	}
	return rumah_sakit, nil
}

func (repo *MysqlRepository) RemoveRumahSakit(id int) (*covid.Rumah_sakit, error) {
	var rumah_sakit *covid.Rumah_sakit
	err := repo.db.Where("ID = ?", id).First(&rumah_sakit).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Delete(rumah_sakit, id).Error
	if err != nil {
		return nil, err
	}
	return rumah_sakit, nil
}
