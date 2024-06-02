package main

import (
	"fmt"
	"os"
	"ulascan-be/config"
	"ulascan-be/controller"
	"ulascan-be/database"
	"ulascan-be/middleware"
	"ulascan-be/repository"
	"ulascan-be/routes"
	"ulascan-be/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("STARTING...")

	var (
		// DATABASE
		db *gorm.DB = config.SetupDatabaseConnection()

		// REPOSITORY
		userRepository repository.UserRepository = repository.NewUserRepository(db)

		// SERVICE
		jwtService  service.JWTService  = service.NewJWTService()
		userService service.UserService = service.NewUserService(userRepository, jwtService)

		// CONTROLLER
		userController controller.UserController = controller.NewUserController(userService)
	)

	defer config.CloseDatabaseConnection(db)

	fmt.Println("MIGRATING DATABASE...")
	if err := database.MigrateFresh(db); err != nil {
		panic(err)
	}
	fmt.Println("> Database Migrated")

	fmt.Println("SEEDING DATABASE...")
	if err := database.Seeder(db); err != nil {
		panic(err)
	}
	fmt.Println("> Database Seeded")

	// SERVER
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// ROUTES
	routes.User(server, userController, jwtService)

	// RUNING THE SERVER
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := server.Run(":" + port); err != nil {
		fmt.Println("Server failed to start: ", err)
		return
	}
}
