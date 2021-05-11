package api

import (
	arfcreatepatient "migrdoctor/pkg/rule/ARF_CreatePatient"
	"reflect"
)

import (
	"github.com/jinzhu/copier"
	"migrdoctor/pkg/model"
	"migrdoctor/pub"
)

func CreatePatient(patient *pub.Patient) (*pub.Patient, error) {

	mPatient := model.Patient{}
	copier.Copy(&mPatient, &patient)
	result := arfcreatepatient.ARF_CreatePatient(&mPatient)
	if reflect.TypeOf(result) == reflect.TypeOf(model.Patient{}) || reflect.TypeOf(result) == reflect.TypeOf(&model.Patient{}) {
		var data *pub.Patient
		copier.Copy(&data, &result)
		return data, nil
	} else {
		var data *pub.Patient
		return data, nil
	}
}
