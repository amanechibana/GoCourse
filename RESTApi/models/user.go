package models

import (
	"errors"

	"example.com/restapi/db"
	"example.com/restapi/utils"
)

type User struct {
	ID 		 int64 
	Email 	 string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error{
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil{
		return err
	}

	res, err := stmt.Exec(u.Email,hashedPassword)

	if err != nil{
		return err
	}

	id, err := res.LastInsertId()
	u.ID = id

	return err
}

func (u User) Validate() error{
	query := "SELECT id, password FROM users where email = ?"
	row := db.DB.QueryRow(query,u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	if !utils.CheckPass(u.Password,retrievedPassword) {
		return errors.New("credentials invalid")
	}

	return nil
}