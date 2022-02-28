package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type info struct {
	id    int
	name  string
	email string
	role  string
}

func Create(db *sql.DB, a string) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS ? (id INT NOT NULL, name VARCHAR(100), email VARCHAR(100), role VARCHAR(100),PRIMARY KEY(id))", a)
	if err != nil {
		return errors.New("error while creating a table")
	}
	return nil
}

func Read(db *sql.DB, i int) (*info, error) {
	var a info

	if i < 1 {
		fmt.Println("ID invalid")
		return nil, errors.New("INVALID ID")
	}

	row, err := db.Query("SELECT * FROM employee WHERE id = ?", i)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	for row.Next() {
		err = row.Scan(&(a.id), &(a.name), &(a.email), &(a.role))
		if err != nil {
			return nil, errors.New("error while scaning")
		}
	}
	return &a, nil
}

func Delete(db *sql.DB, i int) error {
	_, err := db.Exec("DELETE FROM employee WHERE id = ?", i)
	if err != nil {
		return errors.New("errror while deleting")
	}

	return nil
}

func Insert(db *sql.DB, a info) error {
	query := "INSERT INTO employee VALUES(id,name,email,role) VALUES(?,?,?,?)"

	_, err := db.Exec(query, a.id, a.name, a.email, a.role)
	if err != nil {
		// log.Print(err)
		return errors.New("errror while inserting")
	}

	return nil
}

func Update(db *sql.DB, i int, k int, a info) error {

	switch k {
	case 0:
		s := "update employee set name=? where id=?"

		_, err2 := db.Exec(s, a.name, i)
		if err2 != nil {
			return errors.New("name not changed")
		}

	case 1:
		s := "update employee set email=? where id=?"

		_, err2 := db.Exec(s, a.email, i)
		if err2 != nil {
			return errors.New("email not changed")
		}

	case 2:
		s := "update employee set role=? where id=?"

		_, err2 := db.Exec(s, a.role, i)
		if err2 != nil {
			return errors.New("role not changed")
		}

	}
	return nil
}

func main() {

	// fmt.Println("GO MYSQL TUTORIAL")
	//CONNECTION

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Connection done Successfully")

	// ins, e := db.Query("INSERT INTO employee VALUES(1,'shivam','singhal','sde')")

	// if e != nil {
	// 	panic(e.Error())
	// }

	// defer ins.Close()

	// fmt.Println("Inserted value in the database")

	//READ THE DATABASE WITH ID
	// var a info
	a, err := Read(db, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.name, a.id, a.email, a.role)
	//DELETE THE DATABASE WITH ID
	Delete(db, 3)

	//INSERT IN THE TABLE
	// x := info{
	// 	id:    111,
	// 	name:  "Samahss",
	// 	email: "ajhasinsansha@gmail.com",
	// 	role:  "sd1213",
	// }
	// Insert(db, x)

	x := info{
		id:    100,
		name:  "Changed name",
		email: "EMAIL@gmail.com",
		role:  "SDTTTT",
	}
	Update(db, 4, 2, x)
}
