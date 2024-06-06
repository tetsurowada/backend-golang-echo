package main

import (
	"context"

	"github.com/labstack/echo/v4"
	petstore "github.com/tetsurowada/backend-golang-echo/api"
)

func main() {
	e := echo.New()
	h := NewHandler(context.Background())
	petstore.RegisterHandlers(e, h)
	e.Logger.Fatal(e.Start(":8080"))
}

type Handler struct {
	c context.Context
}

func NewHandler(c context.Context) *Handler {
	return &Handler{c: c}
}

func (h *Handler) ListPets(ctx echo.Context, params petstore.ListPetsParams) error {
	return nil
}

func (h *Handler) CreatePets(ctx echo.Context) error {
	return nil
}

func (h *Handler) ShowPetById(ctx echo.Context, petID string) error {
	return nil
}
