package main

import (
	"flag"
	"fmt"
	"log"

	"gin-starter/internal/config"
	"gin-starter/internal/model"
	"gin-starter/internal/repository"
	"gin-starter/pkg/utils"
)

const USAGE_DOCS = `
# Run migrations
cli --migrate

# Seed an admin user
cli --seed --email=<amdin email> --password=<admin password>`

func main() {
	migrate := flag.Bool("migrate", false, "Run database migration")
	seed := flag.Bool("seed", false, "Seed admin user")
	email := flag.String("email", "", "Email for seeding admin user")
	password := flag.String("password", "", "Password for seeding admin user")

	flag.Parse()

	if *migrate == false && *seed == false {
		log.Fatalln(USAGE_DOCS)
	}

	if err := config.LoadConfig(".env"); err != nil {
		log.Fatalln(`Please make sure .env file exists or env variable TWELVE_FACTOR_MODE is set to "true": `, err)
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatalln("Error: database connection failed: ", err)
	}

	if *migrate {
		runMigration()
	}

	if *seed {
		if *email == "" || *password == "" {
			fmt.Println("Error: Both email and password must be provided for seeding.")
			return
		}
		seedAdmin(*email, *password)
	}
}

func runMigration() {
	if err := config.Config.DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalln("Database migration failed: ", err)
	}
	fmt.Println("Database migration was successful.")
}

func seedAdmin(email, password string) {
	hashedPassword, _ := utils.HashPassword(password)

	userReposiroty := repository.NewUserRepository(config.Config.DB)

	user, err := userReposiroty.Create(model.User{
		Email:    email,
		Password: hashedPassword,
	})
	if err != nil {
		log.Fatalln("Seeding admin user failed: ", err)
	}

	log.Printf("Seeding was successful: %+v", user)
}
