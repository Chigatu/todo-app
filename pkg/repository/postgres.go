package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
	todoListTable = "todo_lists"
	usersListsTable = "users_list"
	todoItemTable = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct{
	Host string
	Port string
	Username string
	Password string
	DBname string
	SSLMode string

}

func NewPostgresDb(cfg Config) (*sqlx.DB, error){
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}