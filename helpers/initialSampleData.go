package helpers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hecharles/doctortest/database"
	"github.com/hecharles/doctortest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

//InitialSampleData for sample data
func InitialSampleData() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	doctorCollection, err := database.GetMongoDbCollection("doctortest", "doctor")

	if err != nil {
		log.Fatal(err)

	}

	opts := options.Update().SetUpsert(true)

	//sample doctor 1
	sampleDoctor := models.Doctor{
		ID:            "sample-doctor-1",
		Prefix:        "Dr.",
		Name:          "Mitchell Fishbach",
		Suffix:        "MD",
		Gender:        "Male",
		Age:           69,
		Email:         "mitchell_fishbach@email.com",
		Mobile:        "(256) 290-4957",
		Qualification: []string{"MD", "Cardiology"},
		Summary:       "Dr. Mitchell Fishbach, MD is a Cardiology Specialist in Scarsdale, NY.  Dr. Fishbach has more experience with Adult Congenital Cardiac Disorders and Cardiac Care than other specialists in his area.  He is affiliated with medical facilities New York-Presbyterian Lawrence Hospital and NewYork-Presbyterian/Columbia University Medical Center.  He is accepting new patients and has indicated that he accepts telehealth appointments.  Be sure to call ahead with Dr. Fishbach to book an appointment.",
	}

	filter := bson.M{"_id": sampleDoctor.ID}
	update := bson.M{"$set": bson.M{
		"Prefix":        sampleDoctor.Prefix,
		"Name":          sampleDoctor.Name,
		"Suffix":        sampleDoctor.Suffix,
		"Gender":        sampleDoctor.Gender,
		"Age":           sampleDoctor.Age,
		"Email":         sampleDoctor.Email,
		"Mobile":        sampleDoctor.Mobile,
		"Qualification": sampleDoctor.Qualification,
		"Summary":       sampleDoctor.Summary,
	},
	}

	result, err := doctorCollection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		fmt.Println("UpdateOne() result ERROR:", err)

	} else {
		fmt.Println("UpdateOne() result:", result)
		fmt.Println("UpdateOne() result MatchedCount:", result.MatchedCount)
		fmt.Println("UpdateOne() result ModifiedCount:", result.ModifiedCount)
		fmt.Println("UpdateOne() result UpsertedCount:", result.UpsertedCount)
		fmt.Println("UpdateOne() result UpsertedID:", result.UpsertedID)
	}

	//sample doctor 2
	sampleDoctor = models.Doctor{
		ID:            "sample-doctor-2",
		Prefix:        "Dr.",
		Name:          "Rachel Weller",
		Suffix:        "MD",
		Gender:        "Female",
		Age:           56,
		Email:         "rachel_weller@email.com",
		Mobile:        "(256) 293-4958",
		Qualification: []string{"Pediatic", "MD", "Cardiology"},
		Summary:       "Dr. Rachel Weller, MD is a Pediatric Cardiology Specialist in New York, NY.  She is affiliated with medical facilities NewYork-Presbyterian/Columbia University Medical Center and NewYork-Presbyterian Morgan Stanley Children's Hospital.  She is accepting new patients and has indicated that she accepts telehealth appointments.  Be sure to call ahead with Dr. Weller to book an appointment.",
	}

	filter = bson.M{"_id": sampleDoctor.ID}
	update = bson.M{"$set": bson.M{
		"Prefix":        sampleDoctor.Prefix,
		"Name":          sampleDoctor.Name,
		"Suffix":        sampleDoctor.Suffix,
		"Gender":        sampleDoctor.Gender,
		"Age":           sampleDoctor.Age,
		"Email":         sampleDoctor.Email,
		"Mobile":        sampleDoctor.Mobile,
		"Qualification": sampleDoctor.Qualification,
		"Summary":       sampleDoctor.Summary,
	},
	}

	result, err = doctorCollection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		fmt.Println("UpdateOne() result ERROR:", err)

	} else {
		fmt.Println("UpdateOne() result:", result)
		fmt.Println("UpdateOne() result MatchedCount:", result.MatchedCount)
		fmt.Println("UpdateOne() result ModifiedCount:", result.ModifiedCount)
		fmt.Println("UpdateOne() result UpsertedCount:", result.UpsertedCount)
		fmt.Println("UpdateOne() result UpsertedID:", result.UpsertedID)
	}

	//sample doctor 3
	sampleDoctor = models.Doctor{
		ID:            "sample-doctor-3",
		Prefix:        "Dr.",
		Name:          "Jeffrey Weinberg",
		Suffix:        "MD",
		Gender:        "Male",
		Age:           53,
		Email:         "jeffry_weinberg@email.com",
		Mobile:        "(256) 323-3728",
		Qualification: []string{"Dermatology"},
		Summary:       "Dr. Jeffrey Weinberg, MD is a Dermatologist in Brooklyn, NY.  Dr. Weinberg has more experience with Cosmetic Dermatology, Dermatologic Care, and Dermatologic Oncology than other specialists in his area.  He is affiliated with medical facilities such as Lenox Hill Hospital and Jamaica Hospital Medical Center.  He is accepting new patients and has indicated that he accepts telehealth appointments.  Be sure to call ahead with Dr. Weinberg to book an appointment",
	}

	filter = bson.M{"_id": sampleDoctor.ID}
	update = bson.M{"$set": bson.M{
		"Prefix":        sampleDoctor.Prefix,
		"Name":          sampleDoctor.Name,
		"Suffix":        sampleDoctor.Suffix,
		"Gender":        sampleDoctor.Gender,
		"Age":           sampleDoctor.Age,
		"Email":         sampleDoctor.Email,
		"Mobile":        sampleDoctor.Mobile,
		"Qualification": sampleDoctor.Qualification,
		"Summary":       sampleDoctor.Summary,
	},
	}

	result, err = doctorCollection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		fmt.Println("UpdateOne() result ERROR:", err)

	} else {
		fmt.Println("UpdateOne() result:", result)
		fmt.Println("UpdateOne() result MatchedCount:", result.MatchedCount)
		fmt.Println("UpdateOne() result ModifiedCount:", result.ModifiedCount)
		fmt.Println("UpdateOne() result UpsertedCount:", result.UpsertedCount)
		fmt.Println("UpdateOne() result UpsertedID:", result.UpsertedID)
	}

	// sample customer 1

	customerCollection, err := database.GetMongoDbCollection("doctortest", "customer")

	sampleCustomer := models.Customer{
		ID:     "sample-customer-1",
		Name:   "Hendra Charles",
		Gender: "Male",
		Mobile: "+6281145784756",
		Email:  "hendra@email.com",
	}

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("sample customer hash password ERROR:", err)
		return

	}
	sampleCustomer.Password = string(hashedPassword)

	filter = bson.M{"_id": sampleCustomer.ID}
	update = bson.M{"$set": bson.M{
		"Name":     sampleCustomer.Name,
		"Gender":   sampleCustomer.Gender,
		"Mobile":   sampleCustomer.Mobile,
		"Email":    sampleCustomer.Email,
		"Password": sampleCustomer.Password,
	},
	}

	result, err = customerCollection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		fmt.Println("UpdateOne() result ERROR:", err)

	} else {
		fmt.Println("UpdateOne() result:", result)
		fmt.Println("UpdateOne() result MatchedCount:", result.MatchedCount)
		fmt.Println("UpdateOne() result ModifiedCount:", result.ModifiedCount)
		fmt.Println("UpdateOne() result UpsertedCount:", result.UpsertedCount)
		fmt.Println("UpdateOne() result UpsertedID:", result.UpsertedID)
	}
}
