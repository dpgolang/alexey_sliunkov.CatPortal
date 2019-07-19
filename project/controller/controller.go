package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"project/model"
	"project/repository"
	"strconv"
)

type Controller struct{}

var cats []model.Animal

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetAnimals(db *sql.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		var cat model.Animal
		cats = []model.Animal{}
		catrepo := repository.CatRepository{}

		cats = catrepo.GetAnimals(db,cat,cats)

		json.NewEncoder(w).Encode(cats)
	}
}

func (c Controller) GetAnimal(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var cat model.Animal
		params := mux.Vars(r)
		cats = []model.Animal{}
		catrepo := repository.CatRepository{}

		id,err:= strconv.Atoi(params["id"])
		logFatal(err)

		cat = catrepo.GetAnimal(db,cat,id)
		json.NewEncoder(w).Encode(cat)
	}
}

func (c Controller) AddAnimal(db *sql.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		var cat model.Animal
		var catID int
		json.NewDecoder(r.Body).Decode(&cat)

		catrepo := repository.CatRepository{}
		catID = catrepo.AddAnimal(db,cat)
		json.NewEncoder(w).Encode(catID)
	}
}

func (c Controller) UpdateAnimal(db *sql.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		var cat model.Animal
		json.NewDecoder(r.Body).Decode(&cat)

		catrepo:= repository.CatRepository{}
		rowsUpdated := catrepo.UpdateAnimal(db,cat)
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveAnimal(db *sql.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		params := mux.Vars(r)

		//result,err := db.Exec("delete from public.cats where id = $1",params["id"])
		//logFatal(err)
		catrepo:= repository.CatRepository{}

		id,err := strconv.Atoi(params["id"])
		logFatal(err)
		rowsDeleted := catrepo.RemoveAnimal(db,id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}




