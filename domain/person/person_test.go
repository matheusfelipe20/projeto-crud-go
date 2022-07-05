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

//Teste criar cadastro pessoa
func TestCreate(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/person/", "application/json",
		bytes.NewBuffer([]byte(`{"id":1,"full_name":"Matheus Felipe","cpf":78978978945,"phone":83988774411,"address":"Rua Professora Mocinha Avelar, João Pessoa","date_birth":"20/March/2002"}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Erro no preenchimento dos campos: %v", err)
	}

}

//Teste listar pessoas cadastradas
func TestGetUsers(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/person")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
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

//Teste listar pessoa cadastrada com ID: (1)
func TestGetUsersByID(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/person/1")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
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

//Teste editar pessoa cadastrada
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
		fmt.Printf("Sem sucesso, ID não cadastrado: %d", resp.StatusCode)
	}
}

//Teste deletar pessoa cadastrada pelo ID: (1)
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
		fmt.Printf("Sem sucesso, ID não cadastrado: %d", resp.StatusCode)
	}
}

//Teste de Erro de CPF nulo
func TestErroCreateCpfNulo(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/person/", "application/json",
		bytes.NewBuffer([]byte(`{"id": 2,"full_name":"Bandeira","cpf":0,"phone":83977955590,"address":"Guarabira","date_birth":"22/March/2002"}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Erro no preenchimento dos campos: %v", err)
	}
}

//Teste para dar Erro de deletar cadastro, ID não existente
func TestErroDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/person/100", nil)
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
		fmt.Printf("Sem sucesso, ID não cadastrado: %d", resp.StatusCode)
	}
}

//Teste para dar Erro de editar cadastro, ID não existente
func TestErroEditID(t *testing.T) {

	req, err := http.NewRequest(
		"PUT",
		"http://localhost:8080/person/",
		bytes.NewBuffer([]byte(`{"id":100,"full_name":"Matheus","cpf":78978978555,"phone":83978955578,"address":"Guarabira","date_birth":"20/March/2002"}`)))

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
		fmt.Printf("Sem sucesso, ID não cadastrado: %d", resp.StatusCode)
	}
}

//Teste para dar Erro de ID não encontrado
func TestErroGetID(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/person/100")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))
	pro := Person{}
	err = json.Unmarshal([]byte(string(body)), &pro)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Sem sucesso, ID não cadastrado: %d", resp.StatusCode)
	}

}
