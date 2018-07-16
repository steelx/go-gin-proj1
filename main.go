package main

import (
	"gin_proj1/src/app"
)

func main() {
	r := app.RegisterRoutes()

	r.Run(":3000")
}
