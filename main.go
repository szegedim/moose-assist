package main

import (
	"gitlab.com/eper.io/engine/line"
	"net/http"
)

// main.go - Relay websocket connections for streaming applications. See README.md for usage

// Licensed under Creative Commons CC0.
// To the extent possible under law, the author(s) have dedicated all copyright and related and
// neighboring rights to this software to the public domain worldwide.
// This software is distributed without any warranty.
// You should have received a copy of the CC0 Public Domain Dedication along with this software.
// If not, see <https://creativecommons.org/publicdomain/zero/1.0/legalcode>.

// This example demonstrates a trivial echo server.
func main() {
	line.Setup()
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
