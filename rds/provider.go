package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type RDS struct {
	EC2 *ec2.EC2
}

func (r *RDS) rdsclient() *rds.RDS {
	return rds.New(r.EC2.Config)
}
