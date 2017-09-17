package alert

import (
	"log"
)

func StandardOut(service string, region string, whatsup string) {

	log.Printf("ALERT: %s in %s\nThis is broke %s\n", service, region, whatsup)

}
