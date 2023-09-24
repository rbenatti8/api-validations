package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rbenatti8/api-validations/internal/api/binder"
	"github.com/rbenatti8/api-validations/internal/api/handler"
	"github.com/rbenatti8/api-validations/internal/api/midleware"
	"github.com/rbenatti8/api-validations/internal/api/validator"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()

	e.Binder = binder.New(&echo.DefaultBinder{})
	e.Validator = validator.New()
	e.Use(midleware.ErrorHandler())
	
	userHandler := handler.UserHandler{}

	e.POST("/user", userHandler.Create)

	e.Logger.Fatal(e.Start(":1323"))
}

type pathParams struct {
	params map[string]string
}

func (p pathParams) Get(k string) (string, bool) {
	v, ok := p.params[k]
	return v, ok
}

func buildPathParams(c echo.Context) pathParams {
	params := make(map[string]string)
	for _, p := range c.ParamNames() {
		params[p] = c.Param(p)
	}

	return pathParams{
		params: params,
	}
}
