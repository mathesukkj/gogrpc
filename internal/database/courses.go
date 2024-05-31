package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type CourseDb struct {
	db *sql.DB
}

type Course struct {
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourseDb(db *sql.DB) *CourseDb {
	return &CourseDb{db: db}
}

func (c *CourseDb) Create(name, description, categoryId string) (Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(
		"INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)",
		id,
		name,
		description,
		categoryId,
	)
	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description, CategoryId: categoryId}, nil
}

func (c *CourseDb) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}
		courses = append(
			courses,
			Course{ID: id, Name: name, Description: description, CategoryId: categoryId},
		)
	}
	return courses, nil
}

func (c *CourseDb) FindByCategoryId(categoryId string) ([]Course, error) {
	rows, err := c.db.Query(
		"SELECT id, name, description, category_id FROM courses WHERE category_id = $1",
		categoryId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}
		courses = append(
			courses,
			Course{ID: id, Name: name, Description: description, CategoryId: categoryId},
		)
	}
	return courses, nil
}
