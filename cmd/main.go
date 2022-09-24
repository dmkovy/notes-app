package main

import (
	"github.com/dmkovy/notes-app"
	"log"
)

func main() {
	srv := new(notes.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}
}
