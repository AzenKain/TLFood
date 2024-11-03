package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeliveryInfo struct {
	City     string `json:"city,omitempty" bson:"city,omitempty"`
	District string `json:"district,omitempty" bson:"district,omitempty"`
	Address  string `json:"address" bson:"address"`
}

type PersonalInfo struct {
	Email       string `json:"email" bson:"email"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

type PayDetail struct {
	IsPaid    bool               `json:"is_paid,omitempty" bson:"is_paid,omitempty"`
	Bank      string             `json:"bank,omitempty" bson:"bank,omitempty"`
	TrackID   string             `json:"track_id,omitempty" bson:"track_id,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}


type OrderItem struct {
	ProductID   string  `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Quantity    int     `json:"quantity,omitempty" bson:"quantity,omitempty"`
	DisplayCost float64  `json:"display_cost,omitempty" bson:"display_cost,omitempty"`
	OriginCost  float64  `json:"origin_cost,omitempty" bson:"origin_cost,omitempty"`
	Discount    float64  `json:"discount,omitempty" bson:"discount,omitempty"`
}

type Order struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID          string             `json:"user_id" bson:"user_id"`
	IsDisplay       bool               `json:"is_display" bson:"is_display"`
	TotalAmount     float64            `json:"total_amount" bson:"total_amount"`
	Status          string             `json:"status" bson:"status"`
	ListProducts    []*OrderItem       `json:"list_products,omitempty" bson:"list_products,omitempty"`
	Notes           string            `json:"notes,omitempty" bson:"notes,omitempty"`
	DeliveryInfo    *DeliveryInfo      `json:"delivery_info" bson:"delivery_info"`
	DeliveryType    string             `json:"delivery_type" bson:"delivery_type"`
	PersonalDetails *PersonalInfo      `json:"personal_details" bson:"personal_details"`
	PaymentMethods  string             `json:"payment_methods" bson:"payment_methods"`
	PayDetail       *PayDetail             `json:"is_pay" bson:"is_pay"`
	UpdatedAt       primitive.DateTime `json:"updated_at" bson:"updated_at"`
	CreatedAt       primitive.DateTime `json:"created_at" bson:"created_at"`
}
