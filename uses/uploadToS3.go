package uses

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//UploadToS3 ...
func UploadToS3(based64 string, filename string) (string, error) {
	// data := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEX/TQBcNTh/AAAACklEQVR4nGNiAAAABgADNjd8qAAAAABJRU5ErkJggg=="
	// The actual image starts after the ","
	var url string
	var arr []string
	i := strings.Index(based64, ",")
	if i < 0 {
		log.Fatal("no comma")
	} else {
		arr = strings.Split(based64, `,`)
	}
	// pass reader to NewDecoder
	//dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[i+1:]))
	dec, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		panic(err)
	}

	//convert decoder to file
	f, err := os.Create(filename + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-2")}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("arn:aws:s3:::ustart-bucket"),
		Key:    aws.String(`AKIAIUTR3SDBQJPINNGA`),
		Body:   f,
	})
	if err != nil {
		return url, fmt.Errorf("failed to upload file, %v", err)
	}

	url = result.UploadID
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Debug text: " + result.Location + "|| " + result.UploadID)

	return url, nil

	// w.Header().Set("Content-Type", "image/png")
	// io.(w, dec)
}
