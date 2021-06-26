package main

import "github.com/gorilla/websocket"

type client struct {
	socket *websocket.Conn
	send chan []byte
	room *room
}

func (c *client) read() {
	for {
		if _, message, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- message
		} else {
			break
		}
	}
	c.socket.Close()
}

func(c *client) write()  {
	for message := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, message);
			err != nil {
				break
		}
	}
	c.socket.Close()
}