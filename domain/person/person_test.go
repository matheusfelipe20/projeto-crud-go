package person_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/matheusfelipe20/projeto-crud/domain"
)

func TestCreateUser(t *testing.T) {
	person := domain.Person{
		ID:        1,
		FullName:  "Matheus",
		Cpf:       12345678912,
		Phone:     83988447799,
		Address:   "Rua Epitacio Pessoa, João Pessoa",
		DateBirth: "20/March/2002",
	}

	bodyRequestJson := new(bytes.Buffer)
	encodeJson, erro := json.Marshal(person)
	if erro != nil {
		return erro
	}
	bodyRequestJson.Write(encodeJson)

	request, erro := http.NewRequest("POST", "/person/", bodyRequestJson)
	client := &http.Client{}
	resp, erro := client.Do(request)
	if erro != nil {
		return erro
		//alertar o error
		//tratar o error
	}

	//testar a resposta de usuario criado

}

func TestGetUsers(t *testing.T) {

}

func TestGetUsersByID(t *testing.T) {

}

func TestEditUser(t *testing.T) {

}

func TestDeleteUser(t *testing.T) {

}
