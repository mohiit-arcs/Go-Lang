package models

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode uint   `json:"pincode" bson:"pincode"`
}

type User struct {
	Name    string  `json:"name" bson:"user_name"`
	Age     uint    `json:"age" bson:"user_age"`
	Address Address `json:"address" bson:"user_address"`
}
