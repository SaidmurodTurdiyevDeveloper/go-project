package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"firstname,omitempty"`
	SurName string             `json:"lastname,omitempty"`
	Age     int                `json:"age,omitempty"`
}

func (student *Student) IsValid() bool {
	return student.Name != "" && student.Age > 0 && student.SurName != ""
}
