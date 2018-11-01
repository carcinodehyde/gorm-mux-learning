package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carcinodehyde/rest-api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var people []models.User
var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root@/test-golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.User{})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{privyid}", GetPerson).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{privyid}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetPeople return list of people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	db.Find(&people)
	json.NewEncoder(w).Encode(&people)
}

// GetPerson return a single person by PrivyID
func GetPerson(w http.ResponseWriter, r *http.Request) {
	var person models.User
	params := mux.Vars(r)
	db.Where("privy_id = ?", params["privyid"]).First(&person)
	json.NewEncoder(w).Encode(&person)
}

// CreatePerson creates a new person data
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.User
	_ = json.NewDecoder(r.Body).Decode(&person)
	db.NewRecord(person)
	db.Create(&person)
	json.NewEncoder(w).Encode(&person)
}

// DeletePerson delete person data from collection
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db.Where("privy_id = ?", params["privyid"]).Delete(&models.User{})
}
