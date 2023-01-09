package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"
)

var Cors = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
})
