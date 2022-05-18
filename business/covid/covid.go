package covid

import (
	"time"

	"gorm.io/gorm"
)

type Rumah_sakit struct {
	ID               uint `gorm:"primarykey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Nama_rumah_sakit string         `json:"nama_rumah_sakit"`
	Alamat           string         `json:"alamat"`
}

type Res_Rumah_sakit struct {
	ID               uint `gorm:"primarykey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
	Nama_rumah_sakit string `json:"nama_rumah_sakit"`
	Alamat           string `json:"alamat"`
}

type Detail_covid struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	Status_covid   bool        `json:"covid"`
	Nama           string      `json:"nama"`
	Alamat         string      `json:"alamat"`
	Rumah_sakit_id int         `json:"rumah_sakit_id"`
	Rumah_sakit    Rumah_sakit `json:"rumah_sakit" gorm:"foreignkey:ID;references:Rumah_sakit_id"`
}
type Res_Detail_covid struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	Status_covid   bool            `json:"covid"`
	Nama           string          `json:"nama"`
	Alamat         string          `json:"alamat"`
	Rumah_sakit_id int             `json:"rumah_sakit_id"`
	Rumah_sakit    Res_Rumah_sakit `json:"rumah_sakit" gorm:"foreignkey:ID;references:Rumah_sakit_id"`
}

type Input_Detail_covid struct {
	Status_covid   bool   `json:"covid"`
	Nama           string `json:"nama"`
	Alamat         string `json:"alamat"`
	Rumah_sakit_id int    `json:"rumah_sakit_id"`
}

type Input_Rumah_sakit struct {
	Nama_rumah_sakit string `json:"nama_rumah_sakit"`
	Alamat           string `json:"alamat"`
}
