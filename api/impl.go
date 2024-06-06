package petstore

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// optional code omitted

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /pets)
func (Server) ListPets(ctx echo.Context, params ListPetsParams) error {
	resp := Pong{
		Ping: "pong",
	}

	return ctx.JSON(http.StatusOK, resp)
}
