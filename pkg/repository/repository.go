package repository

import (
	"github.com/LevOrlov5404/task-tracker/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int64, error)
}

type Project interface {

}

type Task interface {

}

type Subtask interface {

}

type Repository struct {
	Authorization
	Project
	Task
	Subtask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}