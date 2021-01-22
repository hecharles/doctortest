package models

import (
	"context"
	"time"

	"github.com/hecharles/doctortest/database"
	"go.mongodb.org/mongo-driver/bson"
)

//Customer model
type Customer struct {
	ID       string `bson:"_id" json:"id,omitempty"`
	Name     string `bson:"Name" json:"Name,omitempty"`
	Gender   string `bson:"Gender" json:"Gender,omitempty"`
	Mobile   string `bson:"Mobile" json:"Mobile,omitempty"`
	Email    string `bson:"Email" json:"Email,omitempty"`
	Password string `bson:"Password" json:"Password,omitempty"`
}

//GetCustomerByEmail find customer
func GetCustomerByEmail(email string) (*Customer, error) {

	customerCollection, err := database.GetMongoDbCollection("doctortest", "customer")
	if err != nil {
		return nil, err
	}

	var customer Customer

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var filter bson.M = bson.M{"Email": email}

	err = customerCollection.FindOne(ctx, filter).Decode(&customer)

	return &customer, err

}

//CreateNewCustomer create new customer
func CreateNewCustomer(customer *Customer) error {

	customerCollection, err := database.GetMongoDbCollection("doctortest", "customer")
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err = customerCollection.InsertOne(ctx, customer)
	if err != nil {
		return err
	}
	return nil

}
