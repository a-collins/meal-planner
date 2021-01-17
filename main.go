package main

import (
	"fmt"
	"os"

	"github.com/a-collins/meal-planner/storage"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	s3Client := storage.S3Client{}
	err := s3Client.Initialise()
	if err != nil {
		fmt.Println("Error creating session: %v", err)
		os.Exit(1)
	}

	_, err = s3Client.Session.Config.Credentials.Get()
	fmt.Println(err)

	storage.Hello("aaron")
	result, err := s3Client.S3.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error listing buckets: %v", err)
		os.Exit(1)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

}
