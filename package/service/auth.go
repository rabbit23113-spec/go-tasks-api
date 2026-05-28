package service

import (
	"log"
	"main/package/database"
	"main/package/types"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Authentication interface {
	SignUp(form types.SignUpDto) error
	SignIn(form types.SignInDto) error
}

type AuthService struct {
	Db *sqlx.DB
}

func (a *AuthService) SignUp(form types.SignUpDto) error {
	db := database.NewPostgresDb{}
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("Db connection error: %v", err)
	}
	a.Db = conn
	hash, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return err
	}
	_, err = a.Db.Exec(`
		INSERT INTO users (firstName, lastName, email, password)
		VALUES ($1, $2, $3, $4);
	`, form.FirstName, form.LastName, form.Email, string(hash))
	if err != nil {
		log.Fatalf("Error: %v", err)
		return err
	}
	return nil
}

func (a *AuthService) SignIn(form types.SignInDto) error {
	var user []types.User
	db := database.NewPostgresDb{}
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("Db connection error: %v", err)
		return err
	}
	a.Db = conn
	err = a.Db.Select(&user, "SELECT * FROM users WHERE email = $1", form.Email)
	if err != nil {
		log.Fatalf("Db connection error: %v", err)
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(form.Password))
	if err != nil {
		log.Fatalf("Db connection error: %v", err)
		return err
	}
	return nil
}
