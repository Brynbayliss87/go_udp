package main

import (
  "net"
  "log"
)

func main() {
  Conn, err := net.Dial("udp", "127.0.0.1:10001")
  if err != nil {
    log.Fatal(err)
  }

  defer Conn.Close()
  Conn.Write([]byte("Hello, This message was sent over UDP, I hope it arrives!"))
}
