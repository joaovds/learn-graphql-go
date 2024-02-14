package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
  db *sql.DB
  ID string
  Name string
  Description string
}

func NewCategory(db *sql.DB) *Category {
  return &Category{db: db}
}

func (c *Category) Create(name, description string) (Category, error) {
  newId := uuid.New().String()

  _, err := c.db.Exec(
    "INSERT INTO category (id, name, description) VALUES ($1, $2, $3)",
    newId, name, description,
  )
  if err != nil {
    return Category{}, err
  }

  return Category{ID: newId, Name: name, Description: description}, nil
}
