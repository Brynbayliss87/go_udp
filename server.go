package main

import (
  "net"
  "fmt"
  "log"
  "time"
)

func main() {
  connection, err := net.ListenPacket("udp", ":10001")
  if err != nil {
    log.Fatal(err)
  }
  defer connection.Close()

  buffer := make([]byte, 1024)

  for {
    n, addr, _ := connection.ReadFrom(buffer)
    fmt.Println("Received at:", time.Now().Unix(), string(buffer[0:n]), "from", addr)

  }
}
