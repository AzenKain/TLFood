package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	URL  string   `json:"url,omitempty" bson:"url,omitempty"`
	Link []string `json:"link,omitempty" bson:"link,omitempty"`
}

type Detail struct {
	Tags      []string `json:"tags,omitempty" bson:"tags,omitempty"`
	Company   string   `json:"company,omitempty" bson:"company,omitempty"`
	Materials []string `json:"materials,omitempty" bson:"materials,omitempty"`
}

type Product struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductName   string             `json:"product_name,omitempty" bson:"product_name,omitempty"`
	IsDisplay     bool               `json:"is_display,omitempty" bson:"is_display,omitempty"`
	OriginCost    float64            `json:"origin_cost,omitempty" bson:"origin_cost,omitempty"`
	DisplayCost   float64            `json:"display_cost,omitempty" bson:"display_cost,omitempty"`
	StockQuantity int                `json:"stock_quantity,omitempty" bson:"stock_quantity,omitempty"`
	ImgDisplay    []*Image            `json:"img_display,omitempty" bson:"img_display,omitempty"`
	ProductType   string             `json:"product_type,omitempty" bson:"product_type,omitempty"`
	BuyCount      int               `json:"buy_count,omitempty" bson:"buy_count,omitempty"`
	Detail        *Detail            `json:"detail,omitempty" bson:"detail,omitempty"`
	Description   *string            `json:"description,omitempty" bson:"description,omitempty"`
	UpdatedAt     primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt     primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
