package main

import (
    "fmt"
    "net"
    "encoding/json"
)

const (
    proto = "tcp"
    serv = "localhost:7777"
    welcome = "\n--| ~Ivy~ Version: XXX |--\n"
)

func handleit(e error) {
    if e != nil {
        fmt.Println("OH NO: %s", e)
    }
}

type Message struct {
    Type      int
    Root      string
    Parent    string
    UUID      string
    Recent    []string
    Content   string
    Username  string
    Timestamp int
    Major     int
    Minor     int
}

func main() {
    fmt.Println(welcome)

	var m Message
    conn, err := net.Dial(proto, serv)
	handleit(err)

	response := json.NewDecoder(conn)

	e := response.Decode(&m)
	handleit(e)

	fmt.Printf("Type: %d\nRoot Message: %s\nRecent Messages: %v\nServer Version %d.%d\n", m.Type, m.Root, m.Recent, m.Major, m.Minor)
}
