package ARF_CreatePatient

import (
	"migrdoctor/pkg/model"
)

type CFGContext struct {
	Patient      *model.Patient
	_returnValue interface{}
	_errorValue  interface{}
}

func NewCFGContext(pPatient *model.Patient) *CFGContext {
	return &CFGContext{

		Patient: pPatient,
	}
}
func ARF_CreatePatient(pPatient *model.Patient) interface{} {

	cfg := NewCFGContext(pPatient)
	return cfg._returnValue
}
