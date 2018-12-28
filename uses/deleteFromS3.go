package uses

import (
	"log"
	urlPackage "net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sea350/ustart_go/globals"
)

//DeleteFromS3  given the full url, it parses for a key and removes an object from s3
func DeleteFromS3(url string) error {

	splt := strings.Split(url, "/")
	key := splt[len(splt)-1]

	key, _ = urlPackage.QueryUnescape(key)
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.Println("Debug text: attempting to delete " + key)

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(globals.S3Region), Credentials: credentials.NewStaticCredentials(globals.S3CredID, globals.S3CredSecret, globals.S3CredToken)}))

	// Create an uploader with the session and default options
	svc := s3.New(sess)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(globals.S3BucketName),
		Key:    aws.String(key),
	}

	_, err := svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err.Error())
		}
		return err
	}

	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.Println("Debug text: " + result.Location)

	return nil

}
