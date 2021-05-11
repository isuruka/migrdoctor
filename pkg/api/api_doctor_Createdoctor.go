package api

import (
	arfcreatedoctor "migrdoctor/pkg/rule/ARF_CreateDoctor"
	"reflect"
)

import (
	"github.com/jinzhu/copier"
	"migrdoctor/pkg/model"
	"migrdoctor/pub"
)

func CreateDoctor(doctor *pub.Doctor) (*pub.Doctor, error) {

	mDoctor := model.Doctor{}
	copier.Copy(&mDoctor, &doctor)
	result := arfcreatedoctor.ARF_CreateDoctor(&mDoctor)
	if reflect.TypeOf(result) == reflect.TypeOf(model.Doctor{}) || reflect.TypeOf(result) == reflect.TypeOf(&model.Doctor{}) {
		var data *pub.Doctor
		copier.Copy(&data, &result)
		return data, nil
	} else {
		var data *pub.Doctor
		return data, nil
	}
}
