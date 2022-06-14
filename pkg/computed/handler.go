package computed

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx) error {
	fmt.Println("Called Computed Handler")
	var request fiber.Map

	err := c.BodyParser(&request)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	go computed(request)

	return c.SendStatus(fiber.StatusAccepted)
}
