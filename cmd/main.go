package main

import (
	"database/sql"
	"log"

	"github.com/dickyanth/eco-bite-v1/cmd/api"
	"github.com/dickyanth/eco-bite-v1/config"
	"github.com/dickyanth/eco-bite-v1/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr:config.Envs.DBAddress,
		DBName:config.Envs.DBName,
		Net:"tcp",
		AllowNativePasswords : true,
		ParseTime: true,
	})

	if err != nil {
		log.Fatal()
	}

	initStorage(db)

	server := api.NewAPIServer(":8000",db)	
	if err:= server.Run(); err != nil{
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}