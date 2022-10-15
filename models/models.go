package models

type User struct {
	Id       string    `json:"id" bson:"_id"`
	Username string    `json:"username" bson:"username"`
	Pwd      string    `json:"password" bson:"password"`
	Email    string    `json:"email" bson:"email"`
	Warbles  []*Warble `json:"warbles" bson:"warbles"`
}

type Warble struct {
	Id      string  `json:"id" bson:"_id"`
	User_id *string `json:"user_id" bson:"user_id"`
	Content string  `json:"content" bson:"content"`
}

type Login struct {
	Username string `json:"username" bson:"username"`
	Pwd      string `json:"password" bson:"password"`
}
