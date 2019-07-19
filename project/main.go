package main

import(
	//"fmt"
	//"encoding/json"
	"log"
	"net/http"
	"project/controller"
	//"strconv"
	"github.com/gorilla/mux"
	"project/model"
	"project/driver"
	"database/sql"
	_ "github.com/lib/pq"
	//_ "github.com/subosito/gotenv"
)

//const (
////	host     = "localhost"
////	port     = 5432
////	user     = "postgres"
////	password = "1878"
////	dbname   = "postgres"
////)
////
var db *sql.DB
var cats []model.Animal

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main(){
	db =driver.ConnectDB()
	router := mux.NewRouter()
	controller1 := controller.Controller{}

	//cats = append(cats,model.Animal{ID:1,Breed:"Abyssinian cat",Size:"24*36, 4,5kg",Diet:"Purina A45X",Motherland:"England",Description:"extremely inquisitive"},
	//model.Animal{ID:2,Breed:"American Bobcat",Size:"30*45, 5,5kg",Diet:"Purina A16M",Motherland:"US",Description:"highly intelligent"},
	//model.Animal{ID:3,Breed:"Bengal Cat",Size:"26*35, 4kg",Diet:"Purina M45C",Motherland:"US",Description:"highly intelligent"})

	router.HandleFunc("/cats",controller1.GetAnimals(db)).Methods("GET")
	router.HandleFunc("/cats/{id}",controller1.GetAnimal(db)).Methods("GET")
	router.HandleFunc("/cats",controller1.AddAnimal(db)).Methods("POST")
	router.HandleFunc("/cats",controller1.UpdateAnimal(db)).Methods("PUT")
	router.HandleFunc("/cats/{id}",controller1.RemoveAnimal(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",router))
}

//func getAnimals(w http.ResponseWriter, r *http.Request){
//	json.NewEncoder(w).Encode(cats)
//}
//
//func getAnimal(w http.ResponseWriter, r *http.Request){
//	params := mux.Vars(r)
//
//	i,_ := strconv.Atoi(params["id"])
//
//	for _,cat := range cats{
//		if cat.ID == i {
//			json.NewEncoder(w).Encode(&cat)
//		}
//	}
//}
//
//func addAnimal(w http.ResponseWriter, r *http.Request){
//	var cat model.Animal
//	_ = json.NewDecoder(r.Body).Decode(&cat)
//
//	cats = append(cats,cat)
//	json.NewEncoder(w).Encode(cats)
//}
//
//func updateAnimal(w http.ResponseWriter, r *http.Request){
//	var cat model.Animal
//	json.NewDecoder(r.Body).Decode(&cat)
//
//	for i,item := range cats{
//		if item.ID == cat.ID{
//			cats[i] = cat
//		}
//	}
//	json.NewEncoder(w).Encode(cats)
//}
//
//func removeAnimal(w http.ResponseWriter, r *http.Request){
//	params := mux.Vars(r)
//
//	id,_ := strconv.Atoi(params["id"])
//	for i,item := range cats{
//		if item.ID == id{
//			cats = append(cats[:i],cats[i+1:]...)
//		}
//	}
//	json.NewEncoder(w).Encode(cats)
//}

//func ConnectDB() *sql.DB {
//	var err error
//
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//
//	db, err = sql.Open("postgres", psqlInfo)
//	logFatal(err)
//
//	err = db.Ping()
//	logFatal(err)
//
//	return db
//}
