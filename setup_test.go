package main

import (
	"log"
	"os"
	"testing"

	"github.com/betopompolo/project_playlist_server/config"
)

func TestMain(m *testing.M) {
	config.AppInstance = config.Setup()
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}
