package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dickyanth/eco-bite-v1/config"
	"github.com/dickyanth/eco-bite-v1/db"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr:   config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net:    "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	
	migrationsPath := filepath.Join(workingDir, "cmd", "migrate", "migrations")
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		log.Fatalf("Migration path does not exist: %s", migrationsPath)
	}
	
	m, err := migrate.NewWithDatabaseInstance(
		"file://" + migrationsPath,
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	
	cmd := os.Args[len(os.Args)-1]
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	default:
		log.Println("Unknown command:", cmd)
	}
}
