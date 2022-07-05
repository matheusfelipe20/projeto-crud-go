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

func TestEditUser(t *testing.T) {
	req, err := http.NewRequest(
		"PUT",
		"http://localhost:8080/person/1",
		bytes.NewReader([]byte(`{"full_name":"Matheus", "cpf":12345678978, "phone":83955447788, 
		"address": "Rua tal tal, joao pessoa", "date_birth": "20/March/2005"}`)))
	if err != nil {
		t.Error(err)
	}
	defer req.Body.Close()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Sem sucesso pessoa não cadastrada: %d", resp.StatusCode)
	}

}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/person/1", nil)
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
