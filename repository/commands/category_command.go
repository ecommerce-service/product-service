package commands

import (
	"github.com/ecommerce-service/product-service/domain/commands"
	"github.com/ecommerce-service/product-service/domain/models"
	"github.com/thel5coder/pkg/postgresql"
)

type CategoryCommand struct {
	db    postgresql.IConnection
	model *models.Category
}

func NewCategoryCommand(db postgresql.IConnection, model *models.Category) commands.ICategoryCommand {
	return &CategoryCommand{
		db:    db,
		model: model,
	}
}

func (c CategoryCommand) Add() (res string, err error) {
	statement := `INSERT INTO categories (name,slug,created_at,updated_at) VALUES ($1,$2,$3,$4) RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.Name(), c.model.Slug(), c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CategoryCommand) Edit() (res string, err error) {
	statement := `UPDATE categories SET name=$1,slug=$2,updated_at=$3 WHERE id=$4 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.Name(), c.model.Slug(), c.model.UpdatedAt(), c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CategoryCommand) Delete() (res string, err error) {
	statement := `UPDATE categories SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
