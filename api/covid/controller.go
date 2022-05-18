package covid

import (
	covidBusiness "go-covid/business/covid"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service covidBusiness.Service
}

func NewController(service covidBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Get All Data Pasien
// @description Get All Data Pasien with data
// @tags Pasien
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} []covid.Res_Detail_covid
// @Router /covid [get]
func (Controller *Controller) GetPasien(c echo.Context) error {
	pasien, err := Controller.service.GetPasien()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    pasien,
	})
}

// Create godoc
// @Summary Get Data Pasien By ID
// @description Get Data Pasien By ID
// @tags Pasien
// @Accept json
// @Produce json
// @Param id path int true "anything id"
// @Success 200 {object} covid.Res_Detail_covid
// @Router /covid/{id} [get]
func (Controller *Controller) GetPasienById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	pasien, err := Controller.service.GetPasienById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    pasien,
	})
}

// Create godoc
// @Summary Create data pasien
// @description Create data pasien
// @tags Pasien
// @Accept json
// @Produce json
// @Param pasien body covid.Input_Detail_covid true "covid"
// @Success 200 {object} covid.Res_Detail_covid
// @Router /covid [post]
func (Controller *Controller) CreatePasien(c echo.Context) error {
	pasien := covidBusiness.Detail_covid{}
	c.Bind(&pasien)
	pasiens, err := Controller.service.CreatePasien(&pasien)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"Error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    pasiens,
	})
}

// Create godoc
// @Summary Update data pasien
// @description Update data pasien
// @tags Pasien
// @Accept json
// @Produce json
// @Param id path int true "anything id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param pasien body covid.Input_Detail_covid true "covid"
// @Success 200 {object} covid.Res_Detail_covid
// @Router /covid/{id} [put]
func (Controller *Controller) UpdatePasien(c echo.Context) error {
	var pasien *covidBusiness.Detail_covid
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&pasien)
	pasien, err := Controller.service.UpdatePasien(id, pasien)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    pasien,
	})
}

// Create godoc
// @Summary Delete data pasien
// @description Delete data pasien
// @tags Pasien
// @Accept json
// @Produce json
// @Param id path int true "anything id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} covid.Res_Detail_covid
// @Router /covid/{id} [delete]
func (Controller *Controller) DeletePasien(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := Controller.service.DeletePasien(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"messages": "deleted pasien",
	})
}

// Create godoc
// @Summary Get all data rumah sakit
// @description Get all data rumah sakit with data
// @tags Pasien
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} []covid.Res_Rumah_sakit
// @Router /hospital [get]
func (Controller *Controller) GetRumahSakit(c echo.Context) error {
	rumah_sakit, err := Controller.service.GetRumahSakit()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    rumah_sakit,
	})
}

// Create godoc
// @Summary Get data rumah sakit by id
// @description Get all data rumah sakit with data
// @tags Pasien
// @Accept json
// @Produce json
// @Param id path int true "anything id"
// @Success 200 {object} covid.Res_Rumah_sakit
// @Router /hospital/{id} [get]
func (Controller *Controller) GetRumahSakitById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	rumah_sakit, err := Controller.service.GetRumahSakitById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    rumah_sakit,
	})
}

// Create godoc
// @Summary Create data rumah sakit
// @description Create data rumah sakit
// @tags Pasien
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param rumah_sakit body covid.Input_Rumah_sakit true "hospital"
// @Success 200 {object} covid.Res_Rumah_sakit
// @Router /hospital [post]
func (Controller *Controller) CreateRumahSakit(c echo.Context) error {
	rumah_sakit := covidBusiness.Rumah_sakit{}
	c.Bind(&rumah_sakit)
	rumah_sakits, err := Controller.service.CreateRumahSakit(&rumah_sakit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"Error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    rumah_sakits,
	})
}

// Create godoc
// @Summary Update data rumah sakit
// @description Update data rumah sakit
// @tags Pasien
// @Accept json
// @Produce json
// @Param id path int true "anything id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param rumah_sakit body covid.Input_Rumah_sakit true "hospital"
// @Success 200 {object} covid.Res_Rumah_sakit
// @Router /hospital/{id} [put]
func (Controller *Controller) UpdateRumahSakit(c echo.Context) error {
	var rumah_sakit *covidBusiness.Rumah_sakit
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&rumah_sakit)
	rumah_sakit, err := Controller.service.UpdateRumahSakit(id, rumah_sakit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    rumah_sakit,
	})
}

// Create godoc
// @Summary Delete data rumah sakit
// @description Delete data rumah sakit
// @tags Pasien
// @Accept json
// @Produce json
// @Param id path int true "anything id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} covid.Res_Rumah_sakit
// @Router /hospital/{id} [delete]
func (Controller *Controller) DeleteRumahSakit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	rumah_sakit, err := Controller.service.DeleteRumahSakit(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "deleted rumah sakit",
		"data":     rumah_sakit,
	})
}
