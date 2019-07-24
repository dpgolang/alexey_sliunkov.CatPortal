package repository

import (
	"database/sql"
	"log"
	"project/model"
	"github.com/jmoiron/sqlx"
)

type CatRepository struct {}

type UserRepository struct{}

type FoodRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (ca CatRepository) GetAnimals(db *sqlx.DB,cat model.Animal,cats [] model.Animal) []model.Animal{
	rows,err := db.Query("select id,breed,size,diet,motherland,description from public.cats")
	logFatal(err)

	defer rows.Close()
	err = sqlx.StructScan(rows, &cats)
	if err != nil {
		return []model.Animal{}
	}

	return cats
}

func (ca CatRepository) GetAnimal(db *sqlx.DB,cat model.Animal,id int) model.Animal{
	rows := db.QueryRow("select id,breed,size,diet,motherland,description from public.cats where id=$1",id)
	err :=rows.Scan(&cat.ID,&cat.Breed,&cat.Size,&cat.Diet,&cat.Motherland,&cat.Description)
	logFatal(err)

	return cat
}

func (ca CatRepository) AddAnimal(db *sqlx.DB,cat model.Animal) int{
	err:= db.QueryRow("insert into public.cats (id,breed,size,diet,motherland,description) values ($1, $2, $3, $4, $5,$6) returning id;",
		cat.ID,cat.Breed,cat.Size,cat.Diet,cat.Motherland,cat.Description).Scan(&cat.ID)

	logFatal(err)

	return cat.ID
}

func (ca CatRepository) UpdateAnimal(db *sqlx.DB,cat model.Animal) int64{
	result,err := db.Exec("update public.cats breed=$1, size=$2, diet=$3,motherland=$4,description=$5 where id=$6 returning id",
		cat.Breed,cat.Size,cat.Diet,cat.Motherland,cat.Description,cat.ID)
	logFatal(err)

	rowsUpdated,err :=result.RowsAffected()
	logFatal(err)

	return rowsUpdated
}

func (ca CatRepository) RemoveAnimal(db *sqlx.DB,id int) int64{
	result,err := db.Exec("delete from public.cats where id = $1",id)
	logFatal(err)

	rowsDeleted,err:= result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}

func (u UserRepository) Signup(db *sqlx.DB, user model.User) int {
	err := db.QueryRow("insert into public.users (id,firstname,lastname,password) values ($1,$2,$3,$4) RETURNING id;",
		user.Id, user.Firstname, user.Lastname, user.Password).Scan(&user.Id)
	logFatal(err)
	return user.Id
}
func (u UserRepository) Signin(db *sqlx.DB, userСhecking model.User, userFromBase model.User) (string, bool) {
	err := db.QueryRow("select password from public.users where id=$1", userСhecking.Id).Scan(&userFromBase.Password)
	if err == sql.ErrNoRows {
		return "", false
	}
	logFatal(err)
	return userFromBase.Password, true
}

func (foo FoodRepository) GetMeals(db *sqlx.DB,meal model.Food,meals [] model.Food)[] model.Food{
	rows,err := db.Query("select * from public.food")
	logFatal(err)

	defer rows.Close()

	for rows.Next(){
		err :=rows.Scan(&meal.Id,&meal.Name,&meal.Price,&meal.Composition)
		logFatal(err)

		meals = append(meals,meal)
	}
	return meals
}

func (foo FoodRepository) GetMeal(db *sqlx.DB,meal model.Food,id int) model.Food{
	rows := db.QueryRow("select * from public.food where id=$1",id)
	err :=rows.Scan(&meal.Id,&meal.Name,&meal.Price,&meal.Composition)
	logFatal(err)

	return meal
}

func (foo FoodRepository) AddMeal(db *sqlx.DB,meal model.Food) int{
	err:= db.QueryRow("insert into public.food (id,name,price,composition) values ($1, $2, $3, $4) returning id;",
		meal.Id,meal.Name,meal.Price,meal.Composition).Scan(&meal.Id)

	logFatal(err)

	return meal.Id
}

func (foo FoodRepository) RemoveMeal(db *sqlx.DB,id int) int64{
	result,err := db.Exec("delete from public.food where id = $1",id)
	logFatal(err)

	rowsDeleted,err:= result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}
