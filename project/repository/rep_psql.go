package repository

import (
	"database/sql"
	"log"
	"project/model"
)

type CatRepository struct {}

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (ca CatRepository) GetAnimals(db *sql.DB,cat model.Animal,cats [] model.Animal)[] model.Animal{
	rows,err := db.Query("select * from public.cats")
	logFatal(err)

	defer rows.Close()

	for rows.Next(){
		err :=rows.Scan(&cat.ID,&cat.Breed,&cat.Size,&cat.Diet,&cat.Motherland,&cat.Description)
		logFatal(err)

		cats = append(cats,cat)
	}
	return cats
}

func (ca CatRepository) GetAnimal(db *sql.DB,cat model.Animal,id int) model.Animal{
	rows := db.QueryRow("select * from public.cats where id=$1",id)
	err :=rows.Scan(&cat.ID,&cat.Breed,&cat.Size,&cat.Diet,&cat.Motherland,&cat.Description)
	logFatal(err)

	return cat
}

func (ca CatRepository) AddAnimal(db *sql.DB,cat model.Animal) int{
	err:= db.QueryRow("insert into public.cats (id,breed,size,diet,motherland,description) values ($1, $2, $3, $4, $5,$6) returning id;",
		cat.ID,cat.Breed,cat.Size,cat.Diet,cat.Motherland,cat.Description).Scan(&cat.ID)

	logFatal(err)

	return cat.ID
}

func (ca CatRepository) UpdateAnimal(db *sql.DB,cat model.Animal) int64{
	result,err := db.Exec("update public.cats breed=$1, size=$2, diet=$3,motherland=$4,description=$5 where id=$6 returning id",
		cat.Breed,cat.Size,cat.Diet,cat.Motherland,cat.Description,cat.ID)
	logFatal(err)

	rowsUpdated,err :=result.RowsAffected()
	logFatal(err)

	return rowsUpdated
}

func (ca CatRepository) RemoveAnimal(db *sql.DB,id int) int64{
	result,err := db.Exec("delete from public.cats where id = $1",id)
	logFatal(err)

	rowsDeleted,err:= result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}

func (u UserRepository) Signup(db *sql.DB, user model.User) int {
	err := db.QueryRow("insert into public.users (id,firstname,lastname,password) values ($1,$2,$3,$4) RETURNING id;",
		user.Id, user.Firstname, user.Lastname, user.Password).Scan(&user.Id)
	logFatal(err)
	return user.Id
}
func (u UserRepository) Signin(db *sql.DB, userСhecking model.User, userFromBase model.User) (string, bool) {
	err := db.QueryRow("select password from public.users where id=$1", userСhecking.Id).Scan(&userFromBase.Password)
	if err == sql.ErrNoRows {
		return "", false
	}
	logFatal(err)
	return userFromBase.Password, true
}
