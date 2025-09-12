package main

import (
	"github.com/xandervanderweken/GoHomeNet/internal/api"
	"github.com/xandervanderweken/GoHomeNet/internal/container"
)

func main() {
	c := container.New()
	r := api.NewRouter(c)
	api.StartServer(r)
}
