package midleware

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type httpError interface {
	error
	Code() int
}

//	var mapError = map[error]int{
//		binder.BindError: 400,
//	}
func ErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				return handleErrors(c, err)
			}

			return nil
		}
	}
}

func handleErrors(c echo.Context, err error) error {
	var he httpError

	if errors.As(err, &he) {
		return c.JSON(he.Code(), he)
	}

	return c.JSON(500, "internal server error")
}
