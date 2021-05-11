package api_echo

import (
    "fmt"
    model "migrdoctor/pkg/model"
	arfcreatedoctor "migrdoctor/pkg/rule/ARF_CreateDoctor"
	"reflect"
)

import (
	"github.com/labstack/echo/v4"
)

// @Summary CreateDoctor
// @Tags root
// @Accept json
// @Produce json
// @Param body-param body model.Doctor true  "model.Doctor body data"
// @Success 200 {object} model.Doctor
// @Router /api/doctor [post]
func CreateDoctor(c echo.Context) (*model.Doctor, error) {

	doctor := model.Doctor{}
	if error := c.Bind(&doctor); error != nil {
		return nil, error
	}
	result := arfcreatedoctor.ARF_CreateDoctor(&doctor)
	if reflect.TypeOf(result) == reflect.TypeOf(model.Doctor{}) || reflect.TypeOf(result) == reflect.TypeOf(&model.Doctor{}) {
		return result.(*model.Doctor), nil
	} else {
		return nil, nil
	}
}

func CreateDoctor2 (doctor *model.Doctor) {


    result := arfcreatedoctor.ARF_CreateDoctor(doctor)
    if reflect.TypeOf(result) == reflect.TypeOf(model.Doctor{}) || reflect.TypeOf(result) == reflect.TypeOf(&model.Doctor{}) {
        fmt.Println("Write Completed")
        fmt.Println(doctor)
    } else {
        fmt.Println("Write Failed")
        fmt.Println(result)
    }
}
