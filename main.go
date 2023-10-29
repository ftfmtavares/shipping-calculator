package main

import (
	"flag"
	"fmt"

	"github.com/ftfmtavares/shipping-calculator/database"
	"github.com/ftfmtavares/shipping-calculator/server"
)

func main() {
	port := flag.Int("port", 80, "Server Port")
	databaseFile := flag.String("db", "products.db", "Sqlite DB file")
	flag.Parse()
	if *port < 0 {
		fmt.Println("Invalid Port")
		flag.PrintDefaults()
		return
	}

	db, err := database.NewDatabase(*databaseFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	srv := server.NewServer(*port, db)
	err = srv.ListenAndServe()
	fmt.Println(err.Error())
}
