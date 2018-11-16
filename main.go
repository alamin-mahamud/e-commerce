package main

import (
	"github.com/alamin-mahamud/e-commerce-go-api/cmd"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	cmd.Execute()
}
