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
    "INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
    newId, name, description,
  )
  if err != nil {
    return Category{}, err
  }

  return Category{ID: newId, Name: name, Description: description}, nil
}

func (c *Category) GetAll() ([]Category, error) {
  rows, err := c.db.Query("SELECT id, name, description FROM categories")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  categories := []Category{}

  for rows.Next() {
    var category Category

    err := rows.Scan(&category.ID, &category.Name, &category.Description)
    if err != nil {
      return nil, err
    }

    categories = append(categories, category)
  }

  return categories, nil
}

func (c *Category) GetByCourseID(courseID string) (Category, error) {
  var category Category

  err := c.db.QueryRow(
    "SELECT cate.id, cate.name, cate.description FROM categories cate JOIN courses cour ON cate.id = cour.category_id WHERE cour.id = $1",
    courseID,
  ).Scan(&category.ID, &category.Name, &category.Description)
  if err != nil {
    return Category{}, err
  }

  return category, nil
}
