package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Model       string             `bson:"model" json:"model"`
    Brand       string             `bson:"brand" json:"brand"`
    Year        string             `bson:"year" json:"year"`
    Color       string             `bson:"color" json:"color"`
    Price       int                `bson:"price" json:"price"`
    Description string             `bson:"description" json:"description"`
    UserID      primitive.ObjectID `bson:"userId" json:"userId"`
    ImageUrl    string             `bson:"imgurl" json:"imgurl"`
}
