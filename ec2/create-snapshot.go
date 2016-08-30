package ec2

import (
	"encoding/json"
	"flag"
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"os"
	"time"
)

func CreateSnapshot() {
	flags := flag.NewFlagSet("create-snapshot", flag.ExitOnError)
	regionString := flags.String("r", "us-east-1", "the AWS region")
	descriptionString := flags.String("d", "", "description")
	flags.Parse(os.Args[4:])
	volumeIdString := os.Args[3]
	auth, authErr := aws.GetAuth("", "", "", time.Now().Add(time.Second*3600))
	if authErr != nil {
		panic(authErr)
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
