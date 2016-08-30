package ec2

import (
	"encoding/json"
	"flag"
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"os"
	"time"
)

func CreateTags() {
	flags := flag.NewFlagSet("create-tags", flag.ExitOnError)
	regionString := flags.String("r", "us-east-1", "the AWS region")
	tagKey := flags.String("k", "", "tag key")
	tagValue := flags.String("v", "", "tag value")
	flags.Parse(os.Args[4:])
	instanceIdString := os.Args[3]
	auth, authErr := aws.GetAuth("", "", "", time.Now().Add(time.Second*3600))
	if authErr != nil {
		panic(authErr)
	}
	region := aws.Regions[*regionString]
	connection := ec2.New(auth, region)
	instanceIds := make([]string, 1)
	instanceIds[0] = instanceIdString
	tag := ec2.Tag{Key: *tagKey, Value: *tagValue}
	tags := make([]ec2.Tag, 1)
	tags[0] = tag
	resp, err := connection.CreateTags(instanceIds, tags)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	j, err := json.Marshal(resp)
	os.Stdout.Write(j)
}
