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

func DescribeInstances() {
	var filterFlags filterOption
	filter := ec2.NewFilter()
	flags := flag.NewFlagSet("describe-instances", flag.ExitOnError)
	//config := flags.String("c", "none", "the configuration file")
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
		//fmt.Printf("foo bar %s", *config)
	}
	resp, err := connection.DescribeInstances(nil, filter)
	if err != nil {
		panic(err)
	}
	reservations := resp.Reservations
	j, err := json.Marshal(reservations)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(j)
}
