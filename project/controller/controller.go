package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"project/model"
	"project/repository"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/sessions"
)

type Controller struct{}

var cats []model.Animal

var goods []model.Food

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func (c Controller) GetAnimals(db *sqlx.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		var cat model.Animal
		cats = []model.Animal{}
		catrepo := repository.CatRepository{}

		cats = catrepo.GetAnimals(db,cat,cats)

		json.NewEncoder(w).Encode(cats)
	}
}

func (c Controller) GetAnimal(db *sqlx.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
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

func (c Controller) AddAnimal(db *sqlx.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["admin"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		var cat model.Animal
		var catID int
		json.NewDecoder(r.Body).Decode(&cat)

		catrepo := repository.CatRepository{}
		catID = catrepo.AddAnimal(db,cat)
		json.NewEncoder(w).Encode(catID)
	}
}

func (c Controller) UpdateAnimal(db *sqlx.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["admin"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		var cat model.Animal
		json.NewDecoder(r.Body).Decode(&cat)

		catrepo:= repository.CatRepository{}
		rowsUpdated := catrepo.UpdateAnimal(db,cat)
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveAnimal(db *sqlx.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["admin"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		params := mux.Vars(r)

		catrepo:= repository.CatRepository{}

		id,err := strconv.Atoi(params["id"])
		logFatal(err)
		rowsDeleted := catrepo.RemoveAnimal(db,id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}

func (c Controller) Signup(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		var userID int
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&user)
		logFatal(err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		user.Password = string(hashedPassword)
		userRepo := repository.UserRepository{}
		userID = userRepo.Signup(db, user)
		json.NewEncoder(w).Encode(&userID)
	}
}

func (c Controller) Signin(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		session.Options.MaxAge = 300
		w.Header().Set("Content-Type", "application/json")
		userFromBase := model.User{}
		userChecking := model.User{}
		var userIsFound bool
		err := json.NewDecoder(r.Body).Decode(&userChecking)
		logFatal(err)
		userRepo := repository.UserRepository{}
		userFromBase.Password, userIsFound = userRepo.Signin(db, userChecking, userFromBase)
		if userIsFound {

			if err = bcrypt.CompareHashAndPassword([]byte(userFromBase.Password), []byte(userChecking.Password)); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			session.Values["authenticated"] = true
			session.Save(r, w)
			json.NewEncoder(w).Encode(fmt.Sprintf("welcome user, id: %d",userChecking.Id))
		}
	}
}

func (c Controller) SignAdmin(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		session.Options.MaxAge = 300
		w.Header().Set("Content-Type", "application/json")
		userFromBase := model.User{}
		adminChecking := model.User{0,"admin","admin","admin1111"}
		var adminIsFound bool
		err := json.NewDecoder(r.Body).Decode(&adminChecking)
		logFatal(err)
		userRepo := repository.UserRepository{}
		userFromBase.Password, adminIsFound = userRepo.Signin(db, adminChecking, userFromBase)
		if adminIsFound {

			if err = bcrypt.CompareHashAndPassword([]byte(userFromBase.Password), []byte(adminChecking.Password)); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				//json.NewEncoder(w).Encode()
				return
			}
			session.Values["admin"] = true
			session.Save(r, w)
			json.NewEncoder(w).Encode(fmt.Sprintf("welcome admin, id: %d",adminChecking.Id))
		}
	}
}

func (c Controller) GetAnimalsAdmin(db *sqlx.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["admin"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		var cat model.Animal
		cats = []model.Animal{}
		catrepo := repository.CatRepository{}

		cats = catrepo.GetAnimals(db,cat,cats)

		json.NewEncoder(w).Encode(cats)
	}
}

func (c Controller) GetAnimalAdmin(db *sqlx.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["admin"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
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

func (c Controller) GetMeals(db *sqlx.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		var food model.Food
		goods = []model.Food{}
		foodrepo := repository.FoodRepository{}

		goods = foodrepo.GetMeals(db,food,goods)

		json.NewEncoder(w).Encode(goods)
	}
}

func (c Controller) GetMeal(db *sqlx.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		var food model.Food
		params := mux.Vars(r)
		goods = []model.Food{}
		foodrepo := repository.FoodRepository{}

		id,err:= strconv.Atoi(params["id"])
		logFatal(err)

		food = foodrepo.GetMeal(db,food,id)
		json.NewEncoder(w).Encode(food)
	}
}

func (c Controller) AddMeal(db *sqlx.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["admin"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		var food model.Food
		var mealID int
		json.NewDecoder(r.Body).Decode(&food)

		foodrepo := repository.FoodRepository{}
		mealID = foodrepo.AddMeal(db,food)
		json.NewEncoder(w).Encode(mealID)
	}
}

func (c Controller) RemoveMeal(db *sqlx.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["admin"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		params := mux.Vars(r)

		foodrepo := repository.FoodRepository{}

		id,err := strconv.Atoi(params["id"])
		logFatal(err)
		rowsDeleted := foodrepo.RemoveMeal(db,id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}

func (c Controller) GetMealsAdmin(db *sqlx.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")

		if auth, ok := session.Values["admin"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		var food model.Food
		goods = []model.Food{}
		foodrepo := repository.FoodRepository{}

		goods = foodrepo.GetMeals(db,food,goods)

		json.NewEncoder(w).Encode(goods)
	}
}

func (c Controller) GetMealAdmin(db *sqlx.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["admin"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		var food model.Food
		params := mux.Vars(r)
		goods = []model.Food{}
		foodrepo := repository.FoodRepository{}

		id,err:= strconv.Atoi(params["id"])
		logFatal(err)

		food = foodrepo.GetMeal(db,food,id)
		json.NewEncoder(w).Encode(food)
	}
}



