package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type CategoryDb struct {
	db *sql.DB
}

type Category struct {
	ID          string
	Name        string
	Description string
}

func NewCategoryDb(db *sql.DB) *CategoryDb {
	return &CategoryDb{db: db}
}

func (c *CategoryDb) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(
		"INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id,
		name,
		description,
	)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *CategoryDb) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}
	return categories, nil
}

func (c *CategoryDb) Find(id string) (Category, error) {
	rows, err := c.db.Query(
		"SELECT id, name, description FROM categories WHERE id = $1",
		id,
	)
	if err != nil {
		return Category{}, err
	}
	defer rows.Close()

	var category Category
	for rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return Category{}, err
		}
	}
	return category, nil
}

func (c *CategoryDb) FindByCourseId(courseId string) (Category, error) {
	rows, err := c.db.Query(
		"SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1",
		courseId,
	)
	if err != nil {
		return Category{}, err
	}
	defer rows.Close()

	var category Category
	for rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return Category{}, err
		}
	}
	return category, nil
}
