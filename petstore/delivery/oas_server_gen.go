// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreatePets implements createPets operation.
	//
	// Create a pet.
	//
	// POST /pets
	CreatePets(ctx context.Context, req *Pet) error
	// ListPets implements listPets operation.
	//
	// List all pets.
	//
	// GET /pets
	ListPets(ctx context.Context, params ListPetsParams) (*PetsHeaders, error)
	// ShowPetById implements showPetById operation.
	//
	// Info for a specific pet.
	//
	// GET /pets/{petId}
	ShowPetById(ctx context.Context, params ShowPetByIdParams) (*Pet, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
