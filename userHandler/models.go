package userHandler

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHandler struct {
	clientOptions   *options.ClientOptions `json:-`
	client          *mongo.Client          `json:-`
	err             error                  `json:-`
	database        *mongo.Database        `json:-`
	usersCollection *mongo.Collection      `json:-`
}

type Location struct {
	GeoJSONType string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type UserLocation struct {
	Id       string   `json:"id" bson:"_id"`
	UserId   int64    `json:"userId" bson:"usrId"`
	Location Location `json:"location" bson:"location"`
	CityName string   `json:"cityName" bson:"cityName"`
}
