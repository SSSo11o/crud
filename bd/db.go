package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=1112 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	DB = db
	log.Println("База данных успешно подключена")

	/* когда не локально и потом используем docker-compose.yaml
	host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, user, password, dbname, port)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Ошибка подключения к базе данных:", err)
		}

		DB = db
		log.Println("База данных успешно подключена")
	*/
}
