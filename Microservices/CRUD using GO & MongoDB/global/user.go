package global

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User is the default user struct
type User struct {
	ID      primitive.ObjectID `bson:"_id"`
	EID     string             `bson:"eid"`
	Empname string             `bson:"ename"`
	Level   string             `bson:"level"`
	Stream  string             `bson:"stream"`
}

//NilUser when no user found
var NilUser User

//GetToken returns the JSON's web tokens
func (u User) GetToken() string {
	byteSlc, _ := json.Marshal(u)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": string(byteSlc),
	})
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString

}

//returns user authenticated with this token
func UserFromToken(token string) User {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	var result User
	json.Unmarshal([]byte(claims["data"].(string)), &result)
	return result
}
