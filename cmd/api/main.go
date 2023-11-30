package main

import (
	"github.com/pezdel/go-auth/routes"
)

func main() {
	r := routes.NewRouter()
	r.Serve(":8080")
}
