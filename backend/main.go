package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend/db"
	"backend/middlewares"
	"backend/router"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dotenvInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("DotEnv initialized")
	}
}
func InitSupabase() error {
	var err error
	supabaseClient, err := middlewares.SupabaseInit(os.Getenv("SUPABASE_KEY"), os.Getenv("SUPABASE_URL"))
	if err != nil {
		return fmt.Errorf("error initializing Supabase client: %v", err)
	} else {
		fmt.Println("Supabase Initialized")
	}
	middlewares.SupabaseClient = supabaseClient
	return nil
}
func InitDatabase() error {
	var err error
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	pgdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.GormDB = pgdb
	// Run Migration
	db.MigrateSchema(db.GormDB)
	if err != nil {
		return fmt.Errorf("error initalizing PostgreSQL: %v", err)
	} else {
		fmt.Println("PostgreSQL Database Initialized")
	}

	return nil
}

func main() {
	dotenvInit()
	InitSupabase()
	InitDatabase()
	r := gin.Default()
	router.InitRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
