package service

import (
	"github.com/jmoiron/sqlx"
)

type Tasks interface {
	Action() error
}

type TasksService struct {
	Db *sqlx.DB
}

func (t *TasksService) Action() error {
	return nil
}
