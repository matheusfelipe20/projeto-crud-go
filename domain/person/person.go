package person

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"os"

	"github.com/matheusfelipe20/projeto-crud/domain"
)

type Service struct {
	dbFilePath string
	people     domain.People
}

func NewService(dbFilePath string) (Service, error) {
	//Verificação do arquivo existente
	_, err := os.Stat(dbFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			//criar um arquivo vazio
			err = createEmptyFile(dbFilePath)
			if err != nil {
				return Service{}, err
			}
			return Service{
				dbFilePath: dbFilePath,
				people:     domain.People{},
			}, nil
		}
	}

	//Caso já exista leia e atualize a variavel people com as pessoas do arquivo
	jsonFile, err := os.Open(dbFilePath)
	if err != nil {
		return Service{}, fmt.Errorf("Erro ao tentar abrir arquivo que contém todas as pessoas: %s", err.Error())
	}

	jsonFileContenByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Service{}, fmt.Errorf("Erro ao tentar ler o arquivo: %s", err.Error())
	}

	var allPeople domain.People
	json.Unmarshal(jsonFileContenByte, &allPeople)

	return Service{
		dbFilePath: dbFilePath,
		people:     allPeople,
	}, nil

}

//Função de criar arquivo
func createEmptyFile(dbFilePath string) error {
	var people domain.People = domain.People{
		People: []domain.Person{},
	}
	peopleJSON, err := json.Marshal(people)
	if err != nil {
		return fmt.Errorf("Erro ao tentar codificar pessoas como JSON?: %s", err.Error())
	}

	err = ioutil.WriteFile(dbFilePath, peopleJSON, 0755)
	if err != nil {
		return fmt.Errorf("Erro ao tentar gravar no arquivo. Error: %s", err.Error())
	}

	return nil
}

func (s *Service) Create(person domain.Person) error {
	//Verificar se a pessoa já existe
	if s.exists(person) {
		return fmt.Errorf("Erro ao tentar criar pessoa. Existe uma pessoa com este ID ou CPF já cadastrado")
	}

	//Verificar quantidade de digitos
	if s.checkQuantity(person) {
		return fmt.Errorf("Erro ao tentar criar pessoa. Erro número incorreto de CPF ou números de telefone")
	}

	//Verificar obrigatoriedade campos
	if s.obrigatory(person) {
		return fmt.Errorf("Erro ao tentar criar pessoa. Erro ao preencher os dados")
	}

	// adicinar pessoa
	s.people.People = append(s.people.People, person)

	//salvar o arquivo
	err := s.saveFile()
	if err != nil {
		return fmt.Errorf("Erro ao tentar salvar arquivo no método Create. Error: %s", err.Error())
	}

	return nil
}

//Função para verificar a quantidade de digito
func (s Service) checkQuantity(person domain.Person) bool {
	quantityCPF := strconv.Itoa(person.Cpf)
	quantityPhone := strconv.Itoa(person.Phone)
	//o cpf será verificado se tem 11 digitos
	if len(quantityCPF) != 11 {
		return true
	}
	//Se o telefone foi preechido será verificado se tem 11 digitos
	if person.Phone != 0 {
		if len(quantityPhone) != 11 {
			return true
		}
	}
	return false
}

//Função para tornar os campos obrigatorios
func (s Service) obrigatory(person domain.Person) bool {
	if person.FullName == "" || person.Cpf == 0 || person.Address == "" {
		return true
	}
	return false
}

//Função para verificar se a pessoa já existe
func (s Service) exists(person domain.Person) bool {
	for _, currentPerson := range s.people.People {
		if currentPerson.ID == person.ID || currentPerson.Cpf == person.Cpf {
			return true
		}
	}
	return false
}

func (s Service) saveFile() error {
	allPeopleJSON, err := json.Marshal(s.people)
	if err != nil {
		return fmt.Errorf("Erro ao tentar codificar pessoas como json: %s", err.Error())
	}
	return ioutil.WriteFile(s.dbFilePath, allPeopleJSON, 0755)
}

func (s Service) List() domain.People {
	return s.people
}

func (s Service) GetByID(personID int) (domain.Person, error) {
	for _, currentPerson := range s.people.People {
		if currentPerson.ID == personID {
			return currentPerson, nil
		}
	}
	return domain.Person{}, fmt.Errorf("Pessoa não encontrada")
}

//Função para encontrar a pessoa para editar pelo ID
func (s *Service) Edit(person domain.Person) error {
	var indexToEdit int = -1
	for index, currentPerson := range s.people.People {
		if currentPerson.ID == person.ID {
			indexToEdit = index
			break
		}
	}
	if indexToEdit < 0 {
		return fmt.Errorf("Não há pessoa com o ID fornecido em nosso banco de dados")
	}

	s.people.People[indexToEdit] = person
	return s.saveFile()
}

func (s *Service) DeleteByID(personID int) error {
	var indexToDelete int = -1
	for index, currentPerson := range s.people.People {
		if currentPerson.ID == personID {
			indexToDelete = index
			break
		}
	}
	if indexToDelete < 0 {
		return fmt.Errorf("Não há nenhuma pessoa com o ID fornecido em nosso banco de dados")
	}

	s.people.People = append(s.people.People[:indexToDelete], s.people.People[indexToDelete+1:]...)

	return s.saveFile()
}
