package main

import (
	// database/sql is an interface
	"database/sql"
	
	// go-sql-drive implements the sql interface
	// the underscore means it is an anonymous import
	_ "github.com/go-sql-driver/mysql"

	"log"
)

func main() {
	// get a database object
	db, err := sql.Open("mysql", "root:wwg@tcp(127.0.0.1:3306)/job_board_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	// check that you can establish a connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	

	err = getAll(db)
	if err != nil {
		log.Fatal(err)
	}

	// // add another record and get all results again
	// query := `insert into user_roles_t (login, type) VALUES (?, ?)`
	// _, err = db.Exec(query, "testuser2", "user")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = getAll(db)
	// if err != nil {
	// 	log.Fatal(err)

}

func getAll(db *sql.DB) error{
	// get all rows	
	rows, err := db.Query("select * from user_roles_t")
	if err != nil {
		return err
	}
	defer rows.Close()
	
	var login string
	var roleType string
	
	// iterate on the results and print each one
	for rows.Next() {
		err := rows.Scan(&login, &roleType)
		if err != nil {
			return err
		}
		log.Println(login, roleType)
	}
	
	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
