package route

import(
	"github.com/gofiber/fiber/v2"
	"restapi-gorm-gofiber/handler"
)


func RouteInit(r *fiber.App) {
	r.Get ("/user", handler.UserHandlerGetAll)
	r.Post ("/user", handler.UserHandlerCreate)

}
