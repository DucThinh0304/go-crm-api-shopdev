package main

import (
	"go-crm-api-shopdev/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8080")
}
