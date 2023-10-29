package main

import (
	"flag"
	"fmt"

	"github.com/ftfmtavares/shipping-calculator/database"
	"github.com/ftfmtavares/shipping-calculator/server"
)

func main() {
	port := flag.Int("port", 80, "Server Port")
	flag.Parse()
	if *port < 0 {
		flag.PrintDefaults()
		return
	}

	productsDB := database.NewProducts()
	srv := server.NewServer(*port, productsDB)
	err := srv.ListenAndServe()
	fmt.Println(err.Error())
}
