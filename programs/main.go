package main

import (
	"encoding/json"

	"fmt"

	"github.com/gofiber/fiber"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func sayHey(c *fiber.Ctx) {
	fmt.Println("Hey!")
	c.Send("Hey!")
}

func sayHello(c *fiber.Ctx) {
	fmt.Println("Hello!")
	c.Send("Hello!")
}

func sayWorld(c *fiber.Ctx) {
	fmt.Println("World!")
	c.Send("World!")
}

func main() {
	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) {
		users := []User{{ID: 1, Name: "John Doe"}, {ID: 2, Name: "Jane Doe"}}

		jsonUsers := make(chan []byte)

		go func() {
			json, _ := json.Marshal(users)
			jsonUsers <- json
		}()

		c.Send(<-jsonUsers)

	})

	app.Get("/helloworld", func(c *fiber.Ctx) {

		hey := make(chan string)
		hello := make(chan string)
		world := make(chan string)

		go func() {
			sayHey(c)
			hey <- "Hey!"
		}()
		go func() {
			sayHello(c)
			hello <- "Hello!"
		}()
		go func() {
			sayWorld(c)
			world <- "World!"
		}()

		c.Send(<-hey + <-hello + <-world)

	})

	app.Listen(3000)
}
