package models

type Address struct {
	State string `json:"state" bson:"user_state"`
	City string `json:"city" bson:"user_city"`
	Pincode int `json:"pincode" bson:"user_pincode"`
}

type User struct {
	FirstName string `json:"firstName" bson:"user_first_name"`
	LastName  string `json:"lastName" bson:"user_last_name"`
	Age int32 `json:"age" bson:"user_age"`
	Address Address `json:"address" bson:"user_address"`
}