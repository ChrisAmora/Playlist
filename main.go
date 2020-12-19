package main

import "github.com/betopompolo/project_playlist_server/config"

func main() {
	a := config.Setup()
	a.RunGraphql()
}
