package main

import(
	"github.com/jmoiron/sqlx"
	//"fmt"
	//"encoding/json"
	"log"
	"net/http"
	"os"
	"project/controller"
	//"strconv"
	"github.com/gorilla/mux"
	"project/model"
	"project/driver"
	_ "github.com/lib/pq"
	"github.com/gorilla/handlers"
	//_ "github.com/subosito/gotenv"
)

var db *sqlx.DB
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

	router.HandleFunc("/signup", controller1.Signup(db)).Methods("POST")
	router.HandleFunc("/signin", controller1.Signin(db)).Methods("POST")
	router.HandleFunc("/signadmin", controller1.SignAdmin(db)).Methods("POST")
	router.HandleFunc("/cats",controller1.GetAnimals(db)).Methods("GET")
	router.HandleFunc("/cats/{id}",controller1.GetAnimal(db)).Methods("GET")
	router.HandleFunc("/meals",controller1.GetMeals(db)).Methods("GET")
	router.HandleFunc("/meals/{id}",controller1.GetMeal(db)).Methods("GET")
	router.HandleFunc("/mealsadmin",controller1.GetMealsAdmin(db)).Methods("GET")
	router.HandleFunc("/mealsadmin/{id}",controller1.GetMealAdmin(db)).Methods("GET")
	router.HandleFunc("/catsadmin",controller1.GetAnimalsAdmin(db)).Methods("GET")
	router.HandleFunc("/catsadmin/{id}",controller1.GetAnimalAdmin(db)).Methods("GET")
	router.HandleFunc("/cats",controller1.AddAnimal(db)).Methods("POST")
	router.HandleFunc("/meals",controller1.AddMeal(db)).Methods("POST")
	router.HandleFunc("/cats",controller1.UpdateAnimal(db)).Methods("PUT")
	router.HandleFunc("/cats/{id}",controller1.RemoveAnimal(db)).Methods("DELETE")
	router.HandleFunc("/meals/{id}",controller1.RemoveMeal(db)).Methods("DELETE")

	logRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8000",logRouter))
}


