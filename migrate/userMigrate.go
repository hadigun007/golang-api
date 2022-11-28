package main

// func MigrateUser() {

// 	type UserModel struct {
// 		gorm.Model
// 		Title string
// 		Price string
// 	}

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Failed Load .env")
// 	}
// 	dsn := os.Getenv("DB_URL")
// 	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	DB.AutoMigrate(UserModel{})
// }
