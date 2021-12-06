package global

import "go.mongodb.org/mongo-driver/bson/primitive"

//User is the default user struct
type User struct {
	ID      primitive.ObjectID `bson:"_id"`
	EID     string             `bson:"eid"`
	Empname string             `bson:"ename"`
	Level   string             `bson:"level"`
	Stream  string             `bson:"stream"`
}
