package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/vavarodrigues/test-aws-sdk-go-v2/rds"
	"log"
)

func main() {
	ec2client, err := ec2client()
	if err != nil {
		log.Fatal("unable to create a client for EC2 ", err)
	}

	params := &ec2.DescribeInstancesInput{
		Filters: []ec2.Filter{
			{
				Name: aws.String("private-dns-name"),
				Values: []string{
					"ip-172-20-15-213.us-east-2.compute.internal",
				},
			},
		},
	}
	log.Println("trying to describe instance")
	req := ec2client.DescribeInstancesRequest(params)
	res, err := req.Send()

	r := rds.RDS{EC2: ec2client}

	fmt.Println(r)
	fmt.Println(res)

}

func ec2client() (*ec2.EC2, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	return ec2.New(cfg), nil
}
