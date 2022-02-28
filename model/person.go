package model

type Person struct {
	Name      string `jason:"name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Ocupation string `json:"ocupation"`
}

type Peopel []Person
