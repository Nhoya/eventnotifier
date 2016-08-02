package main

import (
	"log"
	)

var dn *DesktopNotifications

func main() {

	dbs, err := newDbusServer();
        dn = newDesktopNotifications()

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	for dbs.run == true {
		/* Run forever */
        }

}
