package main
import (
//	"fmt"
	"log"
)

func main(){
	dbo, err := newDbusObject()
	if err != nil {
		log.Fatal("Error connecting to SystemBus: %v", err)
	}

	// obj.alertObj(message,showFullscreen,Priority)
	//Priority has 3 values, 0 for INFO, 1 for WARING and 3 for CRITICAL

	dbo.alertObj("INFO",false,0)

	//dbo.alertObj("WARNING",false,1)

	//dbo.alertObj("CRITICAL",true,2)
}
