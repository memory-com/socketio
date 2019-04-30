package polling

import (
	"github.com/memory-com/pschlump/engineio/transport"
)

var Creater = transport.Creater{
	Name:      "polling",
	Upgrading: false,
	Server:    NewServer,
	Client:    NewClient,
}
