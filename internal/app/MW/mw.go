package MW

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		userRole := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(userRole, roleAdmin) {
			log.Println("Admin role detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
