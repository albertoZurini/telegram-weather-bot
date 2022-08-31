package userHandler

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewDBHandler() (*DBHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_HOST"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	db := client.Database("telegram-weather-bot")
	usersCollection := db.Collection("users")

	dbh := &DBHandler{
		clientOptions:   clientOptions,
		client:          client,
		err:             err,
		database:        db,
		usersCollection: usersCollection,
	}

	return dbh, nil
}

func (dbh *DBHandler) PrintDatabases() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := dbh.client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{}
	dbs, err := dbh.client.ListDatabaseNames(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", dbs)
}

func (dbh *DBHandler) SelectUser(UserId int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// attempt to get existing user
	filterUser, err := dbh.usersCollection.Find(ctx, bson.M{"UserId": UserId})
	if err != nil {
		log.Fatal(err)
	}
	var usersFiltered []bson.M
	if err = filterUser.All(ctx, &usersFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Print(usersFiltered)
}

// Write database
func (dbh *DBHandler) SetCityForUser(UserId int64, CityName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	/*
		filter := bson.D{{"UserId", UserId}}

		update := bson.D{{"$set", bson.D{{"CityName", CityName}}}}
		options := options.Update().SetUpsert(true)

		updated, err := dbh.usersCollection.UpdateOne(ctx, filter, update, options)
	*/
	doc := bson.D{{"UserId", UserId}, {"CityName", CityName}}
	updated, err := dbh.usersCollection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(updated)
	dbh.SelectUser(UserId)
}

func (dbh *DBHandler) SetCoordinatesForUser(UserId int64, location Location) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"UserId", UserId}}

	options := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"Location", location}, {"UserId", UserId}, {"CityName", ""}}}}

	updated, err := dbh.usersCollection.UpdateOne(ctx, filter, update, options)

	//updated, err := dbh.usersCollection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updated)
}

func (dbh *DBHandler) SetCoordinatesAndCityForUser(UserId int64, location Location, CityName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"UserId", UserId}}

	options := options.Update().SetUpsert(true)
	update := bson.D{{"$set",
		bson.D{{"Location", location},
			{"UserId", UserId},
			{"CityName", CityName}}}}

	updated, err := dbh.usersCollection.UpdateOne(ctx, filter, update, options)

	//updated, err := dbh.usersCollection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updated)
}

// Read database
func (dbh *DBHandler) GetCityForUser(UserId int64) string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// attempt to get existing user
	filterUser, err := dbh.usersCollection.Find(ctx, bson.M{"UserId": UserId})
	if err != nil {
		log.Fatal(err)
	}
	var usersFiltered []bson.M
	if err = filterUser.All(ctx, &usersFiltered); err != nil {
		log.Fatal(err)
	}

	return usersFiltered[0]["CityName"].(string)
}

func (dbh *DBHandler) GetCoordinatesForUser(UserId int64) string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// attempt to get existing user
	filterUser, err := dbh.usersCollection.Find(ctx, bson.M{"UserId": UserId})
	if err != nil {
		log.Fatal(err)
	}
	var usersFiltered []bson.M
	if err = filterUser.All(ctx, &usersFiltered); err != nil {
		log.Fatal(err)
	}

	return usersFiltered[0]["CityName"].(string) // TODO
}

func (dbh *DBHandler) GetLocationForUser(UserId int64) (Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// attempt to get existing user
	filterUser, err := dbh.usersCollection.Find(ctx, bson.M{"UserId": UserId})
	if err != nil {
		log.Fatal(err)
	}
	var usersFiltered []bson.M
	if err = filterUser.All(ctx, &usersFiltered); err != nil {
		log.Fatal(err)
	}

	return Location{usersFiltered[0]["CityName"].(string),
		[]float64{
			usersFiltered[0]["Location"].(primitive.M)["coordinates"].(primitive.A)[0].(float64),
			usersFiltered[0]["Location"].(primitive.M)["coordinates"].(primitive.A)[1].(float64),
		}}, nil
}
