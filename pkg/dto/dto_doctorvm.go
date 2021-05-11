package dto

import "time"

type Doctorvm struct {
	Id          int64      `bson:"id" json:"id" xml:"id"`
	Createdat   *time.Time `bson:"createdat" json:"createdat" xml:"createdat"`
	Updatedat   *time.Time `bson:"updatedat" json:"updatedat" xml:"updatedat"`
	Deletedat   *time.Time `bson:"deletedat" json:"deletedat" xml:"deletedat"`
	Name        string     `bson:"name" json:"name" xml:"name"`
	Age         int64      `bson:"age" json:"age" xml:"age"`
	Nic         string     `bson:"nic" json:"nic" xml:"nic"`
	Address     string     `bson:"address" json:"address" xml:"address"`
	Email       string     `bson:"email" json:"email" xml:"email"`
	Contact     string     `bson:"contact" json:"contact" xml:"contact"`
	Nationality string     `bson:"nationality" json:"nationality" xml:"nationality"`
}
