package seeders

import (
	"belajar_go_echo/config"
	"belajar_go_echo/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedUser() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	user := models.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: string(hashedPassword),
	}

	config.DB.Create(&user)
}
