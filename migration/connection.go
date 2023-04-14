package migration

import (
	"boilerplate/config"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	// "host=localhost user=postgres password=riyan dbname=boilerplate2 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	dsn := fmt.Sprintf(
		`host='%v' user='%v' password='%v' dbname='%v' port='%v' sslmode=disable TimeZone=%s`,
		config.GetEnv("DB_HOST"), config.GetEnv("DB_USERNAME"), config.GetEnv("DB_PASSWORD"), config.GetEnv("DB_NAME"), config.GetEnv("DB_PORT"), config.GetEnv("TIME_ZONE"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal("migration: can't connect to database")
	}

	sqlDB, _ := database.DB()
	maxIdleConns, _ := strconv.Atoi(config.GetEnv("DB_MAX_IDDLE_CONN"))
	maxOpenConns, _ := strconv.Atoi(config.GetEnv("DB_MAX_OPEN_CONN"))
	connMaxLifetime, _ := strconv.Atoi(config.GetEnv("DB_CONN_MAX_LIFETIME_HOUR"))
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(connMaxLifetime))
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("migration: can't ping to database")
	}

	fmt.Println("migration: connection opened to database")
	fmt.Println("migration: start")
	err = database.AutoMigrate(
		&UserTypes{}, &Users{}, &UserDatas{},
	)
	if err != nil {
		log.Fatal("migration: migration failed")
	}
	// initial user type
	if err := database.First(&UserTypes{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		userTypes := []UserTypes{
			{Code: "super_admin", Name: "Super Admin"},
			{Code: "admin", Name: "Admin"},
			{Code: "admin_keuangan", Name: "Admin Keuangan"},
			{Code: "admin_humas", Name: "Admin Hubungan Masyarakat"},
		}
		database.Create(userTypes)
	}
	//
	if err := database.First(&UserDatas{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		userData := []UserDatas{
			{ID: "e1aa18e1-3400-45cb-a883-7394d87f2abd", Name: "nama saya sebenarnya ada 2"},
		}
		database.Create(userData)
	}
	if err := database.First(&Users{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		users := []Users{
			{UserDataID: "e1aa18e1-3400-45cb-a883-7394d87f2abd", UserTypeCode: "super_admin", UserName: "admin ganteng", Email: "adminganteng@mail.com", Password: "$2a$10$nEr9y50xz8oMQQaAzAjDaOif4/75XyH8DCQ1WqtZ.bYyNUEkC28aK"},
		}
		database.Create(users)
	}
	fmt.Println("migration: done")
	sqlDB.Close()
}
