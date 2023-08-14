package repository

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

func TestRepository_Registry(t *testing.T) {
	dsn := "admin:admin@tcp(localhost:3306)/conversions_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("Failed to connect to the test database:", err)
	}
	defer db.Exec("DELETE FROM conversion_models")
	repository := NewRepository(db)

	success := repository.Registry(nil, "EUR", "USD", 10000, 1200, 12000000)

	if !success {
		t.Fatal("Success false failed to registry on database")
	}

	log.Println("[Test Repository] Success:", success)

}
