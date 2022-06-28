package domain

type Person struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Cpf      int    `json:"cpf"`
	Phone    int    `json:"phone"`
	Address  string `json:"address"`
}

type People struct {
	People []Person `json:"people"`
}
