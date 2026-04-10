package config

import (
	"log"

	"github.com/optimus/backend/models"
	"github.com/optimus/backend/utils"
)

func SeedData() {
	roles := []string{
		"admin",
		"ppk",
		"pokja/pp",
		"kepala bagian",
		"kepala biro",
	}

	for _, roleName := range roles {
		var role models.Role
		err := DB.Where("name = ?", roleName).First(&role).Error
		if err != nil {
			// If not found, create it
			role = models.Role{
				Name: roleName,
			}
			if err := DB.Create(&role).Error; err != nil {
				log.Printf("Failed to seed role %s: %v", roleName, err)
			} else {
				log.Printf("Role seeded: %s", roleName)
			}
		}
	}

	// Ensure Admin User exists and is linked to the 'admin' role
	var adminRole models.Role
	if err := DB.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		log.Printf("Admin role not found, skipped user seeding: %v", err)
		return
	}

	var adminUser models.User
	email := "adminmonev@gmail.com"
	err := DB.Where("email = ?", email).First(&adminUser).Error
	if err != nil {
		// If not found, create it
		isActive := true
		fullName := "Administrator"

		adminUser = models.User{
			FullName: &fullName,
			Email:    email,
			Password: utils.HashSHA512("AdM0nev!n#88"),
			IsActive: &isActive,
			RoleId:   adminRole.Id,
		}

		if err := DB.Create(&adminUser).Error; err != nil {
			log.Printf("Failed to seed admin user: %v", err)
			return
		}
		log.Println("Admin user seeded: " + email)
	} else {
		// Check if admin user has the correct role (optional, but good for consistency)
		if adminUser.RoleId != adminRole.Id {
			adminUser.RoleId = adminRole.Id
			DB.Save(&adminUser)
			log.Println("Admin user role updated to 'admin'")
		}
		log.Println("Admin user already exists")
	}
}
