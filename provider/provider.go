package provider

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"log"
)

type AWS struct {
	EC2 *ec2.EC2
	RDS *rds.RDS
}

func (a *AWS) FetchSubnet() (*rds.DescribeDBSubnetGroupsOutput, error) {
	svc := a.RDS
	sf := &rds.DescribeDBSubnetGroupsInput{}
	req := svc.DescribeDBSubnetGroupsRequest(sf)
	res, err := req.Send()

	if err != nil {
		log.Println(err)
	}
	return res, err
}

func (a *AWS) FetchEc2() (*ec2.DescribeInstancesOutput, error) {
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
	req := a.EC2.DescribeInstancesRequest(params)
	res, err := req.Send()

	return res, err
}
