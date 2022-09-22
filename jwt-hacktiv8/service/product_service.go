package service

import (
	"context"
	"jwt-hacktiv8/dto"
	"jwt-hacktiv8/entity"
	"jwt-hacktiv8/repository"

	"github.com/mashingan/smapping"
)

type ProductService interface {
	GetProductByID(ctx context.Context, id uint) (entity.Product, error)
	InsertProduct(ctx context.Context, productDTO dto.ProductCreateDTO) (entity.Product, error)
	UpdateProduct(ctx context.Context, productDTO dto.ProductUpdateDTO) (entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(pr repository.ProductRepository) ProductService {
	return &productService{
		productRepository: pr,
	}
}

func (s *productService) InsertProduct(ctx context.Context, productDTO dto.ProductCreateDTO) (entity.Product, error) {
	var createdProduct entity.Product

	err := smapping.FillStruct(&createdProduct, smapping.MapFields(&productDTO))
	if err != nil {
		return entity.Product{}, err
	}
	return s.productRepository.InsertProduct(ctx, createdProduct)
}

func (s *productService) UpdateProduct(ctx context.Context, productDTO dto.ProductUpdateDTO) (entity.Product, error) {
	var createdProduct entity.Product

	err := smapping.FillStruct(&createdProduct, smapping.MapFields(&productDTO))
	createdProduct.ID = productDTO.ID
	if err != nil {
		return entity.Product{}, err
	}
	return s.productRepository.UpdateProduct(ctx, createdProduct)
}

func (s *productService) GetProductByID(ctx context.Context, id uint) (entity.Product, error) {
	return s.productRepository.GetProductByID(ctx, id)
}
