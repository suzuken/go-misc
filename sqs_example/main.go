package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type svc struct {
	*sqs.SQS
}

func New(region string) *svc {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{Region: aws.String(region)},
	}))
	return &svc{sqs.New(sess)}
}

func (s *svc) listQueues() ([]*string, error) {
	result, err := s.ListQueues(nil)
	if err != nil {
		return nil, err
	}
	return result.QueueUrls, nil
}

func (s *svc) queueExists(name string) bool {
	name, err := s.queueURL(name)
	return err == nil && name != ""
}

func (s *svc) queueURL(name string) (string, error) {
	result, err := s.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})
	if err != nil {
		return "", err
	}
	return *result.QueueUrl, nil
}

func (s *svc) createQueue(queueName string) (string, error) {
	result, err := s.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
		Attributes: aws.StringMap(map[string]string{
			"DelaySeconds":                  "60",
			"MessageRetentionPeriod":        "86400",
			"ReceiveMessageWaitTimeSeconds": "10",
		}),
	})
	if err != nil {
		return "", err
	}
	log.Printf("%s created.", *result.QueueUrl)
	return *result.QueueUrl, nil
}

func (s *svc) deleteQueue(queueURL string) error {
	_, err := s.DeleteQueue(&sqs.DeleteQueueInput{
		QueueUrl: aws.String(queueURL),
	})
	if err != nil {
		return err
	}
	log.Printf("%s deleted.", queueURL)
	return nil
}

type T struct {
	SomeString string
	Numbers    []int
	IDs        []string
	Points     []float64
}

func (s *svc) sendJSON(queueURL string, thing T) (*sqs.SendMessageOutput, error) {
	b, err := json.Marshal(thing)
	if err != nil {
		return nil, err
	}
	return s.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"something": &sqs.MessageAttributeValue{
				DataType:    aws.String("Binary.json"),
				BinaryValue: b,
			},
		},
		MessageBody: aws.String("test binary"),
		QueueUrl:    &queueURL,
	})
}

func (s *svc) sendBytes(queueURL string) (*sqs.SendMessageOutput, error) {
	return s.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"something": &sqs.MessageAttributeValue{
				DataType:    aws.String("Binary"),
				BinaryValue: []byte("some"),
			},
		},
		MessageBody: aws.String("test binary"),
		QueueUrl:    &queueURL,
	})
}

func (s *svc) send(queueURL string) (*sqs.SendMessageOutput, error) {
	return s.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
		MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
		QueueUrl:    &queueURL,
	})
}

func (s *svc) receiveWithLongPoll(queueURL string) (*sqs.ReceiveMessageOutput, error) {
	return s.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &queueURL,
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(0),
		WaitTimeSeconds:     aws.Int64(20),
	})
}

func (s *svc) receive(queueURL string) (*sqs.ReceiveMessageOutput, error) {
	return s.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &queueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(0),
		WaitTimeSeconds:     aws.Int64(0),
	})
}

func (s *svc) receiveJSON(queueURL string) ([]T, error) {
	result, err := s.receiveWithLongPoll(queueURL)
	if err != nil {
		return nil, err
	}
	ts := make([]T, 0, 10)
	for _, m := range result.Messages {
		for k, v := range m.MessageAttributes {
			if strings.HasSuffix(*v.DataType, ".json") {
				var t T
				if err := json.Unmarshal(v.BinaryValue, &t); err != nil {
					log.Printf("unmarshal %s err: %s", k, err)
				}
				ts = append(ts, t)
			}
		}
	}
	return ts, nil
}

func (s *svc) receiveWithDelete(queueURL string) (*sqs.DeleteMessageOutput, error) {
	result, err := s.receive(queueURL)
	if err != nil {
		return nil, err
	}
	if len(result.Messages) == 0 {
		return nil, nil
	}
	return s.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &queueURL,
		ReceiptHandle: result.Messages[0].ReceiptHandle,
	})
}

func realmain() error {
	var (
		region    = flag.String("region", "ap-northeast-1", "your region")
		queueName = flag.String("queueName", "my-test-queue-long", "queue name")
	)
	flag.Parse()
	svc := New(*region)
	urls, err := svc.listQueues()
	if err != nil {
		return err
	}
	for i, u := range urls {
		fmt.Printf("%d: %s\n", i, *u)
	}

	var queueURL string
	queueURL, err = svc.queueURL(*queueName)
	if err != nil || queueURL == "" {
		queueURL, err = svc.createQueue(*queueName)
		if err != nil {
			log.Fatalf("create queue failed: %s", err)
		}
	}

	// svc.sendBytes(queueURL)
	svc.sendJSON(queueURL, T{"test", []int{1, 2, 3}, []string{"id1", "id2"}, []float64{1.1, 2.1}})

	if _, err := svc.receiveWithLongPoll(queueURL); err != nil {
		log.Printf("receive message failed: %s", err)
	}
	// log.Println(received)

	ts, err := svc.receiveJSON(queueURL)
	if err != nil {
		panic(err)
	}
	for _, t := range ts {
		fmt.Printf("t = %+v\n", t)
	}

	// for _, m := range received.Messages {
	// for k, v := range m.MessageAttributes {
	// fmt.Printf("%s = %v\n", k, &v)
	// fmt.Printf("%s = %s\n", k, string(v.BinaryValue))
	// }
	// }

	// if err := svc.deleteQueue(queueURL); err != nil {
	// 	log.Fatalf("delete queue %s failed: %s", queueURL, err)
	// }
	return nil
}

func main() {
	if err := realmain(); err != nil {
		log.Fatal(err)
	}
}
