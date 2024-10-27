package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DataSources struct {
	DB          *gorm.DB
	redisClient *redis.Client
}

func initDS() (*DataSources, error) {

	fmt.Println("Initializing data sources")

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
