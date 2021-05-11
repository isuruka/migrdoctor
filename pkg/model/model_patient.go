package model

import (
    "gorm.io/gorm"
)

type Patient struct {
	Model
	Name        string `bson:"name" json:"name" xml:"name"`
	Age         int64  `bson:"age" json:"age" xml:"age"`
	Nic         string `bson:"nic" json:"nic" xml:"nic"`
	Address     string `bson:"address" json:"address" xml:"address"`
	Email       string `bson:"email" json:"email" xml:"email"`
	Contact     string `bson:"contact" json:"contact" xml:"contact"`
	Nationality string `bson:"nationality" json:"nationality" xml:"nationality"`
	Doctor      string `bson:"doctor" json:"doctor" xml:"doctor"`
}

func (Patient) TableName() string {
	return "patient"
}
func (m *Patient) PreloadPatient(db *gorm.DB) *gorm.DB {
	return db
}
