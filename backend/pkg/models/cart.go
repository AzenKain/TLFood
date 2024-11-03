package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartItem struct {
	ProductId string `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Quantity  int    `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

type Cart struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId    string             `json:"user_id,omitempty" bson:"user_id,omitempty" unique:"true"`
	IsDisplay bool               `json:"is_display,omitempty" bson:"is_display,omitempty"`
	Items     []*CartItem        `json:"items,omitempty" bson:"items,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
