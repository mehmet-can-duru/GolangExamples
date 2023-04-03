package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func logError(err error) {
	if err != nil {
		logFile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Failed to write to log file:", err)
			return
		}
		defer logFile.Close()

		logTime := time.Now().Format("2006-01-02 15:04:05")
		logMsg := fmt.Sprintf("[%s] Error: %s\n", logTime, err.Error())
		logFile.WriteString(logMsg)

		fmt.Println("Error logged:", err)
	}
}

func main() {
	// listen server and port number
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	logError(err)

	// UDP connection
	conn, err := net.ListenUDP("udp", addr)
	logError(err)
	defer conn.Close()

	fmt.Println("UDP server ", addr, " listening at.")

	// A map that holds players' connection information
	players := make(map[string]*net.UDPAddr)

	// endless loop
	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		logError(err)
		// process the message
		msg := string(buf[0:n])
		switch msg {
		case "connect":
			// connect address added
			players[addr.String()] = addr
			fmt.Println("New player connected:", addr)
		case "disconnect":
			/// connect address deleted
			delete(players, addr.String())
			fmt.Println("Player disconnected:", addr)
		case "update":
			// player location update, message to other players
			for _, playerAddr := range players {
				if playerAddr.String() != addr.String() {
					conn.WriteToUDP([]byte(msg), playerAddr)
				}
			}
		}
	}
}
