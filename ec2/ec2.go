package ec2

import (
	"fmt"
	"os"
)
func Run() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "%s %s you need to specify a command, available commands are: %s\n", os.Args[0], "ec2", "describe-instances")
		os.Exit(1)
	}
	cmd := os.Args[2]
	if cmd == "describe-instances" {
		DescribeInstances()
	} else {
		panic(fmt.Sprintf("unknown ec2 command %v", cmd))
	}
}
