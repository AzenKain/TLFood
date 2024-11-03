package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDetail struct {
	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Birthday  primitive.DateTime `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Gender    string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Address   string             `json:"address,omitempty" bson:"address,omitempty"`
	PhoneNumber  string          `json:"phone_number,omitempty" bson:"phone_number,omitempty" unique:"true"`
}
type User struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CartId       string             `json:"cart_id,omitempty" bson:"cart_id,omitempty" unique:"true"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty" unique:"true"`
	IsDisplay    bool               `json:"is_display,omitempty" bson:"is_display,omitempty"`
	Hash         string             `json:"hash,omitempty" bson:"hash,omitempty"`
	Role         []string           `json:"role,omitempty" bson:"role,omitempty"`
	RefreshToken string             `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
	HeartList    []string           `json:"heart_list,omitempty" bson:"heart_list,omitempty"`
	Detail       *UserDetail         `json:"detail,omitempty" bson:"detail,omitempty"`
	UpdatedAt    primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt    primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
