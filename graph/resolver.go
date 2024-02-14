package graph

import "github.com/joaovds/learn-graphql-go/internal/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
  CategoryDB *db.Category
  CourseDB *db.Course
}
