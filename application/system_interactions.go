package application

import (
	"fmt"
	"os"
)

type systemInteractions interface {
	Printf(str string, args ...interface{})
	Exit(exitCode int)
}

type realSystemInteractions struct{}

func (r realSystemInteractions) Exit(exitCode int) {
	os.Exit(exitCode)
}

func (r realSystemInteractions) Printf(str string, args ...interface{}) {
	_, err := fmt.Printf(str, args...)

	if err != nil {
		panic(err)
	}
}


