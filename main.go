package main

import (
	"books-api/domain/book"
	"books-api/domain/users"
	"books-api/infrastructure/authentication"
	"books-api/infrastructure/config"
	"books-api/infrastructure/database"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	jwtAuth, err := authentication.ConstructJwtAuth()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Initialize()

	if err != nil {
		log.Fatal(err)
	}

	err = database.MigrateIfNeed(db)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	userUseCase := users.ConstructUserUseCase(db, jwtAuth)
	users.ConstructUserHandler(router, userUseCase, jwtAuth)

	bookUseCase := book.ConstructBookUseCase(db, jwtAuth)
	book.ConstructBookHanlder(router, bookUseCase, jwtAuth)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.GlobalConfig.Server.Host, config.GlobalConfig.Server.Port),
		Handler: router,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown", err)
	}

	log.Println("Server exiting")
}
