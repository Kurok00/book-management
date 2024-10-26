package database

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"book_management/models"
)

var DB *gorm.DB

func Connect() {
	// Cấu hình thông tin kết nối
	//dsn := "host=localhost user=admin password=asdqwe123 dbname=book-management port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	dsn := "host=dpg-cse57ptsvqrc73esl8ag-a.singapore-postgres.render.com user=book_management_p7nw_user password=eOBW5Flts5evwvk3VoIui6aSYJM69UNK dbname=book_management_p7nw port=5432 sslmode=require TimeZone=Asia/Shanghai"
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to the database successfully!")

	// Tự động migrate mô hình dữ liệu
	DB.AutoMigrate(&models.Book{})
}