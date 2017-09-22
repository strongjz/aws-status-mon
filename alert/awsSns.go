package alert

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

func awsSNS(topic string, running_region string, service string, region string, whatsup string) {

	log.Printf("[ALERT][SNS] %s in %s %s\n", service, region, whatsup)

	message := fmt.Sprintf("[ALERT][SNS] %s in %s %s\n", service, region, whatsup)

	// Create a Session with a custom region
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(running_region),
	}))

	svc := sns.New(sess)

	params := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(topic),
	}

	resp, err := svc.Publish(params)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	log.Println(resp)

}
