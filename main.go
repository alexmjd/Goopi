package main

import (
	"os"
)

func main() {

	a := App{}
	a.Init(
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBNAME"),
	)

	a.Run(":9999")
}
