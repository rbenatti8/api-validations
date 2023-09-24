package binder

import "github.com/labstack/echo/v4"

type binder interface {
	Bind(a any, c echo.Context) error
	BindHeaders(c echo.Context, a any) error
}

type EchoBinder struct {
	echoBinder binder
}

func New(binder binder) *EchoBinder {
	return &EchoBinder{
		echoBinder: binder,
	}
}

func (b *EchoBinder) Bind(i interface{}, c echo.Context) error {
	if err := b.echoBinder.Bind(i, c); err != nil {
		return handleErrors(err)
	}

	if err := b.echoBinder.BindHeaders(c, i); err != nil {
		return handleErrors(err)
	}

	return nil
}
