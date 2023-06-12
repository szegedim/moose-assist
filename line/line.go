package line

import (
	"fmt"
	"gitlab.com/eper.io/engine/metadata"
	"golang.org/x/net/websocket"
	"math"
	"net/http"
	"os"
	"path"
	"time"
)

// main.go - Relay websocket connections for streaming applications. See README.md for usage

// Licensed under Creative Commons CC0.
//
// To the extent possible under law, the author(s) have dedicated all copyright
// neighboring rights to this software to the public domain worldwide.
// This software is distributed without any warranty.
// You should have received a copy of the CC0 Public Domain Dedication along wi
// If not, see <https:#creativecommons.org/publicdomain/zero/1.0/legalcode>.

func Setup() {
	activationKey := os.Getenv("APIKEY")
	if activationKey != "" {
		metadata.ActivationKey = activationKey
	}
	siteUrl := os.Getenv("SITEURL")
	if siteUrl != "" {
		metadata.SiteUrl = siteUrl
	}
	fmt.Printf("Try local %s\n", metadata.SiteUrl+"/line.html?apikey="+metadata.ActivationKey+"#generate_leaf")
	http.Handle("/ws", websocket.Handler(relayLine))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/line.html", http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/line.html", func(writer http.ResponseWriter, request *http.Request) {
		if "" == request.URL.Query().Get("apikey") && metadata.ActivationKey == "" {
			if metadata.ActivationKey != "" {
				http.Redirect(writer, request, fmt.Sprintf("/line.html?apikey=%s#generate_leaf", metadata.ActivationKey), http.StatusTemporaryRedirect)
			} else {
				http.Redirect(writer, request, fmt.Sprintf("/documentation.html"), http.StatusTemporaryRedirect)
			}
			return
		}
		http.ServeFile(writer, request, path.Join("./line/res/line.html"))
	})

	http.HandleFunc("/news.html", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, path.Join("./line/res/news.html"))
	})

	http.HandleFunc("/moose.png", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, path.Join("./line/res/moose.png"))
	})

	http.HandleFunc("/main.css", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, path.Join("./line/res/main.css"))
	})

	http.HandleFunc("/e2e.js", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, path.Join("./line/res/e2e.js"))
	})

	http.HandleFunc("/pcm.js", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, path.Join("./line/res/pcm.js"))
	})

	http.HandleFunc("/documentation.html", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, path.Join("./line/res/documentation.html"))
	})
}

func relayLine(ws *websocket.Conn) {
	start := time.Now()
	var lastSecond float64
	defer func() { _ = ws.Close() }()
	apiKey := ws.Request().URL.Query().Get("apikey")
	if metadata.ActivationKey != apiKey {
		time.Sleep(15 * time.Millisecond)
		fmt.Println("error")
		return
	}
	lock.Lock()
	if len(peer) > 1 {
		fmt.Println("connection cannot exceed", peer)
		return
	}
	ws.PayloadType = websocket.BinaryFrame
	peer[ws] = 1
	lock.Unlock()
	fmt.Println("connection received")

	finished := false
	for !finished {
		buf := make([]byte, math.MaxUint16)
		n, _ := ws.Read(buf)
		if n == 0 {
			buf = nil
			fmt.Println("connection lost")
			break
		}
		lock.Lock()
		for cr := range peer {
			if cr != ws || loopback {
				m, _ := cr.Write(buf[0:n])
				currentSecond := time.Now().Sub(start).Seconds()
				if currentSecond > lastSecond+3.0 {
					lastSecond = currentSecond
					go func(buffer []byte) {
						var min, max, num, sum float64
						min = math.MaxFloat64
						max = -math.MaxFloat64
						for i := range buffer {
							if i%2 == 1 && i > 0 {
								pulse := (int16(buffer[i])+math.MinInt8)*math.MaxUint8 + (int16(buffer[i-1]) + math.MinInt8)
								if float64(pulse) > max {
									max = float64(pulse)
								}
								if float64(pulse) < min {
									min = float64(pulse)
								}
								sum = sum + float64(pulse)
								num++
							}
						}
						sum = sum / num
						fmt.Printf("%6.2f %6.2f %6.2f\n", min, sum, max)
					}(buf)
				}
				if m == 0 {
					finished = true
					fmt.Println("connection closed")
					break
				}
			}
		}
		lock.Unlock()
	}
	lock.Lock()
	delete(peer, ws)
	if len(peer) == 0 {
		fmt.Println("lost all connections")
	}
	lock.Unlock()
}
