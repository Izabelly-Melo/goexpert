package main

import (
	"context"
	"net/http"

	"github.com/Izabelly-Melo/goexpert/api/configs"
	"github.com/Izabelly-Melo/goexpert/api/internal/entity"
	"github.com/Izabelly-Melo/goexpert/api/internal/infra/database"
	"github.com/Izabelly-Melo/goexpert/api/internal/infra/webserver/handlers"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"gorm.io/gorm"
)

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

	http.ListenAndServe(":8080", r)
}
