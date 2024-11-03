package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Otp struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" unique:"true"`
	IsDisplay bool               `json:"is_display,omitempty" bson:"is_display,omitempty"`
	OtpCode   string             `json:"otp_code,omitempty" bson:"otp_code,omitempty"`
	OtpType   string             `json:"otp_type,omitempty" bson:"otp_type,omitempty"`
	Value     bool               `json:"value,omitempty" bson:"value,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
