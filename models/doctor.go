package models

import (
	"context"
	"time"

	"github.com/hecharles/doctortest/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Doctor model
type Doctor struct {
	ID            string   `bson:"_id" json:"id,omitempty"`
	Prefix        string   `bson:"Prefix" json:"Prefix,omitempty"`
	Name          string   `bson:"Name" json:"Name,omitempty"`
	Suffix        string   `bson:"Suffix" json:"Suffix,omitempty"`
	Gender        string   `bson:"Gender" json:"Gender,omitempty"`
	Age           int      `bson:"Age" json:"Age,omitempty"`
	Email         string   `bson:"Email" json:"Email,omitempty"`
	Mobile        string   `bson:"Mobile" json:"Mobile,omitempty"`
	Qualification []string `bson:"Qualification" json:"Qualification,omitempty"`
	Summary       string   `bson:"Summary" json:"Summary,omitempty"`
}

//GetAllDoctorPublic get simple doctor from database
func GetAllDoctorPublic() (*[]Doctor, error) {

	doctorCollection, err := database.GetMongoDbCollection("doctortest", "doctor")

	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	fields := bson.M{"_id": 1, "Prefix": 1, "Name": 1, "Suffix": 1, "Gender": 1}

	opts := options.Find().SetProjection(fields)
	filter := bson.M{}

	cursor, err := doctorCollection.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}

	doctors := []Doctor{}

	err = cursor.All(ctx, &doctors)
	if err != nil {
		return nil, err
	}

	return &doctors, err

}

//GetCustomerByEmail find customer
func GetAllDoctorCustomer() (*[]Doctor, error) {

	doctorCollection, err := database.GetMongoDbCollection("doctortest", "doctor")

	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{}

	cursor, err := doctorCollection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	doctors := []Doctor{}

	err = cursor.All(ctx, &doctors)
	if err != nil {
		return nil, err
	}

	return &doctors, err

}
