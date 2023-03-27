// Package config provides the database connection
package config

import (
	"fmt"
	"log"
	"os"
	"time"

	// driver mysql on this implementation
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

type infoDatabasePostgreSQL struct {
	Read struct {
		Hostname   string
		Name       string
		Username   string
		Password   string
		Port       string
		Parameter  string
		Timezone   string
		DriverConn string
	}
	Write struct {
		Hostname   string
		Name       string
		Username   string
		Password   string
		Port       string
		Parameter  string
		Timezone   string
		DriverConn string
	}
}

func (infoDB *infoDatabasePostgreSQL) getPostgreConn(nameMap string) (err error) {

	viper.SetConfigFile("config.json")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = mapstructure.Decode(viper.GetStringMap(nameMap), infoDB)
	if err != nil {
		return
	}

	infoDB.Read.DriverConn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		infoDB.Read.Hostname, infoDB.Read.Username, infoDB.Read.Password, infoDB.Read.Name, infoDB.Read.Port, infoDB.Read.Timezone)
	infoDB.Write.DriverConn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		infoDB.Write.Hostname, infoDB.Write.Username, infoDB.Write.Password, infoDB.Write.Name, infoDB.Write.Port, infoDB.Write.Timezone)
	return
}

func initPostgreDB(inGormDB *gorm.DB, infoPg infoDatabasePostgreSQL) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	inGormDB, err := gorm.Open(postgres.Open(infoPg.Write.DriverConn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	err = inGormDB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{postgres.Open((infoPg.Read.DriverConn))},
	}))
	if err != nil {
		return nil, err
	}

	return inGormDB, nil
}
