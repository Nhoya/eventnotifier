package main

import (
//  "fmt"
        "github.com/godbus/dbus"
)

const busName = "com.NSAway.EventNotifier"
const objectPath = "/com/NSAway/EventNotifier"

type dbusObject struct {
        dbus.BusObject
}

type slmData struct {
    Msg string
    FullScreen bool
    Priority int
}

func newDbusObject() (*dbusObject, error) {
        conn, err := dbus.SystemBus()

        if err != nil {
                return nil, err
        }

        return &dbusObject{conn.Object(busName, objectPath)}, nil
}

func (ob *dbusObject) alertObj(Msg string, FullScreen bool, Priority int) {
    dobj := slmData{Msg, FullScreen, Priority}
    call := ob.Call("com.NSAway.EventNotifier.Alert", 0, dobj)
    if call.Err != nil {
        panic(call.Err)
    }
}


