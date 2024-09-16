package route

import(
	"github.com/gofiber/fiber/v2"
	"restapi-gorm-gofiber/handler"
)


func RouteInit(r *fiber.App) {
	r.Get ("/user", handler.UserHandlerGetAll)
	r.Get ("/user/:id", handler.UserHandlerGetById)
	r.Post ("/user", handler.UserHandlerCreate)
	r.Put ("/user/:id", handler.UserHandlerUpdate)
	r.Put ("/user/:id/updateEmail", handler.UserHandlerUpdateEmail)
	r.Delete ("/user/:id", handler.UserHandlerDelete)

}
