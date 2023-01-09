package router

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"gormTest/handler"
	md "gormTest/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() (*echo.Echo, error) {
	e := echo.New()
	e.HTTPErrorHandler = handler.NewHttpErrorHandler(handler.NewErrorStatusCodeMaps()).ErrorHandler

	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(md.Cors)

	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	userGroup := e.Group("api/v1/user")
	{
		userGroup.GET("", h.GetUser)
	}

	return e, nil
}

func StartServer() {
	defer func() {
		recover()
		log.Println("Recovery")
		retryDurationSecond, getEnvErr := strconv.Atoi("10")
		if getEnvErr != nil {
			log.Println(getEnvErr)
			panic(getEnvErr)
		}

		log.Println("Server Start Error -> Retrying after", fmt.Sprintf("%v", retryDurationSecond), "Seconds...")
		time.Sleep(time.Second * time.Duration(retryDurationSecond))

		go StartServer()
	}()

	server, err := Router()
	if err == nil {
		fmt.Println(" Server Start ")
		address := "0.0.0.0:8088"
		if err := server.Start(address); err != http.ErrServerClosed {
			panic(err)
		}
		/*if err := server.StartTLS(address, "cert.pem", "key.pem"); err != http.ErrServerClosed {
			log.Fatal(err)
		}*/
	} else {
		panic(err)
	}
}
