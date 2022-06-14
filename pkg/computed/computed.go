package computed

import (
	"fmt"
	"sseexample/pkg/notisystem"
	"time"

	"github.com/gofiber/fiber/v2"
)

func computed(request fiber.Map) {
	fmt.Printf("Processing: %+v\n", request)
	time.Sleep(5 * time.Second)

	response := fmt.Sprintf("Completed: %v", request["transaction_id"])
	fmt.Printf("%v\n", response)
	notisystem.NotifierController.Append(response)
}
