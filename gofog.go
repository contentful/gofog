package main

import (
	"os"
	"fmt"
	"github.com/hungryblank/gofog/ec2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s you need to specify a service, available services are: %s\n", os.Args[0], "ec2")
		os.Exit(1)
	}
	cmd := os.Args[1]
	if cmd == "ec2" {
		ec2.Run()
	} else {
		panic(fmt.Sprintf("unknown command %v", cmd))
	}
}
