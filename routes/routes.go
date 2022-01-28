package routes

import (
	"github.com/Chaksack/staples-backend/controllers"
	"github.com/Chaksack/staples-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//controllers endpoints
	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)

	app.Use(middleware.Isauthenticated)

	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)

	app.Get("api/users", controllers.AllUsers)
	app.Post("api/users", controllers.CreateUser)
	app.Get("api/users/:id", controllers.GetUser)
	app.Put("api/users/:id", controllers.UpdateUser)
	app.Delete("api/users/:id", controllers.DeleteUser)

	app.Get("api/roles", controllers.AllRoles)
	app.Post("api/roles", controllers.CreateRole)
	app.Get("api/roles/:id", controllers.GetRole)
	app.Put("api/roles/:id", controllers.UpdateRole)
	app.Delete("api/roles/:id", controllers.DeleteRole)

	app.Get("api/permissions", controllers.AllPermissions)
	app.Post("api/permissions", controllers.CreatePermission)
	app.Get("api/permissions/:id", controllers.GetPermission)
	app.Put("api/permissions/:id", controllers.UpdatePermission)
	app.Delete("api/permissions/:id", controllers.DeletePermission)
}
