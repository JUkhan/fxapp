package main

import (
	"fxapp/bundle"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		bundle.Module,
	).Run()

}
