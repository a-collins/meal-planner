package storage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	S3      *s3.S3
	Session *session.Session
}

func Hello(name string) {
	fmt.Printf("Hello, %s's world", name)
}

func (c *S3Client) Initialise() error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	})

	if err != nil {
		return err
	}
	c.Session = sess
	c.S3 = s3.New(sess)
	fmt.Println("Created session successfully")

	return nil
}
