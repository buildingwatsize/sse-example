package sse

import (
	"bufio"
	"sseexample/pkg/notisystem"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func Handler(c *fiber.Ctx) error {

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		notisystem.BrokerList.Subscribe(w)
		notisystem.BrokerList.Listen()
	}))
	return nil
}
