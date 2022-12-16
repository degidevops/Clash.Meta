package net

import (
	"net"
	"sync"
)

type closeSaveConn struct {
	net.Conn
	closeOnce sync.Once
	closeErr  error
}

func (conn *closeSaveConn) Close() error {
	conn.closeOnce.Do(func() {
		conn.closeErr = conn.Conn.Close()
	})
	return conn.closeErr
}

func NewCloseSaveConn(conn net.Conn) net.Conn {
	return &closeSaveConn{Conn: conn}
}

type closeSavePacketConn struct {
	net.PacketConn
	closeOnce sync.Once
	closeErr  error
}

func (conn *closeSavePacketConn) Close() error {
	conn.closeOnce.Do(func() {
		conn.closeErr = conn.PacketConn.Close()
	})
	return conn.closeErr
}

func NewCloseSavePacketConn(pc net.PacketConn) net.PacketConn {
	return &closeSavePacketConn{PacketConn: pc}
}
