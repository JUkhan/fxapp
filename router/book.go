package router

import (
	models "fxapp/db"
	"fxapp/server"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type bookModule struct {
	*server.Server
}

func newBookModule(server *server.Server) *bookModule {
	book := bookModule{server}
	book.setupRoutes()
	return &book
}
func (b *bookModule) setupRoutes() {
	b.App.Get("/api/v1/book", b.GetBooks)
	b.App.Get("/api/v1/book/:id", b.GetBook)
	b.App.Post("/api/v1/book", b.NewBook)
	b.App.Delete("/api/v1/book/:id", b.DeleteBook)
}
func (b *bookModule) GetBooks(c *fiber.Ctx) error {

	var books []models.Book
	b.DB.Find(&books)
	return c.JSON(books)
}

func (b *bookModule) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	b.DB.Find(&book, id)
	return c.JSON(book)
}

func (b *bookModule) NewBook(c *fiber.Ctx) error {

	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	b.DB.Create(&book)
	return c.JSON(book)
}

func (b *bookModule) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book
	b.DB.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No Book Found with ID")
	}
	b.DB.Delete(&book)
	return c.SendString("Book Successfully deleted")
}

var BookModule = fx.Options(
	fx.Invoke(newBookModule),
)
