package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"project/database"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email string
	Name  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllToken(email string, name string) (signedToken string, signedRefresh string, err error) {
	claims := &SignedDetails{
		Email: email,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refresh, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
	}
	return token, refresh, err
}

func UpdateAllTokens(signedToken string, signedRefresh string, user_id string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var updateobj primitive.D
	updateobj = append(updateobj, bson.E{"token", signedToken})
	updateobj = append(updateobj, bson.E{"refresh", signedRefresh})
	update_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateobj = append(updateobj, bson.E{"update_at", update_at})
	upsert := true
	filter := bson.M{"user_id": user_id}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	defer cancel()
	userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"&set", updateobj},
		},
		&opt,
	)
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("token is invalid")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expried")
		msg = err.Error()
		return
	}
	return claims, msg
}
