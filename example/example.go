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
	dbo.alertObj("ok")
}
