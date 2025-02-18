package database

import (
	"context"
	"grapql-to-do/internal/domain/model"
	"grapql-to-do/internal/infrastructure"
)

type UserDAO struct {
	DB *infrastructure.Database
}

func NewUserCreateDAO(db *infrastructure.Database) *UserDAO {
	return &UserDAO{DB: db}
}

func (dao *UserDAO) Create(ctx context.Context, user *model.User) (int32, error) {
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	var id int32
	err := dao.DB.Conn.QueryRow(ctx, query, user.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
