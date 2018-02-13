package client

import (
	"log"
	"net"

	"github.com/Hucaru/Valhalla/common/connection"
	"github.com/Hucaru/gopacket"
)

type Connection struct {
	conn      *connection.ClientConnection
	userID    uint32
	isLogedIn bool
	isAdmin   bool
	hash      string
	worldID   uint32
	chanID    byte
	gender    byte
}

func NewConnection(conn net.Conn) *Connection {
	loginConn := &Connection{conn: connection.NewClientConnection(conn), isAdmin: false}
	return loginConn
}

func (c *Connection) Write(p gopacket.Packet) error {
	return c.conn.Write(p)
}

func (c *Connection) Read(p gopacket.Packet) error {
	return c.conn.Read(p)
}

func (c *Connection) Close() {
	if c.isLogedIn {
		records, err := connection.Db.Query("UPDATE users set isLogedIn=0 WHERE userID=?", c.userID)

		defer records.Close()

		if err != nil {
			log.Println("Error in auto log out of user on disconnect, userID:", c.userID)
		}
	}

	c.conn.Close()
}

func (c *Connection) String() string {
	return c.conn.String()
}

func (c *Connection) SetUserID(val uint32) {
	c.userID = val
}

func (c *Connection) GetUserID() uint32 {
	return c.userID
}

func (c *Connection) SetAdmin(val bool) {
	c.isAdmin = val
}

func (c *Connection) IsAdmin() bool {
	return c.isAdmin
}

func (c *Connection) SetSessionHash(val string) {
	c.hash = val
}

func (c *Connection) GetSessionHash() string {
	return c.hash
}

func (c *Connection) SetIsLogedIn(val bool) {
	c.isLogedIn = val
}

func (c *Connection) GetIsLogedIn() bool {
	return c.isLogedIn
}

func (c *Connection) SetWorldID(val uint32) {
	c.worldID = val
}

func (c *Connection) GetChanID() byte {
	return c.chanID
}

func (c *Connection) SetChanID(val byte) {
	c.chanID = val
}

func (c *Connection) GetWorldID() uint32 {
	return c.worldID
}

func (c *Connection) SetGender(val byte) {
	c.gender = val
}

func (c *Connection) GetGender() byte {
	return c.gender
}