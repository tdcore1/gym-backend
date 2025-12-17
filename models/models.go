package models

type Wallet struct {
	Amount int `json:"amount"`
}
type User struct {
	Name     string `json:"name"`
	Lastname string `json:"last_name"`
	Password int    `json:"password"`
	Coachid  int    `json:"coach_id"`
	Courseid int    `json:"course_id"`
}
type Coach struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Day   int    `json:"day"`
}
type Course struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Day   int    `json:"day"`
}
type Login struct {
	User     int `json:"user"`
	Password int `json:"password"`
}
type Selectt struct {
	Courseid int `json:"Courseid"`
	Coachid  int `json:"Coachid"`
}
