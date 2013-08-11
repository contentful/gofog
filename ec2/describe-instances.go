package ec2

import (
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"encoding/json"
	"fmt"
	"os"
	"flag"
	"strings"
)

type filterOption []string

func (i *filterOption) String() string {
	return fmt.Sprint(*i)
}

func (i *filterOption) Set(value string) error {
	*i = append(*i, value)
	return nil
}


func DescribeInstances() {
	var filterFlags filterOption
	filter := ec2.NewFilter()
	flags := flag.NewFlagSet("describe-instances", flag.ExitOnError)
	//config := flags.String("c", "none", "the configuration file")
	regionString := flags.String("r", "us-east-1", "the AWS region")
	flags.Var(&filterFlags, "f", "the filter")
	flags.Parse(os.Args[3:])
	auth, err := aws.EnvAuth()
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
	resp, err := connection.Instances(nil, filter)
	if err != nil {
		panic(err)
	}
	reservations := resp.Reservations
	j, err := json.Marshal(reservations)
	os.Stdout.Write(j)
}
