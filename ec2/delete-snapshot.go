package ec2

import (
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"encoding/json"
	"os"
	"time"
	"strings"
	"flag"
)

func DeleteSnapshot() {
	flags := flag.NewFlagSet("delete-snapshot", flag.ExitOnError)
	regionString := flags.String("r", "us-east-1", "the AWS region")
	flags.Parse(os.Args[4:])
	snapshotIds := strings.Split(os.Args[3], ",")
	auth, authErr := aws.GetAuth("", "", "", time.Now().Add(time.Second*3600))
	if err != nil {
		panic(err)
	}
	region := aws.Regions[*regionString]
	connection := ec2.New(auth, region)

	resp, err := connection.DeleteSnapshots(snapshotIds)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	j, err := json.Marshal(resp)
	os.Stdout.Write(j)
}
