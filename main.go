package main

import (
	"fmt"
	"github.com/contentful/gofog/ec2"
	"github.com/contentful/gofog/sns"
	"os"
)

var gofogVersion = "unknown"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s you need to specify a service, available services are: %s\n", os.Args[0], "ec2, sns")
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "ec2":
		ec2.Run()
		os.Exit(0)
	case "sns":
		sns.Run()
		os.Exit(0)
    case "-v":
        fmt.Printf("Gofog version %s\n", gofogVersion)
        os.Exit(0)
	}
	fmt.Sprintf("unknown command %v", cmd)
	os.Exit(1)
}
