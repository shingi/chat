package main

import (
	"time"
	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan *message
	// room is the room this client is chatting in.
	room *room
	// userData holds information about the user
	userData map[string]interface{}
}

// read messages from socket, continously sending
// any received messages to the forward channel on the room type
func (c *client) read() {
	// close the socket when the function exists
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
        // get avatar URL if it exists
        if avatarURL, ok := c.userData["avatar_url"]; ok {
            msg.AvatarURL = avatarURL.(string)
        }
		c.room.forward <- msg
	}
}

// continously accept messages from the send channel and
// write everything out of the socket
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
