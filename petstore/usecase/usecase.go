package usecase

import "context"

type Pets []Pet

type Pet struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Tags     []Tag  `json:"tags"`
}

type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// CreatePets creates a new pet record
func CreatePets(ctx context.Context, pet Pet) error {
	return nil
}

// ListPets lists all pets
func ListPets(ctx context.Context) (Pets, error) {
	pets := make(Pets, 0)
	dog := Pet{
		ID:       0,
		Name:     "dog",
		Category: "dog",
	}
	pets = append(pets, dog)
	return pets, nil
}
