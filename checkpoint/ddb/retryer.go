package ddb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Retryer interface contains one method that decides whether to retry based on error
type Retryer interface {
	ShouldRetry(error) bool
}

type DefaultRetryer struct {
	Retryer
}

func (r *DefaultRetryer) ShouldRetry(err error) bool {
	if r != nil {
		awsErr, ok := err.(awserr.Error)
		if ok {
			if awsErr.Code() == dynamodb.ErrCodeProvisionedThroughputExceededException {

				return true
			}
			fmt.Printf("awsErr: %#v", awsErr.Code())
		}
		fmt.Printf("awsErr: %#v", awsErr)

	}
	return false
}
