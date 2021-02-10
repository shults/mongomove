package main

import (
	"github.com/shults/mongomove/application"
	"os"
)

func main() {
	app := application.NewCliApp(os.Args)
	app.Run()
}
