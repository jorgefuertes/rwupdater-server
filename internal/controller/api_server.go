package controller

import (
	"fmt"

	"git.martianoids.com/queru/retroserver/internal/build"

	"github.com/gofiber/fiber/v2"
)

func APIServer(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf(
		"+ SERVER VERSION:\n\n- %s\n- %s\n- %s\n- %s\n",
		build.Version(),
		build.VersionShort(),
		build.BinVersion(),
		build.CompileTime(),
	))
}
