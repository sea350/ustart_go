package uses

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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
	for index, elem := range arr {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Index and len() thereof: " + strconv.Itoa(index) + "|| " + strconv.Itoa(len(elem)))
	}
	dec, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(dec)
	img, err := png.Decode(r)
	if err != nil {
		panic("Bad png")
	}

	//convert decoder to file
	f, err := os.OpenFile(filename+".png", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Cannot open file")
	}

	png.Encode(f, img)

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1"), Credentials: credentials.NewStaticCredentials("AKIAJILB2MI6CPZKYOFA", "dgyZx0eLnJhXue/UBS9BWXvPycOAYjX60M3NJzTP", "")}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("ustart-bucket"),
		Key:    aws.String(filename + ".png"),
		Body:   f,
	})
	if err != nil {
		return url, fmt.Errorf("failed to upload file, %v", err)
	}

	url = result.UploadID
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Debug text: " + result.Location)

	return url, nil

	// w.Header().Set("Content-Type", "image/png")
	// io.(w, dec)
}
