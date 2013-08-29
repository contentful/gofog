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
	switch cmd {
	case "describe-instances":
		DescribeInstances()
		os.Exit(0)
	case "create-snapshot":
		CreateSnapshot()
		os.Exit(0)
	case "describe-snapshots":
		DescribeSnapshots()
		os.Exit(0)
	case "delete-snapshot":
		DeleteSnapshot()
		os.Exit(0)
	}
	fmt.Sprintf("unknown command %v", cmd)
	os.Exit(1)
}

type filterOption []string

func (i *filterOption) String() string {
	return fmt.Sprint(*i)
}

func (i *filterOption) Set(value string) error {
	*i = append(*i, value)
	return nil
}
