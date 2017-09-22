package alert

import (
	"github.com/spf13/viper"
	"log"
)

func Alert(config *viper.Viper, service string, region string, whatsup string) {

	log.Println("Alerting")

	alerts_enabled := config.GetStringSlice("enabled_alerts")

	for _, a := range alerts_enabled {
		log.Printf("Alert Enabled %s", a)
		switch {
		case a == "sns":

			topicARN := config.GetString("sns_topic_arn")
			running_region := config.GetString("region")

			awsSNS(topicARN, running_region, service, region, whatsup)

		case a == "log":
			standardOut(service, region, whatsup)
		}
	}
}
