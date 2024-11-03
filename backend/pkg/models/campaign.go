package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Campaign struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IsDisplay    bool               `json:"is_display" bson:"is_display"`
	CampaignType string             `json:"campaign_type,omitempty" bson:"campaign_type,omitempty"`
	Percents     int                `json:"percents,omitempty" bson:"percents,omitempty"`
	Detail       string             `json:"detail,omitempty" bson:"detail,omitempty"`
	ImgDisplay   []string           `json:"img_display,omitempty" bson:"img_display,omitempty"`
	StartTime    primitive.DateTime `json:"start_time,omitempty" bson:"start_time,omitempty"`
	EndTime      primitive.DateTime `json:"end_time,omitempty" bson:"end_time,omitempty"`
	UpdatedAt    primitive.DateTime `json:"updated_at" bson:"updated_at"`
	CreatedAt    primitive.DateTime `json:"created_at" bson:"created_at"`
}
