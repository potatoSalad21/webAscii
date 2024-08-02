package handlers

import (
	"github.com/labstack/echo/v4"
	base "github.com/potatoSalad21/webAscii/views/home"
)

func HandleHome(c echo.Context) error {
	return render(c, base.Show())
}
