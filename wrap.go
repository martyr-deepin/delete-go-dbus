package dbus

import "pkg.linuxdeepin.com/dbus/transport"

type ObjectPath transport.ObjectPath
type Conn transport.Conn
type Message transport.Message

var SignatureOfType = transport.SignatureOfType

type Variant transport.Variant

func SessionBus() (*Conn, error) {
	return transport.SessionBus()
}
func SystemBus() (*Conn, error) {
	return transport.SystemBus()
}
