package person_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

type Person struct {
	Results []struct {
		ID        int    `json:"id"`
		FullName  string `json:"full_name"`
		Cpf       int    `json:"cpf"`
		Phone     int    `json:"phone"`
		Address   string `json:"address"`
		DateBirth string `json:"date_birth"`
	} `json:"results"`
	Status string `json:"status"`
}

func TestCreate(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/person/", "application/json",
		bytes.NewBuffer([]byte(`{"id":7,"full_name":"Felipe","cpf":78978978555,"phone":83978955578,"address":"Guarabira","date_birth":"19/April/1999"}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Erro no preenchimento dos campos: %v", err)
	}
}

//CPF nulo
func TestCreateErrorCpfNulo(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/person/", "application/json",
		bytes.NewBuffer([]byte(`{"id":,"full_name":"Felipe","cpf":0,"phone":83978955578,"address":"Guarabira","date_birth":"19/April/1999"}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Erro no preenchimento dos campos: %v", err)
	}
}

//Listar pessoas cadastradas
func TestGetUsers(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/person")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	pro := Person{}
	err = json.Unmarshal([]byte(string(body)), &pro)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Sem sucesso: %d", resp.StatusCode)
	}
}

//Listar pessoa cadastraada com ID
func TestGetUsersByID(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/person/1")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	pro := Person{}
	err = json.Unmarshal([]byte(string(body)), &pro)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Sem sucesso: %d", resp.StatusCode)
	}

}

//Editar pessoa cadastrada
func TestEditUser(t *testing.T) {

	req, err := http.NewRequest(
		"PUT",
		"http://localhost:8080/person/",
		bytes.NewBuffer([]byte(`{"id":1,"full_name":"Matheus","cpf":78978978555,"phone":83978955578,"address":"Guarabira","date_birth":"20/March/2002"}`)))

	if err != nil {
		t.Error(err)
	}

	defer req.Body.Close()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Sem sucesso pessoa não cadastrada: %d", resp.StatusCode)
	}
}

//Deletar pessoa cadastrada pelo ID (2)
func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/person/2", nil)
	if err != nil {
		t.Error("*********************************************")
		t.Error(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("---------------------------------------------")
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Sem sucesso pessoa não cadastrada: %d", resp.StatusCode)
	}
}
