package model

type Animal struct{
	ID int `json:id`
	Breed string `json:breed`
	Size string `json:size`
	Diet string `json:diet`
	Motherland string `json:motherland`
	Description string `json:description`
}


