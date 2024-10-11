package main

import (
	"context"
	"mole/internal/app"
)

func main() {
	instance := app.New(context.Background())
	instance.Run()
}
