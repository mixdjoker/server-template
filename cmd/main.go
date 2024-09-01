package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mixdjoker/server-template/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := a.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
