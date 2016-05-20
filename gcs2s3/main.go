// gcs2s3 copy file from gcs to s3. only support single object copy.
//
// same as below:
//
// gsutil cp gs://your/file/path /tmp/hoge
// aws s3 cp /tmp/hoge s3://your/file/path
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3bucket      = flag.String("s3bucket", "", "destination s3 bucket")
	s3obj         = flag.String("s3obj", "", "destination s3 object prefix")
	awsregion     = flag.String("awsregion", "ap-northeast-1", "AWS Region")
	gcscredential = flag.String("gcscredential", "credential.json", "path to service account credential JSON file")
	gcsbucket     = flag.String("gcsbucket", "hoge", "source GCS bucket")
	gcsobj        = flag.String("gcsobj", "kuke", "source GCS object prefix")
)

func main() {
	flag.Parse()
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatalf("failed to create temp file: %s", err)
	}
	b, err := ioutil.ReadFile(*gcscredential)
	if err != nil {
		log.Fatalf("failed to read %s, %s", *gcscredential, err)
	}

	scope := fmt.Sprintf("%s %s",
		storage.DevstorageReadOnlyScope,
		"https://www.googleapis.com/auth/userinfo.profile")
	conf, err := google.JWTConfigFromJSON(b, scope)
	if err != nil {
		log.Fatalf("failed to auth google API by %s, %s", *gcscredential, err)
	}
	gcpClient := conf.Client(context.Background())
	gcs, err := storage.New(gcpClient)
	if err != nil {
		log.Fatalf("failed to initialize gcs client %s", err)
	}
	// sample: https://cloud.google.com/storage/docs/json_api/v1/json-api-go-samples#setup-code
	res, err := gcs.Objects.Get(*gcsbucket, *gcsobj).Do()
	if err != nil {
		log.Fatalf("failed to get object from gcs: %s", err)
	}
	resp, err := http.Get(res.MediaLink)
	if err != nil {
		log.Fatalf("failed to download object from gcs: %s", err)
	}
	defer resp.Body.Close()

	n, err := io.Copy(tmpfile, resp.Body)
	if err != nil {
		log.Fatalf("copy object to tempfile failed %s", err)
	}
	log.Printf("object %s copy to %s: %d bytes written.", res.Name, tmpfile.Name(), n)

	// and, resp.Body to S3
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(*awsregion)})
	po, err := svc.PutObject(&s3.PutObjectInput{
		Body:   tmpfile,
		Bucket: aws.String(*s3bucket),
		Key:    aws.String(*s3obj),
	})
	if err != nil {
		log.Fatalf("failed to put object to s3 %s", err)
	}
	log.Printf("successfully put object to s3: %v", po)
}
