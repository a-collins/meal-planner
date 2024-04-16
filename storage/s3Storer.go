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
	fmt.Printf("Hello, %s's world\n", name)
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

// s3Client := storage.S3Client{}
// err := s3Client.Initialise()
// if err != nil {
// 	fmt.Printf("Error creating session: %v\n", err)
// 	os.Exit(1)
// }

// _, err = s3Client.Session.Config.Credentials.Get()
// fmt.Println(err)

// storage.Hello("aaron")
// result, err := s3Client.S3.ListBuckets(nil)
// if err != nil {
// 	fmt.Printf("Error listing buckets: %v\n", err)
// 	os.Exit(1)
// }

// fmt.Println("Buckets:")

// for _, b := range result.Buckets {
// 	fmt.Printf("* %s created on %s\n",
// 		aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
// }
