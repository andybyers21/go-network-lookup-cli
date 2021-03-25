// cmd/my-cli/cli.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	fmt.Println("Go Network Lookup CLI v0.01")

	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
