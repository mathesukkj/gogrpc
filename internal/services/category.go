package services

import (
	"context"

	"github.com/mathesukkj/gogrpc/internal/database"
	"github.com/mathesukkj/gogrpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.CategoryDb
}

func NewCategoryService(categoryDB database.CategoryDb) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(
	ctx context.Context,
	in *pb.CreateCategoryRequest,
) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}
	return categoryResponse, nil
}
