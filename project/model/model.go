package model

type Animal struct{
	ID int `json:id`
	Breed string `json:breed`
	Size string `json:size`
	Diet string `json:diet`
	Motherland string `json:motherland`
	Description string `json:description`
}

type User struct {
	Id          int     `json:id`
	Firstname   string  `json:firstname,omitempty`
	Lastname    string  `json:lastname,omitempty`
	Password 	string `json:password,omitempty`
}
type Food struct{
	Id	int     `json:id`
	Name   string  `json:name`
	Price    string  `json:price`
	Composition 	string `json:composition`
}


