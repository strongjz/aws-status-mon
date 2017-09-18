package alert

import (
	"log"
)

//StandardOut - Logging the alert to Standard out
func StandardOut(service string, region string, whatsup string) {

	log.Printf("[ALERT] %s in %s %s\n", service, region, whatsup)

}
