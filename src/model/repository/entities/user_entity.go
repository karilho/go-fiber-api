package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// the omitempty tag is used to not return the field if it is empty.
// the bson tag is used to map the field to the database, but it's exclusive to mongoDB, so we dont put it
// with others "classes" of database.
type UserEntityStruct struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Age      int8               `bson:"age,omitempty"`
}
