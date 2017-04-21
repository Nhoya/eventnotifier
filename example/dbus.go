package main

import (
//  "fmt"
        "github.com/godbus/dbus"
)

type dbusObject struct {
        dbus.BusObject
}

type slmData struct {
    Msg string
    FullScreen bool
}

func newDbusObject() (*dbusObject, error) {
        conn, err := dbus.SystemBus()

        if err != nil {
                return nil, err
        }

        return &dbusObject{conn.Object("com.subgraph.EventNotifier", "/com/subgraph/EventNotifier")}, nil
}

func (ob *dbusObject) alertObj(Msg string) {
//  fmt.Println("id = ", id)
//  fmt.Println("xyz: ", line)
    dobj := slmData{Msg, true}
    call := ob.Call("com.subgraph.EventNotifier.Alert", 0, dobj)
    if call.Err != nil {
        panic(call.Err)
    }
}


