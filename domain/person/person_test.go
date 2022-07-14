package person_test

import (
	"bytes"
	"encoding/json"
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
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		t.Errorf("Erro no preenchimento dos campos: %v", err)
	}

	//Testar se o cadastro foi registrado na listar com o ID: 1
	if resp.StatusCode == http.StatusCreated {
		resp, erro := http.Get("http://localhost:8080/person/1")
		if erro != nil {
			t.Errorf("Erro ao fazer requisição: %v", erro)
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
			log.Println(err)
		}

		if resp.StatusCode != 200 {
			t.Errorf("Sem sucesso: %v", string(body))
		}

	} else {
		t.Errorf("Sem sucesso: %v", string(body))
	}

}

//Teste listar todas as pessoas cadastradas
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
		log.Println(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso: %v", string(body))
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
		log.Println(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso: %v", string(body))
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso: %v", string(body))
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso: %v", string(body))
	}
}

//Teste_Erro CPF nulo
func TestErroCreateCpfNulo(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/person/", "application/json",
		bytes.NewBuffer([]byte(`{"id": 12,"full_name":"Bandeira","cpf":0,"phone":83977955590,"address":"Guarabira","date_birth":"22/March/2002"}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Erro no preenchimento dos campos: %v", err)
	}

	if resp.StatusCode == 500 {
		t.Errorf("Sem sucesso: %v", string(body))
	}
}

//Teste_Erro Create: ID negativo
func TestErroCreateIDNegative(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/person/", "application/json",
		bytes.NewBuffer([]byte(`{"id": -7,"full_name":"Bandeira","cpf":12345678989,"phone":83977955590,"address":"Guarabira","date_birth":"22/March/2002"}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Erro no preenchimento dos campos: %v", err)
	}

	if resp.StatusCode == 400 {
		t.Errorf("Sem sucesso: %v", string(body))
	}
}

//Teste_Erro Create: Nome vazio
func TestErroCreateNameNull(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/person/", "application/json",
		bytes.NewBuffer([]byte(`{"id": 12,"full_name":"","cpf":12345678989,"phone":83977955590,"address":"Guarabira","date_birth":"22/March/2002"}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Erro no preenchimento dos campos: %v", err)
	}

	if resp.StatusCode == 500 {
		t.Errorf("Sem sucesso: %v", string(body))
	}
}

//Teste_Erro Create: Endereço vazio
func TestErroCreateAddressNull(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/person/", "application/json",
		bytes.NewBuffer([]byte(`{"id": 12,"full_name":"Bandeira","cpf":12345678989,"phone":83977955590,"address":"","date_birth":"22/March/2002"}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Erro no preenchimento dos campos: %v", err)
	}

	if resp.StatusCode == 500 {
		t.Errorf("Sem sucesso: %v", string(body))
	}
}

//Teste_Erro Create: Date Birth vazio
func TestErroCreateDateBirthNull(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/person/", "application/json",
		bytes.NewBuffer([]byte(`{"id": 12,"full_name":"Bandeira","cpf":12345678900,"phone":83977955590,"address":"Guarabira","date_birth":""}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Erro no preenchimento dos campos: %v", err)
	}

	if resp.StatusCode == 500 {
		t.Errorf("Sem sucesso: %v", string(body))
	}
}

//Teste_Erro Deletar cadastro: ID não existente
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 500 {
		t.Errorf("Sem sucesso: %v", string(body))
	}
}

//Teste_Erro Editar cadastro: ID não existente
func TestErroEditID(t *testing.T) {

	req, err := http.NewRequest(
		"PUT",
		"http://localhost:8080/person/",
		bytes.NewBuffer([]byte(`{"id":100,"full_name":"Matheus","cpf":78978978559,"phone":83978955578,"address":"Guarabira","date_birth":"20/March/2002"}`)))

	if err != nil {
		t.Error(err)
	}

	defer req.Body.Close()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode == 500 {
		t.Errorf("Sem sucesso: %v", string(body))
	}
}

//Teste_Erro Get: ID não cadastrado
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
		log.Println(err)
	}

	if resp.StatusCode == 404 {
		t.Errorf("Sem sucesso: %v", string(body))
	}

}
