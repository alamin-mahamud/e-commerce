package main

import (
	"e-commerce-go-api/cmd"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	cmd.Execute()
}
