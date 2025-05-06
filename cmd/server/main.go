package main

import (
	"go-ecommerce-api/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run()
}
