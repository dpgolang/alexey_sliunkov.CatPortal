package main

import(
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/subosito/gotenv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "viewsonic2000"
	dbname   = "postgres"
)

var db *sql.DB

type Animal struct{
	ID int `json:id`
	Breed string `json:breed`
	Size string `json:size`
	Diet string `json:diet`
	Motherland string `json:motherland`
	Description string `json:description`
}

var cats []Animal

func main(){
	ConnectDB()
	router := mux.NewRouter()

	cats = append(cats,Animal{ID:1,Breed:"Abyssinian cat",Size:"24*36, 4,5kg",Diet:"Purina A45X",Motherland:"England",Description:"extremely inquisitive"},
	Animal{ID:2,Breed:"American Bobcat",Size:"30*45, 5,5kg",Diet:"Purina A16M",Motherland:"US",Description:"highly intelligent"},
	Animal{ID:3,Breed:"Bengal Cat",Size:"26*35, 4kg",Diet:"Purina M45C",Motherland:"US",Description:"highly intelligent"})

	router.HandleFunc("/cats",getAnimals).Methods("GET")
	router.HandleFunc("/cats/{id}",getAnimal).Methods("GET")
	router.HandleFunc("/cats",addAnimal).Methods("POST")
	router.HandleFunc("/cats",updateAnimal).Methods("PUT")
	router.HandleFunc("/cats/{id}",removeAnimal).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",router))
}

func getAnimals(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(cats)
}

func getAnimal(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	i,_ := strconv.Atoi(params["id"])

	for _,cat := range cats{
		if cat.ID == i {
			json.NewEncoder(w).Encode(&cat)
		}
	}
}

func addAnimal(w http.ResponseWriter, r *http.Request){
	var cat Animal
	_ = json.NewDecoder(r.Body).Decode(&cat)

	cats = append(cats,cat)
	json.NewEncoder(w).Encode(cats)
}

func updateAnimal(w http.ResponseWriter, r *http.Request){
	var cat Animal
	json.NewDecoder(r.Body).Decode(&cat)

	for i,item := range cats{
		if item.ID == cat.ID{
			cats[i] = cat
		}
	}
	json.NewEncoder(w).Encode(cats)
}

func removeAnimal(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	id,_ := strconv.Atoi(params["id"])
	for i,item := range cats{
		if item.ID == id{
			cats = append(cats[:i],cats[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(cats)
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
