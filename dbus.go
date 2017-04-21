package main

import (
	"errors"
	"github.com/godbus/dbus"
)

const busName = "com.NSAway.EventNotifier"
const objectPath = "/com/NSAway/EventNotifier"
const interfaceName = "com.NSAway.EventNotifier"

type dbusServer struct {
	conn *dbus.Conn
	run bool
}

type slmData struct {
    Msg string
    FullScreen bool
    Priority int
}

func newDbusServer() (*dbusServer, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	reply, err := conn.RequestName(busName, dbus.NameFlagDoNotQueue)
	if err != nil {
		return nil, err
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		return nil, errors.New("Bus name is already owned")
	}
	ds := &dbusServer{}

	if err := conn.Export(ds, objectPath, interfaceName); err != nil {
		return nil, err
	}

	ds.conn = conn
	ds.run = true
	return ds, nil
}

func (ds *dbusServer) Alert(obj slmData) *dbus.Error {
    dn.show("sysevent", obj.Msg, obj.FullScreen, obj.Priority)

	return nil
}
