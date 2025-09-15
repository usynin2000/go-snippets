package main

import "fmt"

type Conn struct {
	Addr string
}

func NewConn(addr string) (*Conn, error) {
	if addr == "" {
		return nil, fmt.Errorf("empty addr")
	}
	return &Conn{Addr: addr}, nil
}

func main() {
	c, err := NewConn("localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Addr)
}