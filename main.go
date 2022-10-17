package main

import (
	"books-api/domain/book"
	"books-api/domain/users"
	"books-api/infrastructure/authentication"
	"books-api/infrastructure/config"
	"books-api/infrastructure/database"
	"books-api/infrastructure/redis"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	docs "books-api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	config.LoadConfig()

	jwtAuth, err := authentication.ConstructJwtAuth()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Initialize()

	cache, err := redis.Initialize()

	if err != nil {
		log.Fatal(err)
	}

	err = database.MigrateIfNeed(db)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	if config.GlobalConfig.Apm.Active {
		log.Println(config.GlobalConfig.Apm.Active)
		router.Use(apmgin.Middleware(router))
	}

	// swagger
	docs.SwaggerInfo.BasePath = "/api/basic"
	router.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userUseCase := users.ConstructUserUseCase(db, jwtAuth, cache)
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
