package main


import (
	"fmt"
	"flag"
	"errors"
	"strings"
	"strconv"
)

type NetAddress struct {
	Host string
	Port int
}

func (a NetAddress) String() string {
	return a.Host + ":" + strconv.Itoa(a.Port)
}

func (a *NetAddress) Set(s string) error {
	hp := strings.Split(s, ":")
	if len(hp) != 2 {
		return errors.New("Need address in a form host:port")
	}
	port, err := strconv.Atoi(hp[1])
	if err != nil {
		return err
	}
	a.Host = hp[0]
	a.Port = port
	return nil
}

func main() {
	addr := new(NetAddress)

	_ = flag.Value(addr)

	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()

	fmt.Println(addr.Host)
	fmt.Println(addr.Port)

}

// go run 4_FlagSet.go --addr=example.com:60