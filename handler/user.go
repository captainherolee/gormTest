package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUser(c echo.Context) error {

	// #1 Get user in DB
	email := fmt.Sprintf("%v", c.QueryParam("email"))
	if len(email) == 0 {
		return ErrEmailEmpty
	}

	user, err := h.db.GetUser(email)
	if err != nil {
		return ErrInternalServer
	}

	if len(user.Email) == 0 {
		return ErrEmailNotFoundInDB
	}

	retUser := JSONUser{
		Email:        user.Email,
		Name:         user.Name,
		Organization: user.Organization,
		Tag:          user.Tag,
	}

	return c.JSON(http.StatusOK, retUser)
}
