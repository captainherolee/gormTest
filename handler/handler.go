package handler

import (
	"gormTest/db"

	"github.com/labstack/echo/v4"
)

type HandlerInterface interface {
	GetUser(c echo.Context) error
}

type Handler struct {
	db db.DB
}

func NewHandler() (HandlerInterface, error) {
	dbms := "mysql"
	dbConf := "root:1111@tcp(0.0.0.0:3306)/eiffel"
	return NewHandlerWithParams(dbms, dbConf)
}

func NewHandlerWithParams(dbtype, conn string) (HandlerInterface, error) {
	dbInstance, err := db.NewORM(dbtype, conn)
	if err != nil {
		panic(err)
	}

	return &Handler{
		db: dbInstance,
	}, nil
}
