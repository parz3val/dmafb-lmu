package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/parz3val/dmafb-lmu/docsdb"
	"github.com/parz3val/dmafb-lmu/internal/routes"
)

func main() {
	fmt.Println("Welcome to DMAFB: Do Me a Favor Buddy")
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
		panic(err)
	}
	uri := os.Getenv("DB_URI")
	datasource, err := docsdb.NewDataSource(uri)
	if err != nil {
		fmt.Println(err)
		panic("Something went wrong while doing the db connection")
	}

	router := routes.Router(datasource)
	router.Logger.Fatal(router.Start(":3000"))
}
