package sns

import (
	"fmt"
	"os"
)

func Run() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "%s %s you need to specify a command, available commands are: %s\n", os.Args[0], "sns", "publish")
		os.Exit(1)
	}
	cmd := os.Args[2]
	if cmd == "publish" {
		Publish()
	} else {
		fmt.Sprintf("unknown sns command %v", cmd)
		os.Exit(1)
	}
}
