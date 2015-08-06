package main

import (
	"github.com/Stephen304/cmdfolder"

	"./twitter"
)

func main() {
	folder := cmdfolder.New()

	// Add services
	folder.AddFolder("twitter", twitter.New())

	// Run it
	folder.Run()
}
