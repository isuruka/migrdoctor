package api_echo

import (
	model "migrdoctor/pkg/model"
	arfcreatepatient "migrdoctor/pkg/rule/ARF_CreatePatient"
	"reflect"
)

import (
	"github.com/labstack/echo/v4"
)

// @Summary CreatePatient
// @Tags root
// @Accept json
// @Produce json
// @Param body-param body model.Patient true  "model.Patient body data"
// @Success 200 {object} model.Patient
// @Router /api/patient [post]
func CreatePatient(c echo.Context) (*model.Patient, error) {

	patient := model.Patient{}
	if error := c.Bind(&patient); error != nil {
		return nil, error
	}
	result := arfcreatepatient.ARF_CreatePatient(&patient)
	if reflect.TypeOf(result) == reflect.TypeOf(model.Patient{}) || reflect.TypeOf(result) == reflect.TypeOf(&model.Patient{}) {
		return result.(*model.Patient), nil
	} else {
		return nil, nil
	}
}
