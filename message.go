package main

import (
	"time"
)

// represents a single message
type message struct {
	Name string
	Message string
	When time.Time
    AvatarURL string
}
