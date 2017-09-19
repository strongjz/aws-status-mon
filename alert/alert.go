package alert

import (
	"log"
)

//StandardOut - Logging the alert to Standard out
func StandardOut(service string, region string, whatsup string) {

	log.Printf("[ALERT] %s in %s %s\n", service, region, whatsup)

}

/*
func SNSPublish(service string, whatsup string) error {

	return snsHelper(service, whatsup)

}

func snsHelper(service string, whatsup string) error {
	svc := sns.New(session.New())
	// params will be sent to the publish call included here is the bare minimum params to send a message.
	params := &sns.PublishInput{
		Message:  aws.String(fmt.Sprintf("Service %s - Issues %s", service, whatsup)), // This is the message itself (can be XML / JSON / Text - anything you want)
		TopicArn: aws.String(),                                                        //Get this from the Topic in the AWS console.
	}

	resp, err := svc.Publish(params) //Call to puclish the message

	if err != nil { //Check for errors
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return err
	}

	// Pretty-print the response data.
	log.Println(resp)
}
*/
