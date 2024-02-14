package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
  db *sql.DB
  ID string
  Name string
  Description string
  CategoryID string
}

func NewCourse(db *sql.DB) *Course {
  return &Course{db: db}
}

func (c *Course) Create(name, description, categoryID string) (Course, error) {
  newId := uuid.New().String()

  _, err := c.db.Exec(
    "INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)",
    newId, name, description, categoryID,
  )
  if err != nil {
    return Course{}, err
  }

  return Course{ID: newId, Name: name, Description: description, CategoryID: categoryID}, nil
}

func (c *Course) GetAll() ([]Course, error) {
  rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  courses := []Course{}

  for rows.Next() {
    var course Course

    err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
    if err != nil {
      return nil, err
    }

    courses = append(courses, course)
  }

  return courses, nil
}

func (c *Course) GetByCategoryID(categoryID string) ([]Course, error) {
  rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryID)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  courses := []Course{}

  for rows.Next() {
    var course Course

    err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
    if err != nil {
      return nil, err
    }

    courses = append(courses, course)
  }

  return courses, nil
}
