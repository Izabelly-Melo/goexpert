package main

import (
	"context"
	"net/http"

	"github.com/Izabelly-Melo/goexpert/api/configs"
	_ "github.com/Izabelly-Melo/goexpert/api/docs"
	"github.com/Izabelly-Melo/goexpert/api/internal/entity"
	"github.com/Izabelly-Melo/goexpert/api/internal/infra/database"
	"github.com/Izabelly-Melo/goexpert/api/internal/infra/webserver/handlers"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

// @title Goexpert API
// @version 1.0
// @description Product API with user authentication
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open(configs.GetDBDrive()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // Middleware to recover from panics and return a 500 error
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "jwt", configs.TokenAuth)
			ctx = context.WithValue(ctx, "jwtExpiresIn", configs.JwtExpiresIn)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/generate_token", userHandler.GetToken)

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"),
	))

	http.ListenAndServe(":8080", r)
}
