package quic

import (
	"fmt"
	"net"
)

var _ net.Addr = (*netAddr)(nil)

type netAddr struct {
	network  string
	hostport string
}

func (n *netAddr) Network() string {
	return n.network
}

func (n *netAddr) String() string {
	return n.hostport
}

func (c *Conn) LocalAddr() net.Addr {
	addrPort := c.endpoint.LocalAddr()
	return &netAddr{
		"udp",
		fmt.Sprintf("%s:%d", addrPort.Addr().String(), addrPort.Port()),
	}
}

func (c *Conn) RemoteAddr() net.Addr {
	addrPort := c.peerAddr
	return &netAddr{
		"udp",
		fmt.Sprintf("%s:%d", addrPort.Addr().String(), addrPort.Port()),
	}
}
