package api

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/tetsurowada/backend-golang-echo/petstore/usecase"
)

type PetsService struct {
	pets map[int64]Pet
	id   int64
}

func NewPetsService() PetsService {
	return PetsService{
		pets: make(map[int64]Pet),
	}
}

func (p *PetsService) CreatePets(ctx context.Context, req *Pet) error {
	var category string
	switch req.ID {
	case 0:
		category = "dog"
	case 1:
		category = "cat"
	default:
		category = "other"
	}
	pet := usecase.Pet{
		ID:       req.ID,
		Name:     req.Name,
		Category: category,
	}

	err := usecase.CreatePets(ctx, pet)
	if err != nil {
		return fmt.Errorf("failed to create pet: %w", err)
	}
	return nil
}

func (p *PetsService) ListPets(ctx context.Context, params ListPetsParams) (*PetsHeaders, error) {
	pets, err := usecase.ListPets(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list pets: %w", err)
	}
	result := make([]Pet, 0, len(pets))
	for _, pet := range pets {
		slog.Info(fmt.Sprintf("pet: %s", pet.Name))
		obj := Pet{
			ID:   pet.ID,
			Name: pet.Name,
			Tag:  OptString{Value: pet.Category},
		}
		result = append(result, obj)
	}

	petsHeaders := PetsHeaders{
		XNext:    OptString{Value: "next"},
		Response: result,
	}
	return &petsHeaders, nil
}

func (p *PetsService) ShowPetById(ctx context.Context, params ShowPetByIdParams) (*Pet, error) {
	return nil, nil
}

func (p *PetsService) NewError(ctx context.Context, err error) *ErrorStatusCode {
	return &ErrorStatusCode{
		StatusCode: 500,
		Response: Error{
			Code:    500,
			Message: "Internal Server Error",
		},
	}
}
