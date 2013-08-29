package ec2

import (
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"encoding/json"
	"os"
	"strings"
	"fmt"
	"flag"
)

func DeleteSnapshot() {
	flags := flag.NewFlagSet("delete-snapshot", flag.ExitOnError)
	regionString := flags.String("r", "us-east-1", "the AWS region")
	flags.Parse(os.Args[4:])
	snapshotIds := strings.Split(os.Args[3], ",")
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err)
	}
	region := aws.Regions[*regionString]
	fmt.Printf("%v", snapshotIds)
	connection := ec2.New(auth, region)

	resp, err := connection.DeleteSnapshots(snapshotIds)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	j, err := json.Marshal(resp)
	os.Stdout.Write(j)
}
