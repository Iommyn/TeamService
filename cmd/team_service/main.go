package main

import (
	"TeamService/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	bootstrap := app.NewApp()
	bootstrap.Run()
}
