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
		fmt.Println("Error trying to create person service")
		return
	}

	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			var person domain.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				fmt.Printf("Error trying to decode body. Body should be a json. Error: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusBadRequest)
				return
			}
			if person.ID <= 0 {
				http.Error(w, "Error trying to create person. ID should be a positive integer", http.StatusBadRequest)
				return
			}

			//Criar pessoa
			err = personService.Create(person)
			if err != nil {
				fmt.Printf("Error trying to create person: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusInternalServerError)
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
					http.Error(w, "Error trying to list people", http.StatusInternalServerError)
					return
				}
			} else {
				//Listar pessoa com id
				personID, err := strconv.Atoi(path)
				if err != nil {
					http.Error(w, "Invalid id given. Person ID must be an integer", http.StatusBadRequest)
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
					http.Error(w, "Error trying to encode person as json", http.StatusInternalServerError)
					return
				}
			}
		}

		if r.Method == "PUT" {
			var person domain.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				fmt.Printf("Error trying to decode body. Body should be a json. Error: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusBadRequest)
				return
			}
			if person.ID <= 0 {
				http.Error(w, "Error trying to create person. ID should be a positive integer", http.StatusBadRequest)
				return
			}

			//Editar Pessoa
			err = personService.Edit(person)
			if err != nil {
				fmt.Printf("Error trying to edit person: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)

		}

		if r.Method == "DELETE" {
			path := strings.TrimPrefix(r.URL.Path, "/person/")

			if path == "" {
				http.Error(w, "ID must be provided in the url", http.StatusBadRequest)
				return
			} else {
				personID, err := strconv.Atoi(path)
				if err != nil {
					http.Error(w, "Invalid id given. Person ID must be an integer", http.StatusBadRequest)
					return
				}
				err = personService.DeleteByID(personID)
				if err != nil {
					fmt.Printf("Error trying to delete person: %s", err.Error())
					http.Error(w, "Error trying to delete person", http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
			}

		}

	})

	http.ListenAndServe(":8080", nil)
}