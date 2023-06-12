package line

import (
	"golang.org/x/net/websocket"
	"sync"
)

// Licensed under Creative Commons CC0.
//
// To the extent possible under law, the author(s) have dedicated all copyright
// neighboring rights to this software to the public domain worldwide.
// This software is distributed without any warranty.
// You should have received a copy of the CC0 Public Domain Dedication along wi
// If not, see <https:#creativecommons.org/publicdomain/zero/1.0/legalcode>.

// data.go - Relay websocket configuration and runtime data

// Turn this on for testing servers
var loopback = false

// Helps to prevent some rare crashes
var lock sync.Mutex

// Participant list to forward audio packets to
var peer = map[*websocket.Conn]int{}
