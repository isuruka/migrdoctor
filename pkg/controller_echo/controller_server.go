package controller_echo

import (
	echoswagger "github.com/swaggo/echo-swagger"
	apiecho "migrdoctor/pkg/api_echo"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

// @title migrdoctor API Documentation
// @version 1.0

func CreateDoctor(c echo.Context) error {

	result, _ := apiecho.CreateDoctor(c)
	return c.JSON(http.StatusOK, result)
}

func CreatePatient(c echo.Context) error {

	result, _ := apiecho.CreatePatient(c)
	return c.JSON(http.StatusOK, result)
}

func APIControllerEcho(g *echo.Group) {

	g.POST("/api/doctor", CreateDoctor)
	g.POST("/api/patient", CreatePatient)
	g.GET("/api/swagger/*any", echoswagger.WrapHandler)
}

func APIControllerEchoSecured(g *echo.Group) {

}
