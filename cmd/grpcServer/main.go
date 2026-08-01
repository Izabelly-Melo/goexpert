package main

import (
	"database/sql"
	"net"

	"github.com/Izabelly-Melo/goexpert/grpc/internal/database"
	"github.com/Izabelly-Melo/goexpert/grpc/internal/pb"
	"github.com/Izabelly-Melo/goexpert/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
