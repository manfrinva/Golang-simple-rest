package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct para representar um contato
type Contact struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

// Struct para representar um usuário
type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var contacts []Contact
var users = map[string]string{"user": "password"} // Mapeamento de usuários e senhas

// Função para autenticar um usuário
func authenticate(username, password string) bool {
	if pass, ok := users[username]; ok {
		return pass == password
	}
	return false
}

// Endpoint para criar um novo contato
func CreateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

// Endpoint para obter todos os contatos
func GetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

// Endpoint para obter um contato pelo ID
func GetContactByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range contacts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Contact{})
}

// Endpoint para atualizar um contato
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range contacts {
		if item.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			var contact Contact
			_ = json.NewDecoder(r.Body).Decode(&contact)
			contact.ID = params["id"]
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

// Endpoint para deletar um contato
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range contacts {
		if item.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

// Middleware para autenticação
func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !authenticate(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized.\n"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/contacts", BasicAuth(GetContacts)).Methods("GET")
	router.HandleFunc("/contacts/{id}", BasicAuth(GetContactByID)).Methods("GET")
	router.HandleFunc("/contacts", BasicAuth(CreateContact)).Methods("POST")
	router.HandleFunc("/contacts/{id}", BasicAuth(UpdateContact)).Methods("PUT")
	router.HandleFunc("/contacts/{id}", BasicAuth(DeleteContact)).Methods("DELETE")

	// Inicialização de contatos de exemplo
	contacts = append(contacts, Contact{ID: "1", Name: "John Doe", Email: "john@example.com", Phone: "1234567890"})
	contacts = append(contacts, Contact{ID: "2", Name: "Jane Smith", Email: "jane@example.com", Phone: "0987654321"})

	// Inicia o servidor na porta 8000
	fmt.Println("Server running on port :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
