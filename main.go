package main

import (
	"os"

	"github.com/dhiiyaur/go-mangamee/routers"
)

func main() {

	e := routers.Index()

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	// e.Logger.Fatal(e.Start(":8000"))

}
