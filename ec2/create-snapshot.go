package ec2

import (
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"encoding/json"
	"os"
	"time"
	"flag"
)

func CreateSnapshot() {
	flags := flag.NewFlagSet("create-snapshot", flag.ExitOnError)
	regionString := flags.String("r", "us-east-1", "the AWS region")
	descriptionString := flags.String("d", "", "description")
	flags.Parse(os.Args[4:])
	volumeIdString := os.Args[3]
	auth, authErr := aws.GetAuth("", "", "", time.Now().Add(time.Second*3600))
	if err != nil {
		panic(err)
	}
	region := aws.Regions[*regionString]
	connection := ec2.New(auth, region)
	resp, err := connection.CreateSnapshot(volumeIdString, *descriptionString)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	j, err := json.Marshal(resp)
	os.Stdout.Write(j)
}
