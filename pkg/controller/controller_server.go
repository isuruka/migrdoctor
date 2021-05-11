package controller

import api "migrdoctor/pkg/api"

import (
	"context"
	"migrdoctor/pub"
)

type Server struct{}

func (s *Server) CreateDoctor(ctx context.Context, parameters *pub.CreateDoctorRequest) (*pub.Doctor, error) {

	result, err := api.CreateDoctor(parameters.Doctor)
	return result, err
}

func (s *Server) CreatePatient(ctx context.Context, parameters *pub.CreatePatientRequest) (*pub.Patient, error) {

	result, err := api.CreatePatient(parameters.Patient)
	return result, err
}
