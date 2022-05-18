package covid

import "github.com/go-playground/validator/v10"

type Repository interface {
	FindPasien() (detail_covids []Detail_covid, err error)
	FindPasienById(id int) (*Detail_covid, error)
	InsertPasien(detail_covid *Detail_covid) (*Detail_covid, error)
	RenewPasien(id int, detail_covid *Detail_covid) (*Detail_covid, error)
	RemovePasien(id int) (*Detail_covid, error)
	FindRumahSakit() (rumah_sakit []Rumah_sakit, err error)
	FindRumahSakitById(id int) (*Rumah_sakit, error)
	InsertRumahSakit(rumah_sakit *Rumah_sakit) (*Rumah_sakit, error)
	RenewRumahSakit(id int, rumah_sakit *Rumah_sakit) (*Rumah_sakit, error)
	RemoveRumahSakit(id int) (*Rumah_sakit, error)
}
type Service interface {
	GetPasien() (detail_covids []Detail_covid, err error)
	GetPasienById(id int) (detail_covid *Detail_covid, err error)
	CreatePasien(detail_covid *Detail_covid) (*Detail_covid, error)
	UpdatePasien(id int, detail_covid *Detail_covid) (*Detail_covid, error)
	DeletePasien(id int) (*Detail_covid, error)
	GetRumahSakit() (rumah_sakit []Rumah_sakit, err error)
	GetRumahSakitById(id int) (rumah_sakit *Rumah_sakit, err error)
	CreateRumahSakit(rumah_sakit *Rumah_sakit) (*Rumah_sakit, error)
	UpdateRumahSakit(id int, rumah_sakit *Rumah_sakit) (*Rumah_sakit, error)
	DeleteRumahSakit(id int) (*Rumah_sakit, error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) GetPasien() (detail_covids []Detail_covid, err error) {
	return s.repository.FindPasien()
}

func (s *service) GetPasienById(id int) (detail_covid *Detail_covid, err error) {
	return s.repository.FindPasienById(id)
}

func (s *service) CreatePasien(detail_covid *Detail_covid) (*Detail_covid, error) {
	return s.repository.InsertPasien(detail_covid)
}

func (s *service) UpdatePasien(id int, detail_covid *Detail_covid) (*Detail_covid, error) {
	return s.repository.RenewPasien(id, detail_covid)
}

func (s *service) DeletePasien(id int) (*Detail_covid, error) {
	return s.repository.RemovePasien(id)
}

func (s *service) GetRumahSakit() (rumah_sakit []Rumah_sakit, err error) {
	return s.repository.FindRumahSakit()
}

func (s *service) GetRumahSakitById(id int) (rumah_sakit *Rumah_sakit, err error) {
	return s.repository.FindRumahSakitById(id)
}

func (s *service) CreateRumahSakit(rumah_sakit *Rumah_sakit) (*Rumah_sakit, error) {
	return s.repository.InsertRumahSakit(rumah_sakit)
}

func (s *service) UpdateRumahSakit(id int, rumah_sakit *Rumah_sakit) (*Rumah_sakit, error) {
	return s.repository.RenewRumahSakit(id, rumah_sakit)
}

func (s *service) DeleteRumahSakit(id int) (*Rumah_sakit, error) {
	return s.repository.RemoveRumahSakit(id)
}
