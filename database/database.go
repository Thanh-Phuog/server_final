package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Kết nối đến PostgreSQL
func Connect() {
	dns := "host=postgres user=myuser password=mypassword dbname=bookmana port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

	// dns := "host=127.0.0.1 user=myuser password=mypassword dbname=bookmana port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	//dns := "postgresql://book_manager_db_v9pa_user:TvOr4JEPepiAMt33xAn9jrvDwIu9ozfh@dpg-cs9m5s08fa8c73ccj3rg-a.singapore-postgres.render.com/book_manager_db_v9pa"
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//// Migration
	// DB.AutoMigrate(&models.Book{})
	log.Println("Database connected successfully!")
}
