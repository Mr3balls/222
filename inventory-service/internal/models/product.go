package models

type Product struct {
	ID       string  `json:"id" bson:"_id,omitempty"`
	Name     string  `json:"name" bson:"name"`
	Category string  `json:"category" bson:"category"`
	Stock    int     `json:"stock" bson:"stock"`
	Price    float64 `json:"price" bson:"price"`
}
