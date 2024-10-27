package main

import (
	"booking-service/api/handler"
	"booking-service/repository"
	"booking-service/service"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type DataSources struct {
	DB          *gorm.DB
	redisClient *redis.Client
}

func main() {
	fmt.Println("starting booking service")

	dataSource, err := initDS()

	if err != nil {
		fmt.Println(err)
		return
	}

	router, err := setupRouter(dataSource)
	if err != nil {
		log.Fatalf("Booking Service: failed to setup router:  %v", err)
	}

	port := os.Getenv("BOOKING_SERVICE_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Running server on port " + port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start the server: %s\n", err)
		}
	}()

	fmt.Printf("Server started %v\n", srv.Addr)

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}

}

func initDS() (*DataSources, error) {

	fmt.Println("initializing data sources")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Karachi",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	fmt.Println("Connecting to database", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error opening DB: %w", err)
	}
	fmt.Println("Successfully initialized data sources")

	fmt.Println("Initializing Redis Data Source")
	options := redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
		Protocol: 2,
	}
	redisClient := redis.NewClient(&options)
	fmt.Println("Initializing Redis Data Source", options)

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %w", err)
	}
	fmt.Println("Successfully initialized Redis Data Source")
	return &DataSources{DB: db, redisClient: redisClient}, nil
}

func setupRouter(d *DataSources) (*gin.Engine, error) {

	engine := gin.New()

	seatsRepository := repository.PgSeatsRepository{DB: d.DB}
	seatService := service.NewSeatsService(seatsRepository)
	bookingService := service.NewBookingService(seatService, d.redisClient)

	handler.NewBookingHandler(engine, bookingService)

	return engine, nil
}
