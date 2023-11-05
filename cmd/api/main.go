package main

import (
	"github.com/pezdel/go-template/routes"
)

func main() {
	r := routes.NewRouter()
	r.Serve(":8080")
}
