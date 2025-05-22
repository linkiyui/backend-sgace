package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadEnvVariables() (string, string, string, string, string) {
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "sasa"
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "sgace_db"
	}

	return user, password, host, port, dbname
}

var Database = func() *gorm.DB {
	user, password, host, port, dbname := loadEnvVariables()

	// Valida que las variables no estén vacías
	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		panic("Faltan variables de entorno necesarias para la conexión a la base de datos")
	}

	// Construye el DSN dinámicamente con las variables de entorno
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/New_York",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		panic(err) // O maneja el error de manera más elegante según tu caso
	}
	return db
}()
