package main

import (
	api "github.com/xandervanderweken/GoHomeNet/api/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/container"
)

func main() {
	c := container.New()
	r := api.NewRouter(c)
	api.StartServer(r)
}
