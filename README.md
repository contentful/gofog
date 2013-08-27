# goFog

## Palindromic zero dependencies CLI for aws web services

Plain and simple copy a file somewhere in your `$PATH` and start
interacting with AWS cloud, no dependencies.

Output in JSON to make it easy to automate and use the CLI from other
scripts and tools.

## Setup

1. copy the executable
2. set `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment
   variables
3. use it

## Commands Supported

* describe instances

```sh
gofog ec2 describe-instances -r us-east-1
# Is possible to specify filters by using the -f flag
gofog ec2 describe-instances -r us-east-1 -f tag:environment=production
# The -f flag can be used multiple times to express multiple constraints
```

### SNS

* publish message

```sh
echo "the body"|./gofog sns publish -t arn:aws:sns:us-east-1:my_topic -s "this is the subject"
```
