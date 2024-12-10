package main

import (
	"fmt"
	"os"

	_ "github.com/hypermodeinc/modus/sdk/go"
	"github.com/hypermodeinc/modus/sdk/go/pkg/console"
)

func SayHello(name *string) string {

	var s string
	if name == nil {
		s = "World"
	} else {
		s = *name
	}

	return fmt.Sprintf("Hello, %s!", s)
}

func GetAuthClaims() (string, error) {
	console.Log(os.Getenv("CLAIMS"))
	return "foo", nil
}
