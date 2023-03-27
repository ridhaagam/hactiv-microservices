package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB cradential
var (
	// host     = os.Getenv("DB_HOST")
	// port     = os.Getenv("DB_PORT")
	// username = os.Getenv("DB_USERNAME")
	// password = os.Getenv("DB_PASSWORD")
	// dbname   = os.Getenv("DB_POSGRES")

	host     = "127.0.0.1"
	port     = 5432
	username = "root"
	password = "root"
	dbname   = "hacktiv"
)

func New() *gorm.DB {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, username, password, dbname)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(0)
	err = db.AutoMigrate()

	if err != nil {
		panic(err)
	}

	return db
}
