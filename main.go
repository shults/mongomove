package main

import (
	"github.com/shults/mongomove/application"
	"os"
)

func main() {
	app := application.NewCliRunner(os.Args)
	app.Run()
}
