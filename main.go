package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/ryanmoralesaz/colossal-stack/config"
	"github.com/ryanmoralesaz/colossal-stack/graph"
	"github.com/ryanmoralesaz/colossal-stack/middleware"
	"github.com/ryanmoralesaz/colossal-stack/models"
	"github.com/ryanmoralesaz/colossal-stack/routes"
	"github.com/ryanmoralesaz/colossal-stack/storage"
	"log"
	"net/http"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db, err := storage.NewConnection(cfg.GetDSN())
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// Run migrations
	if err := models.MigrateBooks(db); err != nil {
		log.Fatal("Could not migrate database:", err)
	}

	if err := models.MigrateUsers(db); err != nil {
		log.Fatal("Could not migrate users:", err)
	}

	// Create repository
	repo := &models.Repository{DB: db}
	authRepo := &models.AuthRepository{DB: db}
	userRepo := &models.UserRepository{DB: db}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Colossal Stack Demo",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Serve front end
	app.Static("/", "./public")

	// REST API routes
	routes.SetupRoutes(app, repo, authRepo, userRepo)

	// GraphQL setup
	gqlResolver := &graph.Resolver{DB: db}
	gqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: gqlResolver}))
	gqlPlayground := playground.Handler("GraphQL Playground", "/colossal/graphql")

	// GraphQL endpoint with custom handler for context injection
	app.All("/graphql", func(c *fiber.Ctx) error {
		// Create a custom handler that injects the user context
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Inject user context from Fiber
			ctx := r.Context()
			ctx = middleware.InjectUserContext(c, ctx)

			// Create new request with updated context
			r = r.WithContext(ctx)

			// Serve GraphQL
			gqlServer.ServeHTTP(w, r)
		})

		return adaptor.HTTPHandler(handler)(c)
	})

	app.Get("/playground", adaptor.HTTPHandler(gqlPlayground)) // Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Start server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...\n", port)
	log.Printf("REST API: http://localhost:%s/api/books\n", port)
	log.Printf("GraphQL Playground: http://localhost:%s/playground\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Could not start server:", err)
	}
}
