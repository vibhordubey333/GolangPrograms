package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	userList = make([]Users, 0)
)

// For proper marshalling Struct members first character should be capital otherwise response will be empty.
type Users struct {
	UserID     string `json:"userID"`
	Name       string `json:"name"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/api/v1/users", users).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/v1/users", createUser).Methods("POST")
	//Setting up port and address.
	fmt.Println("Listening and serving on http://localhost:7777")
	log.Fatal(http.ListenAndServe(":7777", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome To Home Page")
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	userList := populateUsersList()

	json.NewEncoder(w).Encode(userList)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	id := mux.Vars(r)["id"]
	fmt.Println("Mux :", mux.Vars(r))
	userPayload, err := findUser(id)
	if err != nil {
		json.NewEncoder(w).Encode("UserID not found")
		return
	}
	json.NewEncoder(w).Encode(userPayload)

}

func findUser(userID string) (Users, error) {
	log.Print("UserID: ", userID)
	userList := populateUsersList()
	for _, v := range userList {
		fmt.Println("UserID print : ", v.UserID)
		if v.UserID == userID {
			log.Println("UserID found !!!")
			return Users{
				UserID:     v.UserID,
				Name:       v.Name,
				Country:    v.Country,
				PostalCode: v.PostalCode,
			}, nil
		}
	}
	log.Println("UserID not found !!!")
	return Users{}, errors.New("UserID doesn't exist.")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("CreateUser:", r.Body)
	json.NewEncoder(w).Encode("User Created Successfully")
}

func populateUsersList() []Users {

	userList = append(userList,
		Users{
			UserID:     "1",
			Name:       "Alpha",
			Country:    "IN",
			PostalCode: "200001",
		},
		Users{
			UserID:     "2",
			Name:       "Bravo",
			Country:    "RU",
			PostalCode: "400001",
		},
	)
	return userList
}

/*
Reference:
1. https://www.golinuxcloud.com/go-gorilla-mux/
*/
/*
1. $ curl http://localhost:7777
"Welcome To Home Page"

2. http://localhost:7777/users/1
Method: GET

3. http://localhost:7777/api/v1/users
METHOD: POST
Body:
{
    "userID": "6",
    "name": "BuBerry",
    "country": "XX",
    "postalCode": "343434"
}
*/
