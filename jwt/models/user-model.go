package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Baked_In_Validators_and_Tags
// ep=ADMIN|ep=USER 可选值为ADMIN或USER
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `json:"name"`
	Password  *string            `json:"password"`
	Email     *string            `json:"email" validate:"email,required"`
	Token     *string            `json:"token"`
	Refresh   *string            `json:"refresh"`
	User_type *string            `json:"User_type" validate:"required,eq=ADMIN|eq=USER"`
	Create_at time.Time          `json:"create_at" `
	Update_at time.Time          `json:"update_at" `
	User_id   string             `json:"user_id" `
}
