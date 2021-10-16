package main

import "golang-sonar-example/cmd/project/internal/application"

func main() {
	app := application.New()
	app.Run()
}
