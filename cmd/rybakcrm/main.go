package main

import (
	"context"
	"fmt"

	"crm-backend/internal/rybakcrm/app"
	"crm-backend/internal/rybakcrm/config"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", cfg.Env)

	mainApp := app.New(ctx, cfg)

	mainApp.Start(ctx)
}
