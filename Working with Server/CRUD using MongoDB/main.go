package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Employee struct {
	name   string
	eid    string
	office string
	ldesc  string
	lname  string
	lid    string
	nname  string
}

func main() {

	emp := Employee{
		name:   "Name",
		eid:    "primary key",
		office: "desc",
		lname:  "name associated (asset)",
		ldesc:  "desc of asset ",
		lid:    "foreign key",
		nname:  "name to be updated to",
	}
	Create(emp)
	fmt.Println("Created Successfully")
	Retrieve(emp)
	fmt.Println("Retrieved Successfully")
	Update(emp)
	fmt.Println("Updated Successfully")
	Delete(emp)
	fmt.Println("Deleted Successfully")

}

func Create(e Employee) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.ut86q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Connected to Database", databases)

	quickstartDatabase := client.Database("Employee")
	detailsCollection := quickstartDatabase.Collection("Details")
	assetsCollection := quickstartDatabase.Collection("Asset")

	detailsResult, err := detailsCollection.InsertOne(ctx, bson.D{
		{Key: "Name", Value: e.name},
		{Key: "Employee_Id", Value: e.eid},
		{"tags", bson.A{"ATCI", "Java", e.office}},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(detailsResult.InsertedID)
	assetsResult, err := assetsCollection.InsertOne(ctx, bson.D{
		{"Employee_Id", e.eid},
		{"Name", e.lname},
		{"description", e.ldesc},
		{"Asset_Id", e.lid},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(assetsResult.InsertedID)
}

func Retrieve(e Employee) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://standard:example@cluster0.ut86q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Connected to Database", databases)

	quickstartDatabase := client.Database("Employee")
	detailsCollection := quickstartDatabase.Collection("Details")
	assetsCollection := quickstartDatabase.Collection("Asset")

	filterCursor, err := detailsCollection.Find(ctx, bson.M{"Employee_Id": e.eid}) //Select based on filter value
	if err != nil {
		log.Fatal(err)
	}
	var detailsFiltered []bson.M
	if err = filterCursor.All(ctx, &detailsFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(detailsFiltered)

	filterCursor, err = assetsCollection.Find(ctx, bson.M{"Employee_Id": e.eid}) //Select based on filter value
	if err != nil {
		log.Fatal(err)
	}
	var assetsFiltered []bson.M
	if err = filterCursor.All(ctx, &assetsFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(assetsFiltered)
}

func Update(e Employee) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://standard:example@cluster0.ut86q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Connected to Database", databases)

	quickstartDatabase := client.Database("Employee")
	detailsCollection := quickstartDatabase.Collection("Details")

	result, err := detailsCollection.UpdateOne( //Update One document
		ctx,
		bson.M{"Employee_Id": e.eid},
		bson.D{
			{"$set", bson.D{{"Name", e.nname}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents! \n", result.ModifiedCount)
}

func Delete(e Employee) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://standard:example@cluster0.ut86q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Connected to Database", databases)

	quickstartDatabase := client.Database("Employee")
	detailsCollection := quickstartDatabase.Collection("Details")

	result1, err := detailsCollection.DeleteOne( //Update One document
		ctx,
		bson.M{"Name": e.nname},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v Documents! \n", result1.DeletedCount)
}
