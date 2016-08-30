package ec2

import (
	"encoding/json"
	"flag"
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"os"
	"strings"
	"time"
)

func DescribeSnapshots() {
	var filterFlags filterOption
	filter := ec2.NewFilter()
	flags := flag.NewFlagSet("describe-snapshots", flag.ExitOnError)
	//config := flags.String("c", "none", "the configuration file")
	regionString := flags.String("r", "us-east-1", "the AWS region")
	flags.Var(&filterFlags, "f", "the filter")
	flags.Parse(os.Args[3:])
	auth, authErr := aws.GetAuth("", "", "", time.Now().Add(time.Second*3600))
	if authErr != nil {
		panic(authErr)
	}
	region := aws.Regions[*regionString]
	connection := ec2.New(auth, region)
	for idx := range filterFlags {
		tokens := strings.Split(filterFlags[idx], "=")
		filter.Add(tokens[0], tokens[1])
	}
	resp, err := connection.Snapshots(nil, filter)
	if err != nil {
		panic(err)
	}
	snapshots := resp.Snapshots
	j, err := json.Marshal(snapshots)
	os.Stdout.Write(j)
}
