package admin_test

import (
	"errors"
	"fmt"
	"go-covid/business/admin"
	"os"
	"reflect"
	"testing"
)

var service admin.Service
var admin1, admin2, admin3, UpdateAdmin admin.Admin

var insertSpec, updateSpec, failedSpec, errorspec admin.Admin
var errorFindID int

// var creator, updater string
var errorInsert error = errors.New("error on insert")
var errorFind error = errors.New("error on find")

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetAdminByID(t *testing.T) {
	t.Run("Expect found the content", func(t *testing.T) {
		foundContent, _ := service.GetAdminByID(admin1.ID)
		if !reflect.DeepEqual(*foundContent, admin1) {
			t.Error("Expect content has to be equal with content1", foundContent, admin1)
		}
	})

	t.Run("Expect not found the content", func(t *testing.T) {
		content, err := service.GetAdminByID(90)

		if err != nil {
			t.Error("Expect error is nil. Error: ", err)
		} else if content != nil {
			t.Error("Expect content must be not found (nil)")
		}
	})
}

func TestGetAdminAll(t *testing.T) {
	t.Run("Expect found the admin", func(t *testing.T) {
		admins, _ := service.GetAdmins()

		if len(admins) != 3 {
			t.Error("Expect admin length must be 3")
			t.FailNow()
		}

		if reflect.DeepEqual(admins[0], admin1) {
			if !reflect.DeepEqual(admins[1], admin2) {
				t.Error("Expect 2nd admin is equal to admin2")
			}
		} else if reflect.DeepEqual(admins[0], admin2) {
			if !reflect.DeepEqual(admins[1], admin1) {
				t.Error("Expect 2nd admin is equal to admin")
			}
		} else {
			t.Error("Expect admin is admin1 and admin2")
		}
	})
}

func TestDeleteAdmin(t *testing.T) {
	t.Run("Expect delete admin2", func(t *testing.T) {
		err := service.DeleteAdmin(admin2.ID)
		if err == fmt.Errorf("") {
			t.Error("error delete")
		}
		admin, _ := service.GetAdmins()
		if len(admin) != 2 {
			t.Error("Error")
		}
	})
}
func TestInsertAdmin(t *testing.T) {
	t.Run("Expect Insert the admin", func(t *testing.T) {
		createadmins, err := service.CreateAdmin(&insertSpec)
		if err != nil {
			t.Error("Cannot insert admin")
		}
		if createadmins.ID != 3 {
			t.Error("Expect 2nd admin is equal to admin2")
		}
		NewAdmins, _ := service.GetAdminByID(createadmins.ID)
		if NewAdmins == nil {
			t.Error("Expect admins is not nil after inserted")
			t.FailNow()
		}
		GetAllAdmin, _ := service.GetAdmins()
		if reflect.DeepEqual(GetAllAdmin[2], createadmins) {
			if !reflect.DeepEqual(GetAllAdmin[0], admin1) {
				t.Error("Expect 3rd admin is equal to insertadmin")
			}
		}
		if reflect.DeepEqual(GetAllAdmin[2].ID, createadmins.ID) {
			if !reflect.DeepEqual(GetAllAdmin[0], admin1) {
				t.Error("Expect 1st admin is equal to admin1")
			}
		}
	})
}

func TestUpdateAdmin(t *testing.T) {
	t.Run("Expect Insert the admin", func(t *testing.T) {
		fmt.Println(updateSpec)
		admin, err := service.UpdateAdmin(updateSpec.ID, &updateSpec)
		fmt.Println(admin)
		if err != nil {
			t.Error("Error Update")
		}
		if admin.Name != updateSpec.Name {
			t.Error("Name tidak update")
		}
		if admin.Username != updateSpec.Username {
			t.Error("Username tidak update")
		}
		if admin.Email != updateSpec.Email {
			t.Error("Email tidak update")
		}
		if admin.Password != updateSpec.Password {
			t.Error("Password tidak update")
		}
	})
}

func setup() {
	//initialize admin 1
	admin1.ID = 1
	admin1.Email = "testemail@gmail.com"
	admin1.Username = "testusername"
	admin1.Name = "testname"
	admin1.Password = "testpassword"

	//initialize admin 2
	admin2.ID = 2
	admin2.Email = "testemail2@gmail.com"
	admin2.Username = "testusername2"
	admin2.Name = "testname2"
	admin2.Password = "testpassword2"

	//initialize admin 3
	admin3.ID = 3
	admin3.Email = "testemail3@gmail.com"
	admin3.Username = "testusername3"
	admin3.Name = "testname3"
	admin3.Password = "testpassword3"

	repo := newInMemoryRepository()
	service = admin.NewService(&repo)

	insertSpec.ID = 3
	insertSpec.Name = "InsertName"
	insertSpec.Email = "insertemail@gmail.com"
	insertSpec.Username = "InsertUsername"
	insertSpec.Password = "InsertPassword"

	updateSpec.ID = 1
	updateSpec.Name = "updatename"
	updateSpec.Email = "updateemail@gmail.com"
	updateSpec.Username = "UpdateUsername"
	updateSpec.Password = "updatePassword"

	// failedSpec.Name = ""
	// failedSpec.Description = "Failed Description"
	// failedSpec.Tags = []string{}

	// errorSpec.Name = "Error Content"
	// errorSpec.Description = "Error Description"
	// errorSpec.Tags = []string{}

	// creator = "creator"
	// updater = "updater"

	errorFindID = 3235
}

type inMemoryRepository struct {
	adminByID map[int]admin.Admin
	AllAdmin  []admin.Admin
}

func newInMemoryRepository() inMemoryRepository {
	var repo inMemoryRepository
	repo.adminByID = make(map[int]admin.Admin)
	repo.adminByID[admin1.ID] = admin1
	repo.adminByID[admin2.ID] = admin2
	repo.adminByID[admin3.ID] = admin3

	repo.AllAdmin = []admin.Admin{}
	repo.AllAdmin = append(repo.AllAdmin, admin1)
	repo.AllAdmin = append(repo.AllAdmin, admin2)
	repo.AllAdmin = append(repo.AllAdmin, admin3)

	return repo
}

func (repo *inMemoryRepository) FindAdminByID(id int) (*admin.Admin, error) {
	if id == errorFindID {
		return nil, errorFind
	}

	content, ok := repo.adminByID[id]
	if !ok {
		return nil, nil
	}

	return &content, nil
}

func (repo *inMemoryRepository) FindAdmins() (admins []admin.Admin, err error) {
	admins = repo.AllAdmin
	return admins, err
}

func (repo *inMemoryRepository) InsertAdmin(Admins *admin.Admin) (*admin.Admin, error) {
	if Admins.Name == errorspec.Name {
		return nil, errorInsert
	}
	repo.AllAdmin = append(repo.AllAdmin, *Admins)
	repo.adminByID[Admins.ID] = *Admins

	return Admins, nil
}

func (repo *inMemoryRepository) RemoveAdmin(id int) error {
	id = id - 1
	repo.AllAdmin = append(repo.AllAdmin[:id], repo.AllAdmin[id+1:]...)
	return fmt.Errorf("")
}

func (repo *inMemoryRepository) RenewAdmin(id int, admin *admin.Admin) (*admin.Admin, error) {
	admins, ok := repo.adminByID[id]
	fmt.Println(admins)
	if !ok {
		return nil, nil
	}
	admins.Email = admin.Email
	admins.Username = admin.Username
	admins.Name = admin.Name
	admins.Password = admin.Password

	return &admins, nil
}

func (repo *inMemoryRepository) CreateToken(Admins *admin.Admin) (string, error) {
	return "", fmt.Errorf("")
}
