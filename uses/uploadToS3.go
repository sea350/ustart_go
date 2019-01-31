package uses

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nfnt/resize"
	"github.com/sea350/ustart_go/globals"
)

//UploadToS3 ...
func UploadToS3(based64 string, filename string) (string, error) {
	// data := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEX/TQBcNTh/AAAACklEQVR4nGNiAAAABgADNjd8qAAAAABJRU5ErkJggg=="
	// The actual image starts after the ","
	var url string
	var arr []string
	i := strings.Index(based64, ",")
	if i < 0 {
		log.Panic("no comma")
	} else {
		arr = strings.Split(based64, `,`)
	}
	// pass reader to NewDecoder
	//dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[i+1:]))
	dec, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(dec)
	img, _, err := image.Decode(r)
	if err != nil {
		return url, err
	}

	imgPrime := resize.Resize(0, uint(img.Bounds().Dy()), img, resize.Lanczos3)

	buff := new(bytes.Buffer)

	// encode image to buffer
	err = png.Encode(buff, imgPrime)
	if err != nil {
		return url, err
	}
	// convert buffer to reader
	reader := bytes.NewReader(buff.Bytes())

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(globals.S3Region), Credentials: credentials.NewStaticCredentials(globals.S3CredID, globals.S3CredSecret, globals.S3CredToken)}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(globals.S3BucketName),
		Key:         aws.String(filename + ".png"),
		Body:        reader,
		ContentType: aws.String("image/png"),
	})
	if err != nil {
		return url, fmt.Errorf("failed to upload file, %v", err)
	}

	url = result.Location
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.Println("Debug text: " + result.Location)

	return url, nil

	// w.Header().Set("Content-Type", "image/png")
	// io.(w, dec)
}
