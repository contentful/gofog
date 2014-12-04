package sns

import (
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/exp/sns"
	"os"
	"time"
	"encoding/json"
	"flag"
	"io/ioutil"
)

func Publish() {
	flags := flag.NewFlagSet("publish", flag.ExitOnError)
	regionString := flags.String("r", "us-east-1", "the AWS region")
	subjectString := flags.String("s", "no subject", "the subject of the message")
	formatString := flags.String("f", "", "the format of the message")
	topicString := flags.String("t", "", "the AWS SNS topic")
	helpBool := flags.Bool("h", false, "print this help message")
        flags.Parse(os.Args[3:])
	if *helpBool {
		flags.PrintDefaults()
		os.Exit(0)
	}
	stdinBytes, stdinErr := ioutil.ReadAll(os.Stdin)
	if stdinErr != nil {
		panic(stdinErr)
	}
	auth, authErr := aws.GetAuth("", "", "", time.Now().Add(time.Second*3600))
	if authErr != nil {
		panic(authErr)
	}
	region := aws.Regions[*regionString]
	xns := sns.New(auth, region)
	pubOpt := &sns.PublishOpt{string(stdinBytes), *formatString, *subjectString, *topicString}
	result, snsErr := xns.Publish(pubOpt)
	if snsErr != nil {
		panic(snsErr)
	}
	j, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(j)
	os.Exit(0)
}
