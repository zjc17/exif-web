package main

import (
	"fmt"
	"github.com/zjc17/exif-web/internal/server"
	"github.com/zjc17/exif-web/internal/version"
	"log"
)

func main() {
	fmt.Printf("exif web %s\n\n", version.Version)

	launchParam := server.DefaultLaunchParam()
	if err := server.Launch(launchParam); err != nil {
		log.Fatal(err)
	}
}
