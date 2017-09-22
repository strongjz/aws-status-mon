package alert

import "log"

//StandardOut - Logging the alert to Standard out
func standardOut(service string, region string, whatsup string) {

	log.Printf("[ALERT][LOG] %s in %s %s\n", service, region, whatsup)

}
