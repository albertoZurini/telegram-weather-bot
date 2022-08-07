package userHandler

func NewLocation(lat, lng float64) Location {
	return Location{
		GeoJSONType: "Point",
		Coordinates: []float64{lng, lat},
	}
}

type Point struct {
	Id       string   `json:"id" bson:"_id"`
	Location Location `json:"location" bson:"location"`
}
