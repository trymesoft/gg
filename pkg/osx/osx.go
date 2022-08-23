package osx

import (
	"fmt"
	"github.com/bingoohuang/gg/pkg/osx/env"
	"os"
)

func ExitIfErr(err error) {
	if err != nil {
		Exit(err.Error(), 1)
	}
}

func Exit(msg string, code int) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(code)
}

func EnvSize(envName string, defaultValue int) int {
	return env.Size(envName, defaultValue)
}
