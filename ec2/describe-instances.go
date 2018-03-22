package ec2

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"os"
	"strings"
	"time"
)

const maxRetries = 3

// makeRequestWithRetries performs the DescribeInstances request with a
// Fibonacci backoff up to maxRetries attempts.
func makeRequestWithRetries(connection *ec2.EC2, filter *ec2.Filter) (*ec2.InstancesResp, error) {
	for currentAttempt := 0; currentAttempt < maxRetries; currentAttempt++ {
		resp, err := connection.DescribeInstances(nil, filter)
		if err == nil {
			return resp, err
		}

		fmt.Fprintln(os.Stderr, err)
		time.Sleep(time.Second * time.Duration(fib(currentAttempt)))
	}

	return nil, fmt.Errorf("reached maximum retry attempts on DescribeInstances")
}

func DescribeInstances() {
	var filterFlags filterOption
	filter := ec2.NewFilter()
	flags := flag.NewFlagSet("describe-instances", flag.ExitOnError)
	regionString := flags.String("r", "us-east-1", "the AWS region")

	flags.Var(&filterFlags, "f", "the filter")
	flags.Parse(os.Args[3:])

	auth, err := aws.GetAuth("", "", "", time.Now().Add(time.Second*3600))
	if err != nil {
		panic(err)
	}

	region := aws.Regions[*regionString]
	connection := ec2.New(auth, region)
	for idx := range filterFlags {
		tokens := strings.Split(filterFlags[idx], "=")
		filter.Add(tokens[0], tokens[1])
	}

	resp, err := makeRequestWithRetries(connection, filter)
	if err != nil {
		// Error has already been logged out
		os.Exit(1)
	}

	reservations := resp.Reservations
	enc := json.NewEncoder(os.Stdout)
	_ = enc.Encode(reservations)
}

// fib returns the nth Fibonacci number in the sequence
func fib(n int) int {
	a, b := 1, 1

	for i := 1; i <= n; i++ {
		a, b = b, a+b
	}

	return a
}
