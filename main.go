package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/matheusfelipe20/projeto-crud/domain"
	"github.com/matheusfelipe20/projeto-crud/domain/person"
)

func main() {
	personService, err := person.NewService("person.json")
	if err != nil {
		fmt.Println("Error ao tentar criar o cadastro da pessoa")
		return
	}

	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			var person domain.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				fmt.Printf("Erro ao tentar decodificar o corpo. O corpo deve ser um json. Error: %s", err.Error())
				http.Error(w, "Error ao tentar criar cadastro", http.StatusBadRequest)
				return
			}
			if person.ID <= 0 {
				http.Error(w, "Erro ao tentar criar o cadastro pessoa. O ID deve ser um número inteiro positivo", http.StatusBadRequest)
				return
			}

			//Criar pessoa
			err = personService.Create(person)
			if err != nil {
				fmt.Printf("Erro ao tentar criar cadastro pessoa: %s", err.Error())
				http.Error(w, "Erro ao tentar criar cadastro pessoa", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			return

		}
		if r.Method == "GET" {
			path := strings.TrimPrefix(r.URL.Path, "/person/")
			if path == "" {

				//Listar todas as pessoas
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				people := personService.List()
				err := json.NewEncoder(w).Encode(people)
				if err != nil {
					http.Error(w, "Erro ao tentar listar pessoas", http.StatusInternalServerError)
					return
				}
			} else {
				//Listar pessoa com id
				personID, err := strconv.Atoi(path)
				if err != nil {
					http.Error(w, "ID inválido fornecido. O ID da pessoa deve ser um número inteiro", http.StatusBadRequest)
					return
				}
				person, err := personService.GetByID(personID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-type", "application/json")
				err = json.NewEncoder(w).Encode(person)
				if err != nil {
					http.Error(w, "Erro ao tentar codificar pessoa como json", http.StatusInternalServerError)
					return
				}
			}
		}

		if r.Method == "PUT" {
			var person domain.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				fmt.Printf("Erro ao tentar decodificar o corpo. O corpo deve ser um json. Error: %s", err.Error())
				http.Error(w, "Erro ao tentar criar cadastro de pessoa", http.StatusBadRequest)
				return
			}
			if person.ID <= 0 {
				http.Error(w, "Erro ao tentar criar cadastro de pessoa. O ID deve ser um número inteiro positivo", http.StatusBadRequest)
				return
			}

			//Editar Pessoa
			err = personService.Edit(person)
			if err != nil {
				fmt.Printf("Erro ao tentar editar a pessoa: %s", err.Error())
				http.Error(w, "Erro ao tentar criar cadastro pessoa", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			fmt.Println("Pessoa editada com sucesso")

		}

		if r.Method == "DELETE" {
			path := strings.TrimPrefix(r.URL.Path, "/person/")

			if path == "" {
				http.Error(w, "O ID deve ser fornecido na URL", http.StatusBadRequest)
				return
			} else {
				personID, err := strconv.Atoi(path)
				if err != nil {
					http.Error(w, "ID inválido fornecido. O ID da pessoa deve ser um número inteiro", http.StatusBadRequest)
					return
				}
				err = personService.DeleteByID(personID)
				if err != nil {
					fmt.Printf("Erro ao tentar excluir cadastro de pessoa: %s", err.Error())
					http.Error(w, "Erro ao tentar excluir cadastro de pessoa", http.StatusInternalServerError)
					return
				}
				fmt.Println("Pessoa deletada com sucesso")
				w.WriteHeader(http.StatusOK)
			}

		}

	})

	http.ListenAndServe(":8080", nil)
}
